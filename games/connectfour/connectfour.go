package connectfour

import (
	"fmt"

	"github.com/moul/bolosseum/games"
)

type ConnectfourGame struct {
	games.Game
}

func NewGame() (*ConnectfourGame, error) {
	return &ConnectfourGame{
	//Game: games.Game{},
	}, nil
}

func (g *ConnectfourGame) CheckArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("You need to specify 2 bots")
	}
	return nil
}

func (g *ConnectfourGame) Run() error {
	fmt.Println("COUCOU")
	return nil
}

func (g *ConnectfourGame) Name() string {
	return "connectfour"
}
