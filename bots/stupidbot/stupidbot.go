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
	logrus.Warnf("bot-%d >> %v", msg.PlayerIndex, msg)
	var reply *bots.ReplyMessage
	switch msg.Action {
	case "init":
		reply = b.ia.Init(msg)
	case "play-turn":
		reply = b.ia.PlayTurn(msg)
	default:
		return nil, fmt.Errorf("Unknown action %q", msg.Action)
	}

	logrus.Warnf("bot-%d << %v", msg.PlayerIndex, *reply)
	return reply, nil
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
