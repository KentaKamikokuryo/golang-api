FROM golang:1.19.0-alpine

WORKDIR /usr/src/app
RUN go install github.com/cosmtrek/air@latest

ADD go.mod .
ADD go.sum .

RUN go mod tidy
RUN go mod download
ADD . .

EXPOSE 3000

# ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main