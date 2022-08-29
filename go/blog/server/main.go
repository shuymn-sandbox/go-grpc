package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// blogv1 "github.com/shuymn-sandbox/go-grpc-sample/go/protobuf/protoc/blog/v1"
	blogv1 "github.com/shuymn-sandbox/go-grpc-sample/go/protobuf/buf/blog/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const addr = "127.0.0.1:8080"

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run() error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	server := grpc.NewServer()
	blogv1.RegisterPostServiceServer(server, &postServiceServer{
		db: NewDB(),
	})

	log.Println("listening on", addr)

	if err = server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}
	return nil
}

type postServiceServer struct {
	db DB
}

func (s *postServiceServer) GetPost(ctx context.Context, r *blogv1.GetPostRequest) (*blogv1.GetPostResponse, error) {
	postID := r.GetPostId()
	post, ok := s.db.GetPost(int(postID))
	if !ok {
		return nil, fmt.Errorf("post_id: %d not found", postID)
	}
	user, ok := s.db.GetUser(post.AuthorID)
	if !ok {
		return nil, fmt.Errorf("author_id: %d not found", post.AuthorID)
	}
	return &blogv1.GetPostResponse{
		Post: &blogv1.Post{
			Id: int64(post.ID),
			Author: &blogv1.User{
				Id:   int64(user.ID),
				Name: user.Name,
			},
			Title:       post.Title,
			Content:     post.Content,
			Status:      blogv1.PostStatus(post.Status),
			PublishedAt: timestamppb.New(post.PublishedAt),
			CreatedAt:   timestamppb.New(post.CreatedAt),
			UpdatedAt:   timestamppb.New(post.UpdatedAt),
		},
	}, nil
}
