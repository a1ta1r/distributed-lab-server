FROM golang:alpine

ADD . /go/src/github.com/a1ta1r/distributed-lab-server
WORKDIR /go/src/github.com/a1ta1r/distributed-lab-server

ENV PORT=3001
ENV GIN_MODE=release

CMD ["go", "run", "cmd/server/server.go"]