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
	"github.com/uvuv-643/Web_Construct/sso/internal"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	userRepo internal.UserRepository
}

func NewServer(userRepo internal.UserRepository) *Server {
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
	s.router.HandleFunc("/users/{id}", s.updateUser).Methods("PUT")
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var user internal.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	if err := s.userRepo.Create(ctx, &user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	user, err := s.userRepo.Get(ctx, &internal.FindUserOptions{ID: id})
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user internal.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	user.ID = id

	ctx := context.Background()
	if err := s.userRepo.Update(ctx, &user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	db := internal.ConnectToDB()
	defer db.Close()

	userRepo := internal.NewUserRepository(db)
	server := NewServer(userRepo)

	http.ListenAndServe(":8080", server)
}
