package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
}

type HttpServer struct {
	router *mux.Router
}

func NewServer() *HttpServer {
	s := &HttpServer{
		router: mux.NewRouter(),
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
		jwtToken := strings.Trim(strings.Replace(authHeader, "Bearer:", "", 1), " ")
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
		next.ServeHTTP(w, r)
	})
}

func (s *HttpServer) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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

func (s *server) SendRequest(_ context.Context, in *llmproxy.LLMRequest) (*emptypb.Empty, error) {
	fmt.Println("Send Request to server", in.Jwt, in.Content)
	fmt.Println(internal.GetUserPermissions(in.Jwt))
	//fmt.Println(internal.ValidateAIProxyPermissions(internal.GetUserPermissions(in.Jwt)))
	return nil, nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func startGrpcServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	llmproxy.RegisterLLMProxyServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startHttpServer() {
	server := NewServer()
	log.Printf("server listening at [::]:8080")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		return
	}
}

func main() {
	go startGrpcServer()
	startHttpServer()
}
