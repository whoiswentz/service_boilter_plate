FROM golang:alpine as builder

LABEL maintainer.name="Vinicios Wentz"
LABEL maintainer.email="vinicios@wentz.io"

RUN apk update \
    apk add ca-certificates

ADD . /build
WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

EXPOSE 8080

COPY --from=builder /build/main /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
CMD ["/app/main"]