package coinflip

import (
	"fmt"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

type CoinflipGame struct {
	games.BotsBasedGame
}

func NewGame() (*CoinflipGame, error) {
	game := CoinflipGame{}
	game.Bots = make([]bots.Bot, 0)
	return &game, nil
}

func (g *CoinflipGame) CheckArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("You need to specify 2 bots")
	}
	return nil
}

func (g *CoinflipGame) Run(gameID string) error {
	if err := bots.InitTurnBasedBots(g.Bots, g.Name(), gameID); err != nil {
		return err
	}

	// play
	// FIXME

	return nil
}

func (g *CoinflipGame) Name() string {
	return "coinflip"
}
