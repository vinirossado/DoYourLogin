# Stage 1: Build Go application
FROM golang:1.20-alpine3.17 AS builder

WORKDIR /app
COPY . .

RUN go build -o main ./source/cmd/main.go

# Stage 2: Final image
FROM mcr.microsoft.com/mssql/server:2019-latest

ENV ACCEPT_EULA=Y
ENV SA_PASSWORD=DoYourLoginEuOdeioReact2x
VOLUME ["/var/opt/mssql"]

COPY --from=builder /app/main /app/main

CMD /opt/mssql/bin/sqlservr & /app/main

EXPOSE 8025 1433