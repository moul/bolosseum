package bots

type Bot interface {
	Name() string
	SetName(string)
	Path() string
	Start() error
	SendMessage(QuestionMessage) (*ReplyMessage, error)
}

type QuestionMessage struct {
	GameID  string `json:"game-id"`
	Action  string `json:"action"`
	Game    string `json:"game"`
	Players int    `json:"players"`
}

type ReplyMessage struct {
	Name string      `json:"name"`
	Play interface{} `json:"data"`
}
