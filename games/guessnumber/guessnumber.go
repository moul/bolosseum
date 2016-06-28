package guessnumber

import (
	"fmt"
	"math/rand"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

type round struct {
	Number int    `json:"number"`
	Answer string `json:"answer"`
}

type GuessnumberGame struct {
	games.BotsBasedGame

	board []round
}

func NewGame() (*GuessnumberGame, error) {
	game := GuessnumberGame{}
	game.Bots = make([]bots.Bot, 0)
	game.board = make([]round, 0)
	return &game, nil
}

func (g *GuessnumberGame) CheckArgs(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("You need to specify at least 1 bot")
	}
	return nil
}

func (g *GuessnumberGame) Run(gameID string, steps chan games.GameStep) error {
	if err := bots.InitTurnBasedBots(g.Bots, g.Name(), gameID); err != nil {
		return err
	}

	// think about a number
	targetNumber := rand.Intn(101) // target is between 0 and 100

	// play
	for turn := 0; ; turn++ {
		botIndex := turn % len(g.Bots)
		bot := g.Bots[botIndex]

		question := bots.QuestionMessage{
			GameID:      gameID,
			Players:     len(g.Bots),
			Board:       g.board,
			Game:        g.Name(),
			Action:      "play-turn",
			PlayerIndex: botIndex,
		}
		steps <- games.GameStep{QuestionMessage: &question}
		reply, err := bot.SendMessage(question)
		if err != nil {
			return err
		}
		reply.PlayerIndex = botIndex
		steps <- games.GameStep{ReplyMessage: reply}

		if reply.Play == targetNumber {
			steps <- games.GameStep{Winner: g.Bots[botIndex]}
			return nil
		}

		var answer string
		if targetNumber > reply.Play.(int) {
			answer = ">"
		} else {
			answer = "<"
		}

		g.board = append(g.board, round{
			Number: reply.Play.(int),
			Answer: answer,
		})
	}
}

func (g *GuessnumberGame) Name() string {
	return "guessnumber"
}

func (g *GuessnumberGame) GetAsciiOutput() []byte {
	return nil
}
