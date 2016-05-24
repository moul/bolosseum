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

// InitTurnBasedBots is an helper that starts and discovers connected bots
func InitTurnBasedBots(bots []Bot, gameName, gameID string) error {
	// start bots
	for _, bot := range bots {
		if err := bot.Start(); err != nil {
			return err
		}
	}

	// send init message to bots
	for _, bot := range bots {
		reply, err := bot.SendMessage(QuestionMessage{
			GameID:  gameID,
			Action:  "init",
			Game:    gameName,
			Players: len(bots),
		})
		if err != nil {
			return err
		}

		// parse reply
		if reply.Name != "" {
			bot.SetName(reply.Name)
		}
	}
	return nil
}
