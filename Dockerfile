FROM golang:1.18-alpine3.15 AS builder

COPY . /mentoring-app/
WORKDIR /mentoring-app/

RUN go mod download
RUN go build -o ./bin/server main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /mentoring-app/bin/server .
COPY --from=0 /mentoring-app/config.json .
COPY --from=0 /mentoring-app/.env .

EXPOSE 8080

CMD ["./server"]