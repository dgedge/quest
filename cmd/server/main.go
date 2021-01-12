package main

import (
	"context"
	"fmt"
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

	log.Fatalln(srv.Serve(lis))
}

func (s *server) QuestionService(ctx context.Context, question.Request ) (*question.Question, error) {
	log.Println(fmt.Sprintf("Request: %g", request))

	return &question.CreditResponse{Confirmation: fmt.Sprintf("Credited %g", request.GetAmount())}, nil
}

