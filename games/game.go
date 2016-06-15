package games

import "github.com/moul/bolosseum/bots"

type Game interface {
	Run(string) error
	Name() string
	CheckArgs([]string) error
	RegisterBot(bots.Bot)
	GetAsciiOutput() []byte
}

var RegisteredGames []Game

// RegisterGame append a game to the RegisteredGames list
func RegisterGame(game Game) {
	RegisteredGames = append(RegisteredGames, game)
}
