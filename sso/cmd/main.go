package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/uvuv-643/Web_Construct/common/proto/pkg/sso"
	"github.com/uvuv-643/Web_Construct/sso/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strings"
)

type server struct {
	sso.UnimplementedAuthServer
	sso.UnimplementedPermissionsServer
	userRepo internal.UserRepository
}

func (s *server) createJWTForUser(user *internal.User) (string, error) {
	mySigningKey := []byte("AllYourBase")
	claims := &jwt.RegisteredClaims{
		Subject: string(user.Email),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func (s *server) getUserFromJWT(tokenString string) (*internal.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	switch {
	case token.Valid:
		email, err := token.Claims.GetSubject()
		if err != nil {
			return nil, err
		}
		ctx := context.Background()
		return s.userRepo.GetByEmail(ctx, email)
	default:
		return nil, err
	}
}

func (s *server) Register(_ context.Context, in *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	ctx := context.Background()
	existingUser, err := s.userRepo.GetByEmailAndPassword(ctx, strings.ToLower(in.Email), in.Password)
	if existingUser != nil {
		return nil, status.Errorf(codes.AlreadyExists, "User already exists")
	}
	_, err = s.userRepo.Create(ctx, internal.Email(strings.ToLower(in.Email)), in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "SSO server error")
	}
	return &sso.RegisterResponse{}, nil
}

func (s *server) Login(_ context.Context, in *sso.LoginRequest) (*sso.LoginResponse, error) {
	ctx := context.Background()
	user, err := s.userRepo.GetByEmailAndPassword(ctx, in.Email, in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid username or password")
	} else {
		jwtString, err := s.createJWTForUser(user)
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, "Internal server error")
		}
		return &sso.LoginResponse{Jwt: jwtString}, nil
	}
}

func (s *server) GetUserPermissions(_ context.Context, in *sso.GetUserPermissionsRequest) (*sso.UserPermissions, error) {
	user, err := s.getUserFromJWT(in.Jwt)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid username or password")
	}
	userPermissions := make(map[string][]sso.PermissionType)
	for _, role := range user.Roles {
		if userPermissions[role.ApplicationID] == nil {
			userPermissions[role.ApplicationID] = make([]sso.PermissionType, 0)
		} else {
			userPermissions[role.ApplicationID] = append(userPermissions[role.ApplicationID], internal.GetPermissionType(string(role.Role)))
		}
	}

	return &sso.UserPermissions{UserId: "hello", Apps: []*sso.AppPermission{
		{AppUuid: "backend", Permissions: []sso.PermissionType{sso.PermissionType_PT_SHARE}},
	}}, nil
}

var (
	port = flag.Int("port", 50052, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sso.RegisterAuthServer(s, &server{})
	sso.RegisterPermissionsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//package main
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/go-pg/pg/v10"
//	"github.com/uvuv-643/Web_Construct/sso/internal"
//	"log"
//	"net/http"
//	"strings"
//
//	"github.com/gorilla/mux"
//)
//
//type Server struct {
//	router   *mux.Router
//	userRepo internal.UserRepository
//}
//
//func NewServer(userRepo internal.UserRepository) *Server {
//	s := &Server{
//		router:   mux.NewRouter(),
//		userRepo: userRepo,
//	}
//	s.routes()
//	return s
//}
//
//func (s *Server) routes() {
//	s.router.HandleFunc("/users", s.createUser).Methods("POST")
//	s.router.HandleFunc("/users", s.getUser).Methods("GET")
//}
//
//func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	s.router.ServeHTTP(w, r)
//}
//
//func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
//	var input struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//		FullName string `json:"full_name"`
//	}
//
//	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//		http.Error(w, "Invalid request payload (see api docs)", http.StatusBadRequest)
//		return
//	}
//
//	if strings.ToLower(input.Email) == "" {
//		http.Error(w, "Email is required", http.StatusBadRequest)
//		return
//	}
//	if input.Password == "" {
//		http.Error(w, "Password is required", http.StatusBadRequest)
//		return
//	}
//	if input.FullName == "" {
//		http.Error(w, "Full name is required", http.StatusBadRequest)
//		return
//	}
//
//	ctx := context.Background()
//
//	existingUser, err := s.userRepo.GetByEmailAndPassword(ctx, strings.ToLower(input.Email), input.Password)
//	if existingUser != nil {
//		http.Error(w, "User with this email already exists", http.StatusBadRequest)
//		return
//	}
//
//	err = s.userRepo.Create(ctx, internal.Email(strings.ToLower(input.Email)), input.Password, input.FullName)
//	if err != nil {
//		http.Error(w, fmt.Sprintf("Failed to create user: %s", err.Error()), http.StatusBadRequest)
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//}
//
//func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
//	var input struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
//
//	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//		http.Error(w, "Invalid request payload", http.StatusBadRequest)
//		return
//	}
//
//	ctx := context.Background()
//	user, err := s.userRepo.GetByEmailAndPassword(ctx, input.Email, input.Password)
//	if err != nil {
//		http.Error(w, fmt.Sprintf("Failed to retrieve user: %s", err.Error()), http.StatusUnauthorized)
//		return
//	}
//
//	if err := json.NewEncoder(w).Encode(user); err != nil {
//		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
//	}
//}
//
//func main() {
//	db := internal.ConnectToDB()
//	defer func(db *pg.DB) {
//		err := db.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(db)
//
//	userRepo := internal.NewUserRepository(db)
//	server := NewServer(userRepo)
//
//	err := http.ListenAndServe(":8080", server)
//	if err != nil {
//		return
//	}
//
//}
