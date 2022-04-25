FROM golang:1.18
WORKDIR /app
COPY main.go go.mod go.sum ./
RUN go mod download
EXPOSE 8080/tcp
RUN go get ./...
ENTRYPOINT [ "go", "run", "main.go" ]