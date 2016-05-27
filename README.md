# bolosseum
:hocho: colosseum for bots

## Usage

```console
$ ./bolosseum run tictactoe http://showcase.m.42.am/bolosseum-tictactoe http://showcase.m.42.am/bolosseum-tictactoe
[0000] Initializing game "tictactoe"
[0000] Game: "tictactoe": &{{<nil> []} map["1-1":"" "1-2":"" "2-1":"" "2-2":"" "0-0":"" "0-1":"" "0-2":"" "1-0":"" "2-0":""]}
[0000] Getting bot "http://showcase.m.42.am/bolosseum-tictactoe"
[0000] Registering bot "http://showcase.m.42.am/bolosseum-tictactoe"
[0000] Getting bot "http://showcase.m.42.am/bolosseum-tictactoe"
[0000] Registering bot "http://showcase.m.42.am/bolosseum-tictactoe"
[0000] body> {"name":"moul-tictactoe","play":null,"error":null}
[0000] body> {"name":"moul-tictactoe","play":null,"error":null}
[0000] body> {"name":"","play":"2-2","error":null}
[0001] body> {"name":"","play":"1-1","error":null}
[0001] body> {"name":"","play":"0-2","error":null}
[0001] body> {"name":"","play":"1-2","error":null}
[0001] body> {"name":"","play":"1-0","error":null}
[0002] body> {"name":"","play":"0-1","error":null}
[0002] body> {"name":"","play":"2-1","error":null}
[0002] body> {"name":"","play":"2-0","error":null}
[0002] body> {"name":"","play":"0-0","error":null}
[0002] Draw
```

## Install

```console
$ go get github.com/moul/bolosseum/cmd/bolosseum
```

## License

MIT
