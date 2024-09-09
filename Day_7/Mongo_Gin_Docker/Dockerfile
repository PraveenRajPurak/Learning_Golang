FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY .env .env
RUN apk --no-cache add ca-certificates
EXPOSE 10005
CMD ["./main"]