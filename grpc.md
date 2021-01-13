## Rest Web Services

[Great Go framework for this](https://github.com/gin-gonic/gin)

[Excellent Object-Relational Mapping](https://github.com/go-gorm/gorm)

## gRPC Services
```bash
 go get -u google.golang.org/grpc
```
[This is described here](https://gorm.io)

It also needs *protbuf*

```sh
brew install protobuf
```
Also install Install protoc-gen-go:

```sh
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Using Go modules in a project


A module is a collection of Go packages stored in a file tree with a go.mod file at its root. A good explanation can be found [here](https://blog.golang.org/using-go-modules).

The best approach is to start by creating a Github (or Gitlab) repository with minimal content and the to clone that in a local repository. This directory should be outside *GOPATH/src*. 

Create a new minimal Go file in that directory, for example: 

```go
func Hello() string {
    return "Hello, world."
}
```


Then create the go module with the command (for example) which shows the path to the root of the repository:

```bash
go mod init github.com/dgedge/quest
```
## Unit tests in Go

```go
package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
```

```bash
go test
```

