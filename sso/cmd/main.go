package main

import (
	"flag"
	"fmt"
	"github.com/uvuv-643/Web_Construct/sso/internal"
	"log"
	"net"

	"github.com/uvuv-643/Web_Construct/common/proto/pkg/sso"

	"google.golang.org/grpc"
)

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
	defer db.Close()

	userRepo := internal.NewUserRepository(db)
	authServer := internal.NewAuthServer(userRepo)

	sso.RegisterAuthServer(s, authServer)
	sso.RegisterPermissionsServer(s, authServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
