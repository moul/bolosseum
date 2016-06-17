package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/moul/bolosseum"
	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/bots/filebot"
	"github.com/moul/bolosseum/bots/httpbot"
	"github.com/moul/bolosseum/bots/stupidbot"
	"github.com/moul/bolosseum/games"
	"github.com/moul/bolosseum/games/coinflip"
	"github.com/moul/bolosseum/games/connectfour"
	"github.com/moul/bolosseum/games/russianbullet"
	"github.com/moul/bolosseum/games/tictactoe"
	"github.com/moul/bolosseum/stupid-ias"
	"github.com/moul/bolosseum/stupid-ias/coinflip"
	"github.com/moul/bolosseum/stupid-ias/connectfour"
	"github.com/moul/bolosseum/stupid-ias/tictactoe"
	"github.com/urfave/cli"
)

var availableGames = []string{
	"coinflip",
	"connectfour",
	"russianbullet",
	"tictactoe",
}

func getGame(gameName string) (games.Game, error) {
	switch gameName {
	case "coinflip":
		return coinflip.NewGame()
	case "connectfour":
		return connectfour.NewGame()
	case "russianbullet":
		return russianbullet.NewGame()
	case "tictactoe":
		return tictactoe.NewGame()
	default:
		return nil, fmt.Errorf("unknown game %q", gameName)
	}
}

func getStupidIA(iaPath string) (stupidias.StupidIA, error) {
	logrus.Warnf("Getting stupid IA %q", iaPath)
	switch iaPath {
	case "connectfour":
		return stupidconnectfour.NewIA()
	case "coinflip":
		return stupidcoinflip.NewIA()
	case "tictactoe":
		return stupidtictactoe.NewIA()
	default:
		return nil, fmt.Errorf("unknown stupid IA %q", iaPath)
	}
}

func getBot(botPath string, game games.Game) (bots.Bot, error) {
	logrus.Warnf("Getting bot %q", botPath)

	if botPath == "stupid" {
		botPath = fmt.Sprintf("stupid://%s", game.Name())
	}

	splt := strings.Split(botPath, "://")
	if len(splt) != 2 {
		return nil, fmt.Errorf("invalid bot path")
	}

	scheme := splt[0]
	path := splt[1]

	switch scheme {
	case "file":
		return filebot.NewBot(path)
	case "http+get":
		return httpbot.NewBot(path, "GET", "http")
	case "http+post", "http":
		return httpbot.NewBot(path, "POST", "http")
	case "https+get":
		return httpbot.NewBot(path, "GET", "https")
	case "https+post", "https":
		return httpbot.NewBot(path, "POST", "https")
	case "stupid":
		ia, err := getStupidIA(path)
		if err != nil {
			return nil, err
		}
		return stupidbot.NewStupidBot(path, ia)
	default:
		return nil, fmt.Errorf("invalid bot scheme: %q (%q)", scheme, path)
	}
}

func main() {
	// seed random
	rand.Seed(time.Now().UTC().UnixNano())

	// configure CLI
	app := cli.NewApp()
	app.Name = "bolosseum"
	app.Usage = "colosseum for bots"
	app.Version = bolosseum.VERSION

	app.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "Start a battle",
			Action: run,
		},
		{
			Name:   "list-games",
			Usage:  "List games",
			Action: listGames,
		},
		{
			Name:   "server",
			Usage:  "Start a bolosseum web server",
			Action: server,
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("%v", err)
	}
}

var indexHTML = `<html>
  <head>
    <title>Bolosseum</title>
  </head>
  <body>
    <h1>Bolosseum</h1>
  </body>
</html>`

func server(c *cli.Context) error {
	r := gin.Default()
	r.LoadHTMLGlob("web/*")
	r.GET("/", func(c *gin.Context) {
		//c.Header("Content-Type", "text/html")
		//c.String(http.StatusOK, indexHTML)
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
	return r.Run(":9000")
}

func listGames(c *cli.Context) error {
	fmt.Println("Games:")
	for _, game := range availableGames {
		fmt.Printf("- %s\n", game)
	}
	return nil
}

func run(c *cli.Context) error {
	args := c.Args()
	if len(args) < 1 {
		return cli.NewExitError("You need to specify the game", -1)
	}

	// initialize game
	logrus.Warnf("Initializing game %q", args[0])
	game, err := getGame(args[0])
	if err != nil {
		return cli.NewExitError(fmt.Sprintf("No such game %q", args[0]), -1)
	}
	logrus.Warnf("Game: %q: %q", game.Name(), game)

	if err = game.CheckArgs(args[1:]); err != nil {
		return cli.NewExitError(fmt.Sprintf("%v", err), -1)
	}

	// initialize bots
	hasError := false
	for _, botPath := range args[1:] {
		bot, err := getBot(botPath, game)
		if err != nil {
			hasError = true
			logrus.Errorf("Failed to initialize bot %q", bot)
		} else {
			logrus.Warnf("Registering bot %q", bot.Path())
			game.RegisterBot(bot)
		}
	}
	if hasError {
		return cli.NewExitError("Invalid bots", -1)
	}

	// run
	if err = game.Run("gameid"); err != nil {
		logrus.Errorf("Run error: %v", err)
	}

	// print ascii output
	if output := game.GetAsciiOutput(); len(output) > 0 {
		fmt.Printf("%s", output)
	}

	return nil
}
