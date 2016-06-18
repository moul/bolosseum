FROM golang:1.6-onbuild
RUN go get -d -v ./...
RUN go install -v ./cmd/bolosseum
CMD ["bolosseum", "server"]
EXPOSE 9000