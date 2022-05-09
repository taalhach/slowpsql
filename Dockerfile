FROM golang:alpine as builder
ENV GO111MODULE=on
LABEL maintainer="Muhammad Talha <talhach891@gmail.com>"
RUN apk update && apk add --no-cache git bash make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make
FROM alpine:latest
WORKDIR /root/
COPY config-files/app-configs-docker.ini  .
ENV SETTINGS=/root/app-configs-docker.ini
COPY --from=builder /app/bin/app .
EXPOSE 3000
LABEL Name=app-boiler-plate Version=0.0.1
CMD ["./app", "serve_api"]
