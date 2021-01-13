package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)
import "github.com/dgedge/quest/pkg/proto/question"


func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := question.NewQuestionServiceClient(conn)
	req := question.Request{
		Qno: 1,
	}
	ctx := context.Background()
	res, err := client.GetQuestion(ctx, &req)
	fmt.Println(res.Text)

}