package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/uvuv-643/Web_Construct/common/proto/pkg/sso"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	sso.UnimplementedAuthServer
	sso.UnimplementedPermissionsServer
}

func (s *server) Register(_ context.Context, in *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	fmt.Println("Send Register Request to server", in.Email, in.Password)
	return nil, nil
}

func (s *server) Login(_ context.Context, in *sso.LoginRequest) (*sso.LoginResponse, error) {
	fmt.Println("Send Login Request to server", in.Email, in.Password)
	return nil, nil
}

func (s *server) GetUserPermissions(_ context.Context, in *sso.GetUserPermissionsRequest) (*sso.UserPermissions, error) {
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
