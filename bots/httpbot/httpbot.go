package httpbot

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/moul/bolosseum/bots"
	"github.com/parnurzeal/gorequest"
)

type HttpBot struct {
	path string
	name string
}

func (b *HttpBot) Name() string {
	return b.name
}

func (b *HttpBot) SetName(name string) {
	b.name = name
}

func (b *HttpBot) Path() string {
	return b.path
}

func (b *HttpBot) SendMessage(msg bots.QuestionMessage) (*bots.ReplyMessage, error) {
	// marshal json
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	var query struct {
		Message string `json:"message"`
	}
	query.Message = string(data)
	resp, body, errs := gorequest.New().Get(b.path).Query(query).End()
	if len(errs) > 0 {
		return nil, fmt.Errorf("gorequest errs: %v", errs)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	logrus.Warnf("body> %s", strings.TrimSpace(body))
	// parse json
	var response bots.ReplyMessage
	if err = json.Unmarshal([]byte(body), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (b *HttpBot) Start() error {
	return nil
}

func NewBot(path string) (*HttpBot, error) {
	return &HttpBot{
		path: fmt.Sprintf("http://%s", path),
	}, nil
}
