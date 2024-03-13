# 1 Stage Building App
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o virtual-diary cmd/virtual-diary/main.go

# 2 Stage Run app
FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/virtual-diary .
