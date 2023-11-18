# Build stage
FROM golang:1.21.3 as build
WORKDIR /app
ADD . /app
RUN go mod download
RUN CGO_ENABLED=0 go build -o bin/grpc-calculator server/main.go

# Production stage
FROM alpine:latest as production
WORKDIR /root/
COPY --from=build /app/bin/grpc-calculator .
EXPOSE 9200
CMD ./grpc-calculator
