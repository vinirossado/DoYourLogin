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


#Install MongoDB

RUN apk add --no-cache mongodb

RUN mkdir -p /data/db

#Install Grafana
RUN apk add --no-cache ca-certificates
RUN wget -O /etc/apk/keys/grafana.rsa.pub https://packages.grafana.com/gpg.key
RUN echo "https://packages.grafana.com/oss/deb stable main" | tee -a /etc/apk/repositories
RUN apk update && apk add grafana


EXPOSE 8025 1433 3000 27017

CMD /opt/mssql/bin/sqlservr & mongod --bind_ip_all & /app/main  & grafana-server --config=/etc/grafana/grafana.ini --homepath=/usr/share/grafana

