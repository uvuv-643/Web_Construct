package internal

import (
	"context"
	"flag"
	config "github.com/uvuv-643/Web_Construct/backend/conifg"
	"github.com/uvuv-643/Web_Construct/common/proto/pkg/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func GetUserPermissions(jwt string) []sso.PermissionType {

	cfg := config.New()
	addr := cfg.SSOUrl
	appUuid := cfg.ApplicationUUID

	flag.Parse()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	c := sso.NewPermissionsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUserPermissions(ctx, &sso.GetUserPermissionsRequest{Jwt: jwt})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	permissions := r.GetApps()
	for _, permission := range permissions {
		if permission.AppUuid == appUuid {
			return permission.Permissions
		}
	}

	return make([]sso.PermissionType, 0)
}

func Register(email, password string) (*sso.RegisterResponse, error) {

	cfg := config.New()
	addr := cfg.SSOUrl

	flag.Parse()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	c := sso.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.Register(ctx, &sso.RegisterRequest{Email: email, Password: password})

	return response, err

}

func Login(email, password string) (*sso.LoginResponse, error) {

	cfg := config.New()
	addr := cfg.SSOUrl

	flag.Parse()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	c := sso.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.Login(ctx, &sso.LoginRequest{Email: email, Password: password})

	return response, err

}
