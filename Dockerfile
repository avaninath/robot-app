# Build stage
FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN cd cmd && go build -o robotapp main.go

# Run stage
FROM alpine:3.18.4
WORKDIR /app
COPY --from=builder /app/cmd/robotapp .
CMD [ "/app/robotapp" ]
