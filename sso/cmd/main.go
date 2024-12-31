package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-pg/pg/v10"
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
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		return nil, err
	}
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
	fmt.Println(in)
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
		}
		userPermissions[role.ApplicationID] = append(
			userPermissions[role.ApplicationID],
			internal.GetPermissionsType(string(role.Role)),
		)
	}
	appPermissions := make([]*sso.AppPermission, 0)
	for key, permissions := range userPermissions {
		appPermissions = append(appPermissions, &sso.AppPermission{
			AppUuid: key, Permissions: permissions,
		})
	}
	return &sso.UserPermissions{UserId: user.ID.String(), Apps: appPermissions}, nil
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

	db := internal.ConnectToDB()
	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	userRepo := internal.NewUserRepository(db)

	sso.RegisterAuthServer(s, &server{userRepo: userRepo})
	sso.RegisterPermissionsServer(s, &server{userRepo: userRepo})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
