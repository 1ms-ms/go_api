# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app
ENV GOPROXY=https://proxy.golang.org

COPY main.go /app/
RUN go mod init golang
RUN go mod tidy

EXPOSE 8080
EXPOSE 1123123123
ENTRYPOINT [ "go", "run", "main.go" ]