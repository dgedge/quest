.PHONY: compile
compile: ## Compile the proto file.
	protoc -I pkg/proto/question/ pkg/proto/question/question.proto --go-grpc_out=pkg/proto/question/

.PHONY: server
server: ## Build and run server.
	go build -race -ldflags "-s -w" -o cmd/server cmd/server/main.go
	cmd/server

.PHONY: client
client: ## Build and run client.
	go build -race -ldflags "-s -w" -o cmd/client cmd/client/main.go
	cmd/client

