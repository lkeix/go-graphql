goose up
go run github.com/99designs/gqlgen init
CGO_ENABLED=0 go build -o /go/src/app/serve serve.go
./serve