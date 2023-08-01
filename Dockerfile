# Stage 1: Build Go application
FROM golang:1.20-alpine3.17 AS builder

WORKDIR /app
COPY . .

RUN go build -o main ./source/cmd/main.go

COPY --from=builder /app/main /app/main

CMD /app/main

EXPOSE 8025