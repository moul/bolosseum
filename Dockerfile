FROM golang:1.6
RUN go get github.com/Masterminds/glide
COPY glide.lock glide.yaml /go/src/github.com/moul/bolosseum/
WORKDIR /go/src/github.com/moul/bolosseum
RUN glide install
COPY . /go/src/github.com/moul/bolosseum
RUN go install -v ./cmd/bolosseum
CMD ["bolosseum", "server"]
EXPOSE 9000
