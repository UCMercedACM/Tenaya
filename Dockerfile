FROM golang:latest

LABEL maintainer="UCM ACM Chapter"
LABEL maintainer.email="acm@ucmerced.edu"

WORKDIR /go/src/app

COPY . .

RUN go install

EXPOSE 4203

CMD ["go", "run", "main.go", "service.go"]

HEALTHCHECK --interval=5m --timeout=3s \
    CMD curl -f http://localhost/ || exit 1
