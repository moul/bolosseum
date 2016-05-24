package russianbullet

import (
	"fmt"
	"math/rand"

	"github.com/Sirupsen/logrus"
	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

type RussianbulletGame struct {
	games.BotsBasedGame
}

func NewGame() (*RussianbulletGame, error) {
	game := RussianbulletGame{}
	game.Bots = make([]bots.Bot, 0)
	return &game, nil
}

func (g *RussianbulletGame) CheckArgs(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("You need to specify at least 1 bot")
	}
	return nil
}

func (g *RussianbulletGame) Run(gameID string) error {
	if err := bots.InitTurnBasedBots(g.Bots, g.Name(), gameID); err != nil {
		return err
	}

	// play
	bulletIndex := rand.Intn(6) // 6 slots in the revolver
	for i := 0; i <= bulletIndex; i++ {
		idx := i % len(g.Bots)
		bot := g.Bots[idx]

		reply, err := bot.SendMessage(bots.QuestionMessage{
			GameID:      gameID,
			Players:     len(g.Bots),
			Game:        g.Name(),
			Action:      "play-turn",
			PlayerIndex: idx,
		})
		if err != nil {
			return err
		}

		if reply.Play != "click !" {
			return fmt.Errorf("Invalid bot input: %v", reply.Play)
		}
		if i == bulletIndex {
			logrus.Warnf("Player %d (%s) was killed", idx, bot.Name())
		}
	}

	logrus.Warnf("Draw")
	return nil
}

func (g *RussianbulletGame) Name() string {
	return "russianbullet"
}
