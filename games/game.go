package games

type Game struct {
	Name string
}

var RegisteredGames []*Game

func RegisterGame(game *Game) {
	RegisteredGames = append(RegisteredGames, game)
}

func NewGame(name string) *Game {
	return &Game{
		Name: name,
	}
}
