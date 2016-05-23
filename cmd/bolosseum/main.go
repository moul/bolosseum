package main

import (
	"fmt"
	"os"

	// games
	_ "github.com/moul/bolosseum/games/coinflip"

	"github.com/Sirupsen/logrus"
	"github.com/moul/bolosseum"
	"github.com/moul/bolosseum/games"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "bolosseum"
	app.Usage = "colosseum for bots"
	app.Version = bolosseum.VERSION

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Start a battle",
			Action: func(c *cli.Context) error {
				args := c.Args()
				if len(args) < 1 {
					return cli.NewExitError("You need to specify the game", -1)
				}

				gameName := args[0]
				logrus.Warnf("Initializing game %q", gameName)

				found := false
				var game *games.Game
				for _, entry := range games.RegisteredGames {
					if entry.Name == gameName {
						game = entry
						found = true
						break
					}
				}
				if !found {
					return cli.NewExitError(fmt.Sprintf("No such game %q", gameName), -1)
				}
				logrus.Warnf("Game: %q", game)

				// temporarily hardcoded check
				if len(args) < 3 {
					return cli.NewExitError("You need to specify two bots", -1)
				}
				bots := args[1:3]
				logrus.Warnf("Bots: %q", bots)

				return nil
			},
		},
	}
	app.Run(os.Args)
}
