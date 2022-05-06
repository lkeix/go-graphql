FROM golang:1.16.15

WORKDIR /go/src/app
COPY ./ /go/src/app

ENV POSTGRES_HOST "db"
ENV POSTGRES_PORT "5432"
ENV POSTGRES_USER "user"
ENV POSTGRES_PASSWORD "password"
ENV POSTGRES_DBNAME "go-graphql-sandbox"

# RUN CGO_ENABLED=0 go build -o /go/src/app/serve serve.go
RUN go get -u -v bitbucket.org/liamstask/goose/cmd/goose
RUN go get github.com/cosmtrek/air

CMD sh start.sh

# FROM alpine:3.8

# WORKDIR /app

# COPY --from=builder /go/src/app/serve /serve

# CMD [ "/serve" ]