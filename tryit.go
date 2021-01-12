package main

import (
    "fmt"
    "log"
)
import "github.com/dgedge/quest/internal/proto-files/domain"
import "google.golang.org/grpc"

func main() {
    sid := domain.Question{
        Text:   "Fred",
        Answer: "Bert",
    }

    conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalln(err)
    }
    defer conn.Close()

    fmt.Println(sid)


}