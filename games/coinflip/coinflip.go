package coinflip

import "github.com/moul/bolosseum/games"

func init() {
	games.RegisterGame(games.NewGame("coinflip"))
}
