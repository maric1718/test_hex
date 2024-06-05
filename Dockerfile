FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

RUN go env -w GO111MODULE=auto
RUN go mod tidy 
RUN go install github.com/githubnemo/CompileDaemon@master

EXPOSE 8080

ENTRYPOINT CompileDaemon -exclude-dir=.git -graceful-kill=true -build="go build -o ./pos-app ./cmd/main.go" -command="./pos-app"


