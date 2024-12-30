//package main
//
//import (
//	"context"
//	"flag"
//	"fmt"
//	"github.com/uvuv-643/Web_Construct/common/proto/pkg/sso"
//	"google.golang.org/grpc"
//	"log"
//	"net"
//)
//
//type server struct {
//	sso.UnimplementedAuthServer
//	sso.UnimplementedPermissionsServer
//}
//
//func (s *server) Register(_ context.Context, in *sso.RegisterRequest) (*sso.RegisterResponse, error) {
//	fmt.Println("Send Register Request to server", in.Email, in.Password)
//	return nil, nil
//}
//
//func (s *server) Login(_ context.Context, in *sso.LoginRequest) (*sso.LoginResponse, error) {
//	fmt.Println("Send Login Request to server", in.Email, in.Password)
//	return nil, nil
//}
//
//func (s *server) GetUserPermissions(_ context.Context, in *sso.GetUserPermissionsRequest) (*sso.UserPermissions, error) {
//	return &sso.UserPermissions{UserId: "hello", Apps: []*sso.AppPermission{
//		{AppUuid: "backend", Permissions: []sso.PermissionType{sso.PermissionType_PT_SHARE}},
//	}}, nil
//}
//
//var (
//	port = flag.Int("port", 50052, "The server port")
//)
//
//func main() {
//	flag.Parse()
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	sso.RegisterAuthServer(s, &server{})
//	sso.RegisterPermissionsServer(s, &server{})
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/uvuv-643/Web_Construct/sso/internal"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	userRepo *internal.UserRepository
}

func NewServer(userRepo *internal.UserRepositoryImpl) *Server {
	s := &Server{
		router:   mux.NewRouter(),
		userRepo: userRepo,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("/users", s.createUser).Methods("POST")
	s.router.HandleFunc("/users/{id}", s.getUser).Methods("GET")
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err := s.userRepo.Create(ctx, Email(strings.ToLower(input.Email)), input.Password, input.FullName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %s", err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	user, err := s.userRepo.GetByEmailAndPassword(ctx, input.Email, input.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve user: %s", err.Error()), http.StatusUnauthorized)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
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

	userRepo := internal.NewUserRepository(db)
	server := NewServer(userRepo)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		return
	}

}
