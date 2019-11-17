FROM golang:latest

LABEL maintainer.name="Vinicios Wentz" \
	  maintainer.email="vinicios@wentz.io"

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .

EXPOSE 8080
CMD ["/app/main"]