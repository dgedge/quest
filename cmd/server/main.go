package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/dgedge/quest/pkg/proto/question"
	"google.golang.org/grpc"
)

type server struct {
	question.UnimplementedQuestionServiceServer
}

func main() {
	log.Println("Server running ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	question.RegisterQuestionServiceServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		panic(err)
	}

	log.Fatalln(srv.Serve(lis))
}

func (s *server) GetQuestion(ctx context.Context, req *question.Request ) (*question.Question, error) {
	log.Println(fmt.Sprintf("Request: %g", req))
	qu := question.Question{
		Text: "How big is it?",
	}
	return &qu, nil
}

