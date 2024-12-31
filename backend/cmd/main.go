package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/uvuv-643/Web_Construct/backend/internal"
	"github.com/uvuv-643/Web_Construct/common/proto/pkg/llmproxy"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type server struct {
	llmproxy.UnimplementedLLMProxyServer
}

func (s *server) SendRequest(_ context.Context, in *llmproxy.LLMRequest) (*emptypb.Empty, error) {
	fmt.Println("Send Request to server", in.Jwt, in.Content)
	fmt.Println(internal.GetUserPermissions(in.Jwt))
	fmt.Println(internal.ValidateAIProxyPermissions(internal.GetUserPermissions(in.Jwt)))
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

func main() {

	//fmt.Println(internal.Register("abacaba@gmail.com", "uvuv"))
	fmt.Println(internal.Login("abacaba@gmail.com", "uvuv"))
	fmt.Println(internal.GetUserPermissions("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhYmFjYWJhQGdtYWlsLmNvbSJ9.Ci3c190C4et5h2PBITpU-zV9jThReOc5mj2erH1ymOw"))

	//flag.Parse()
	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()
	//llmproxy.RegisterLLMProxyServer(s, &server{})
	//log.Printf("server listening at %v", lis.Addr())
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
