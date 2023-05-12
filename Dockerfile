FROM golang:1.20-alpine3.17
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main ./source/cmd/main.go
CMD ["/app/main"]
EXPOSE 8025

