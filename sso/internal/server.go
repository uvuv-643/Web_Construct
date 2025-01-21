package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/uvuv-643/Web_Construct/common/proto/pkg/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	sso.UnimplementedAuthServer
	sso.UnimplementedPermissionsServer
	userRepo UserRepository
}

func NewAuthServer(repo UserRepository) *AuthServer {
	return &AuthServer{userRepo: repo}
}

func (s *AuthServer) createJWTForUser(user *User) (string, error) {
	mySigningKey := []byte("AllYourBase")
	claims := &jwt.RegisteredClaims{
		Subject: string(user.Email),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func (s *AuthServer) getUserFromJWT(tokenString string) (*User, error) {
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("hello")
	switch {
	case token.Valid:
		email, err := token.Claims.GetSubject()
		if err != nil {
			return nil, err
		}
		ctx := context.Background()
		return s.userRepo.GetByEmail(ctx, email)
	default:
		return nil, status.Errorf(codes.Unauthenticated, "Invalid username or password")
	}
}

func (s *AuthServer) Register(ctx context.Context, in *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	existingUser, err := s.userRepo.GetByEmailAndPassword(ctx, strings.ToLower(in.Email), in.Password)
	if existingUser != nil {
		return nil, status.Errorf(codes.AlreadyExists, "User already exists")
	}
	_, err = s.userRepo.Create(ctx, Email(strings.ToLower(in.Email)), in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "SSO server error")
	}
	return &sso.RegisterResponse{}, nil
}

func (s *AuthServer) Login(ctx context.Context, in *sso.LoginRequest) (*sso.LoginResponse, error) {
	user, err := s.userRepo.GetByEmailAndPassword(ctx, in.Email, in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid username or password")
	}
	jwtString, err := s.createJWTForUser(user)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Internal server error")
	}
	return &sso.LoginResponse{Jwt: jwtString}, nil
}

func (s *AuthServer) GetUserPermissions(ctx context.Context, in *sso.GetUserPermissionsRequest) (*sso.UserPermissions, error) {
	user, err := s.getUserFromJWT(in.Jwt)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid JWT")
	}
	userPermissions := BuildUserPermissions(user)
	return &sso.UserPermissions{
		UserId: user.ID.String(),
		Apps:   userPermissions,
	}, nil
}
