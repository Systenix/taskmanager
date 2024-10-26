FROM golang:1.23.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/task_queue_system ./cmd/main.go

EXPOSE 8080

CMD ["./bin/task_queue_system"]