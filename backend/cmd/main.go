package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	config "github.com/uvuv-643/Web_Construct/backend/conifg"
	"github.com/uvuv-643/Web_Construct/backend/internal"
	"github.com/uvuv-643/Web_Construct/common/proto/pkg/llmproxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"net/http"
	"strings"
)

type server struct {
	llmproxy.UnimplementedLLMProxyServer
	orderRepo internal.OrderRepository
}

type HttpServer struct {
	router    *mux.Router
	orderRepo internal.OrderRepository
}

func NewServer(orderRepo internal.OrderRepository) *HttpServer {
	s := &HttpServer{
		router:    mux.NewRouter(),
		orderRepo: orderRepo,
	}
	s.routes()
	return s
}

func (s *HttpServer) routes() {
	authRoutes := s.router.PathPrefix("/api/auth").Subrouter()
	authRoutes.HandleFunc("/register", s.createUser).Methods("POST")
	authRoutes.HandleFunc("/login", s.getUser).Methods("POST")

	allRoutes := s.router.PathPrefix("/api").Subrouter()
	allRoutes.Use(s.validateJWT)
	allRoutes.HandleFunc("/hello", s.hello).Methods("POST")

}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *HttpServer) validateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		jwtToken := strings.Trim(strings.Replace(authHeader, "Bearer", "", 1), " ")
		_, err := internal.GetUserPermissions(jwtToken)
		if err != nil {
			st, _ := status.FromError(err)
			if st.Code() == codes.Unauthenticated {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		r.Header.Add("userValidated", jwtToken)
		next.ServeHTTP(w, r)
	})
}

func (s *HttpServer) hello(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Request string `json:"request"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "`request` field is required", http.StatusBadRequest)
		return
	}

	order, err := s.orderRepo.Create(context.Background(), r.Header.Get("userValidated"), input.Request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = internal.SendRequestToLLM(input.Request, order)
	if err != nil {
		s.orderRepo.delete(order)
		fmt.Println()
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(order.ID.String()))
	if err != nil {
		s.orderRepo.delete(order)
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *HttpServer) createUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload (see api docs)", http.StatusBadRequest)
		return
	}

	if strings.ToLower(input.Email) == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if input.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	_, err := internal.Register(strings.ToLower(input.Email), input.Password)
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.AlreadyExists {
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		} else {
			fmt.Println(st.Code())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *HttpServer) getUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := internal.Login(strings.ToLower(input.Email), input.Password)
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.Unauthenticated {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	if err := json.NewEncoder(w).Encode(response.Jwt); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *server) SendReply(_ context.Context, in *llmproxy.LLMReply) (*emptypb.Empty, error) {
	fmt.Println("Send Request to server", in.Jwt)
	fmt.Println(internal.GetUserPermissions(in.Jwt))
	permissions, err := internal.GetUserPermissions(in.Jwt)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Unauthenticated, "Invalid token")
	}
	if !internal.ValidateAIProxyPermissions(permissions) {
		return &emptypb.Empty{}, status.Errorf(codes.PermissionDenied, "Have no PT_SHARE")
	}
	fmt.Println(in.Uuid)
	id, err := uuid.Parse(in.Uuid)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Unknown, "Invalid uuid")
	}
	fmt.Println("Modify triggered")
	_, err = s.orderRepo.Modify(context.Background(), id, in.Response)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "Invalid uuid")
	}
	return &emptypb.Empty{}, nil
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func startGrpcServer(orderRepo internal.OrderRepository) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.New().BackendGrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	llmproxy.RegisterLLMProxyServer(s, &server{orderRepo: orderRepo})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startHttpServer(orderRepo internal.OrderRepository) {
	server := NewServer(orderRepo)
	log.Printf("server listening at [::]:8080")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		return
	}
}

func main() {
	db := internal.ConnectToDB()
	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	orderRepo := internal.NewOrderRepository(db)
	go startGrpcServer(orderRepo)
	startHttpServer(orderRepo)
}
