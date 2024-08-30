FROM golang:1.22-alpine

COPY . /go/src/app

WORKDIR /go/src/app/cmd/

RUN go build -o app main.go

COPY .env .

EXPOSE 6000

CMD ["./app"]