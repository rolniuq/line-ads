package messenger

import (
	"errors"
	"fmt"
	"line-ads/configs"
	ds "line-ads/internal/dial_settings"
	"line-ads/pkgs/http"

	"github.com/submodule-org/submodule.go/v2"
)

var MessengerMod = submodule.MakeModifiable[*Messenger](func(config *configs.Config) *Messenger {
	return NewMessenger(config, nil)
}, configs.ConfigMod)

type Messenger struct {
	config *configs.Config
	token  *string
}

func NewMessenger(config *configs.Config, token *string) *Messenger {
	return &Messenger{
		config: config,
		token:  token,
	}
}

type Message struct {
	Text string
	Type string
}

type MulticastRequest struct {
	Id       string
	To       []string
	Messages []Message
}

// https://developers.line.biz/en/reference/messaging-api/#send-multicast-message
func (s *Messenger) Multicast(req *MulticastRequest) (*SendMessageResponse, error) {
	if s.token == nil {
		return nil, errors.New("token is missing")
	}

	header := map[string][]string{
		"Authorization":    {fmt.Sprintf("Bearer %s", *s.token)},
		"Content-Type":     {"application/json"},
		"X-Line-Retry-Key": {req.Id},
	}

	body := ds.HttpBody{
		"to":       req.To,
		"messages": req.Messages,
	}

	hc := http.NewHttpClient("https://api.line.me/v2/bot/message/multicast", http.WithHeader(header), http.WithBody(body))
	res, err := http.Do[SendMessageResponse](hc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type SendMessageRequest struct {
	Id       string
	To       string
	Messages []Message
}

type SentMessage struct {
	Id         string `json:"id"`
	QuoteToken string `json:"quoteToken"`
}

type SendMessageResponse struct {
	SentMessages []SentMessage `json:"sentMessages"`
}

func (s *Messenger) SendMessage(req *SendMessageRequest) (*SendMessageResponse, error) {
	if s.token == nil {
		return nil, errors.New("token is missing")
	}

	header := map[string][]string{
		"Authorization":    {fmt.Sprintf("Bearer %s", *s.token)},
		"Content-Type":     {"application/json"},
		"X-Line-Retry-Key": {req.Id},
	}

	body := ds.HttpBody{
		"to":       req.To,
		"messages": req.Messages,
	}

	hc := http.NewHttpClient("https://api.line.me/v2/bot/message/push", http.WithHeader(header), http.WithBody(body))
	res, err := http.Do[SendMessageResponse](hc)
	if err != nil {
		return nil, err
	}

	return res, nil
}
