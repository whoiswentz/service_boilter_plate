FROM golang:1.13.4-alpine3.10

RUN mkdir /app
ADD . /app
WORKDIR /app

ENV SERVICE_PORT=:8080
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD="eU_\@?S8tD:QnhF*&&mA.Ve1Q.Sv5JD"
ENV DB_DATABASE=supplier

RUN CGO_ENABLED=0 go build -o main .

EXPOSE 8080:8080/tcp
CMD ["/app/main"]