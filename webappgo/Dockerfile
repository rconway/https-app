FROM golang:latest

WORKDIR /go/src/github.com/rconway/webappgo
COPY ./webappgo/. .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000/tcp

CMD ["webappgo"]
