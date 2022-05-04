FROM golang:1.15.3

WORKDIR /go/src/app
COPY ./ /go/src/app

RUN CGO_ENABLED=0 go build -o /go/src/app/serve serve.go

RUN go get -u -v bitbucket.org/liamstask/goose/cmd/goose

CMD sh start.sh

# FROM alpine:3.8

# WORKDIR /app

# COPY --from=builder /go/src/app/serve /serve

# CMD [ "/serve" ]