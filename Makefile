SOURCE :=	$(shell find . -name "*.go")


bolosseum: $(SOURCE)
	go build -o ./bolosseum ./cmd/bolosseum/main.go


.PHONY: test-coinflip
test-coinflip: bolosseum
	./bolosseum run coinflip file://./test/coinflip.sh file://./test/coinflip.sh


.PHONY: test-tictactoe
test-tictactoe: bolosseum test/tictactoe
	./bolosseum run tictactoe file://./test/tictactoe file://./test/tictactoe


test/tictactoe:
	go get github.com/moul/tictactoe/cmd/tictactoe-bolosseum
	ln -s $(GOPATH)/bin/tictactoe-bolosseum $@
