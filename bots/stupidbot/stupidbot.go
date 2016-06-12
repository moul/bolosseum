package stupidbot

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/moul/bolosseum/bots"
	"github.com/moul/bolosseum/stupid-ias"
)

type StupidBot struct {
	path string
	name string
	ia   stupidias.StupidIA
}

func (b *StupidBot) Name() string {
	return b.name
}

func (b *StupidBot) SetName(name string) {
	b.name = name
}

func (b *StupidBot) Path() string {
	return b.path
}

func (b *StupidBot) SendMessage(msg bots.QuestionMessage) (*bots.ReplyMessage, error) {
	logrus.Warnf("SendMessage: %v", msg)
	switch msg.Action {
	case "init":
		return b.ia.Init(msg), nil
	case "play-turn":
		return b.ia.PlayTurn(msg), nil
	default:
		return nil, fmt.Errorf("Unknown action %q", msg.Action)
	}
}

func (b *StupidBot) Start() error {
	return nil
}

func NewStupidBot(path string, ia stupidias.StupidIA) (*StupidBot, error) {
	return &StupidBot{
		path: path,
		name: path,
		ia:   ia,
	}, nil
}
