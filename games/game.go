package games

import "github.com/moul/bolosseum/bots"

type Game interface {
	Run() error
	Name() string
	CheckArgs([]string) error
	RegisterBot(bots.Bot)
}

var RegisteredGames []Game

func RegisterGame(game Game) {
	RegisteredGames = append(RegisteredGames, game)
}
