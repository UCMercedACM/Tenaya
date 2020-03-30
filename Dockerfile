FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go install

EXPOSE 4203

CMD ["go", "run", "main.go"]
