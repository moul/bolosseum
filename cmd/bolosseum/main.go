package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/moul/bolosseum"
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

				game := args[0]
				logrus.Warnf("Initializing game %q", game)

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
