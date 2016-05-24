SOURCE :=	$(shell find . -name "*.go")


bolosseum: $(SOURCE)
	go build -o ./bolosseum ./cmd/bolosseum/main.go


test: bolosseum
	./bolosseum run coinflip file://./test/coinflip.sh file://./test/coinflip.sh
