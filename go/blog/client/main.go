package main

import (
	"context"
	"fmt"
	"log"

	blogv1 "github.com/shuymn-sandbox/go-grpc-sample/go/protobuf/protoc/blog/v1"

	"google.golang.org/grpc"
)

const serverAddr = "127.0.0.1:8080"

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run(ctx context.Context) error {
	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to BlogService on %s: %w", serverAddr, err)
	}
	log.Println("connect to", serverAddr)

	postSvc := blogv1.NewPostServiceClient(conn)
	resp, err := postSvc.GetPost(ctx, &blogv1.GetPostRequest{
		PostId: 1,
	})
	if err != nil {
		return fmt.Errorf("failed to GetPost: %w", err)
	}

	log.Printf("successfully GetPost: %+v\n", resp)
	return nil
}
