package internal

import (
	"context"
	config "github.com/uvuv-643/Web_Construct/backend/conifg"
	"github.com/uvuv-643/Web_Construct/common/proto/pkg/llmproxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func SendRequestToLLM(data string) error {

	cfg := config.New()
	addr := cfg.ProxyURL

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := llmproxy.NewLLMProxyClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SendRequest(ctx, &llmproxy.LLMRequest{Jwt: cfg.ServiceJWT, Content: data})
	if err != nil {
		return err
	}
	return nil

}
