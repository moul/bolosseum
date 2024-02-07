FROM golang:1.22
COPY . /go/src/github.com/moul/bolosseum
WORKDIR /go/src/github.com/moul/bolosseum
RUN go install -v ./cmd/bolosseum
CMD ["bolosseum", "server"]
EXPOSE 9000
