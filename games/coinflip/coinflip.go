package coinflip

import (
	"fmt"
	"math/rand"

	"github.com/Sirupsen/logrus"
	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

type CoinflipGame struct {
	games.BotsBasedGame

	board []string
}

func NewGame() (*CoinflipGame, error) {
	game := CoinflipGame{
		board: make([]string, 2),
	}
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
	for idx, bot := range g.Bots {
		reply, err := bot.SendMessage(bots.QuestionMessage{
			GameID:      gameID,
			Game:        g.Name(),
			Action:      "play-turn",
			Board:       g.board,
			PlayerIndex: idx,
		})
		if err != nil {
			return err
		}

		// parse reply
		if reply.Play != "ship" && reply.Play != "head" {
			return fmt.Errorf("player %d sent an invalid response: %q", idx, reply.Play)
		}
		g.board[idx] = reply.Play.(string)
	}

	// check if players chose the same value
	if g.board[0] == g.board[1] {
		return fmt.Errorf("the second player played the same value")
	}

	// trigger
	expectedValue := []string{"ship", "head"}[rand.Intn(2)]
	logrus.Warnf("Expected value: %q", expectedValue)

	for idx, value := range g.board {
		if value == expectedValue {
			logrus.Warnf("Player %d (%s) won", idx, g.Bots[idx].Name())
		}
	}

	return nil
}

func (g *CoinflipGame) Name() string {
	return "coinflip"
}
