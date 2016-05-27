# bolosseum
:hocho: colosseum for bots

## Usage

local battle

```console
$ bolosseum run tictactoe file://./test/tictactoe file://./test/tictactoe
[0000] Initializing game "tictactoe"
[0000] Game: "tictactoe": &{{<nil> []} map["0-0":"" "0-1":"" "0-2":"" "1-0":"" "1-1":"" "2-0":"" "2-2":"" "1-2":"" "2-1":""]}
[0000] Getting bot "file://./test/tictactoe"
[0000] Registering bot "./test/tictactoe"
[0000] Getting bot "file://./test/tictactoe"
[0000] Registering bot "./test/tictactoe"
[0000] ./test/tictactoe << {"game-id":"gameid","action":"init","game":"tictactoe","players":2,"board":null,"you":null,"player-index":0}
[0000] body> {"name":"moul-tictactoe","play":null,"error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"init","game":"tictactoe","players":2,"board":null,"you":null,"player-index":1}
[0000] body> {"name":"moul-tictactoe","play":null,"error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"","1-0":"","1-1":"","1-2":"","2-0":"","2-1":"","2-2":""},"you":"X","player-index":0}
[0000] body> {"name":"","play":"2-2","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"","1-0":"","1-1":"","1-2":"","2-0":"","2-1":"","2-2":"X"},"you":"O","player-index":1}
[0000] body> {"name":"","play":"1-1","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"","1-0":"","1-1":"O","1-2":"","2-0":"","2-1":"","2-2":"X"},"you":"X","player-index":0}
[0000] body> {"name":"","play":"0-2","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"X","1-0":"","1-1":"O","1-2":"","2-0":"","2-1":"","2-2":"X"},"you":"O","player-index":1}
[0000] body> {"name":"","play":"1-2","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"X","1-0":"","1-1":"O","1-2":"O","2-0":"","2-1":"","2-2":"X"},"you":"X","player-index":0}
[0000] body> {"name":"","play":"1-0","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"X","1-0":"X","1-1":"O","1-2":"O","2-0":"","2-1":"","2-2":"X"},"you":"O","player-index":1}
[0000] body> {"name":"","play":"0-1","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"O","0-2":"X","1-0":"X","1-1":"O","1-2":"O","2-0":"","2-1":"","2-2":"X"},"you":"X","player-index":0}
[0000] body> {"name":"","play":"2-1","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"O","0-2":"X","1-0":"X","1-1":"O","1-2":"O","2-0":"","2-1":"X","2-2":"X"},"you":"O","player-index":1}
[0000] body> {"name":"","play":"2-0","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"O","0-2":"X","1-0":"X","1-1":"O","1-2":"O","2-0":"O","2-1":"X","2-2":"X"},"you":"X","player-index":0}
[0000] body> {"name":"","play":"0-0","error":null}
[0000] Draw
```

remote battle

```console
$ bolosseum run tictactoe http://showcase.m.42.am/bolosseum-tictactoe http://showcase.m.42.am/bolosseum-tictactoe
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
