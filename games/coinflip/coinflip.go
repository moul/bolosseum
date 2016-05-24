package coinflip

import (
	"fmt"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

type CoinflipGame struct {
	games.Game

	bots []bots.Bot
}

func NewGame() (*CoinflipGame, error) {
	return &CoinflipGame{
		bots: make([]bots.Bot, 0),
	}, nil
}

func (g *CoinflipGame) CheckArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("You need to specify 2 bots")
	}
	return nil
}

func (g *CoinflipGame) Run() error {
	// start bots
	for _, bot := range g.bots {
		if err := bot.Start(); err != nil {
			return err
		}
	}

	// send init message to bots
	for _, bot := range g.bots {
		reply, err := bot.SendMessage(bots.QuestionMessage{
			GameID:  "test",
			Action:  "init",
			Game:    g.Name(),
			Players: len(g.bots),
		})
		if err != nil {
			return err
		}

		// parse reply
		if reply.Name != "" {
			bot.SetName(reply.Name)
		}
	}

	// play
	// FIXME

	return nil
}

func (g *CoinflipGame) Name() string {
	return "coinflip"
}

func (g *CoinflipGame) RegisterBot(bot bots.Bot) {
	g.bots = append(g.bots, bot)
}
