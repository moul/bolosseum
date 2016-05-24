package tictactoe

import (
	"fmt"

	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/games"
)

type TictactoeGame struct {
	games.BotsBasedGame

	board map[string]string
}

func NewGame() (*TictactoeGame, error) {
	game := TictactoeGame{}
	game.Bots = make([]bots.Bot, 0)
	game.board = make(map[string]string, 9)

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			game.board[fmt.Sprintf("%d-%d", y, x)] = ""
		}
	}

	return &game, nil
}

func (g *TictactoeGame) CheckArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("You need to specify 2 bots")
	}
	return nil
}

func (g *TictactoeGame) Run(gameID string) error {
	if err := bots.InitTurnBasedBots(g.Bots, g.Name(), gameID); err != nil {
		return err
	}

	pieces := []string{"X", "O"}

	// play
	for turn := 0; turn < 9; turn++ {
		idx := turn % 2
		bot := g.Bots[idx]
		piece := pieces[idx]

		reply, err := bot.SendMessage(bots.QuestionMessage{
			GameID:      gameID,
			Game:        g.Name(),
			Action:      "play-turn",
			Board:       g.board,
			You:         piece,
			PlayerIndex: idx,
		})
		if err != nil {
			return err
		}

		g.board[reply.Play.(string)] = piece
		// FIXME: check for finished game
	}

	return nil
}

func (g *TictactoeGame) Name() string {
	return "tictactoe"
}
