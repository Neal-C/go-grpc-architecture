# Build stage
FROM golang:1.21.3 as build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go-grpc-calculator

# Production stage
FROM alpine:latest as production
WORKDIR /root/
COPY --from=build /go-grpc-calculator ./
CMD ["./go-grpc-calculator"]
