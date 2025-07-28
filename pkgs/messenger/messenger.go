package messenger

import (
	"line-ads/configs"
	ds "line-ads/internal/dial_settings"
	"line-ads/pkgs/http"

	"github.com/submodule-org/submodule.go/v2"
)

const longLive = "Vcb0SEpE25iDdfzZG1EfOzcBfRFEOd70DEZsS0TQ6FNJqaHOXl/LVPGQNpzKd99vqfQQ73ytIhlvOPnxxR/UyrFMRTO6GN0t9XCHspUSZBmar7mdFSLj/BxTFvnPiAj7mp4PlyySwFN4dhMThW4d9AdB04t89/1O/w1cDnyilFU="

var MessengerMod = submodule.Make[*Messenger](func(config *configs.Config) *Messenger {
	return NewMessenger(config)
}, configs.ConfigMod)

type Messenger struct {
	config *configs.Config
}

func NewMessenger(config *configs.Config) *Messenger {
	return &Messenger{config}
}

type MulticastMessage struct {
	Text string
	Type string
}

type MulticastRequest struct {
	To      []string
	Message []MulticastMessage
}

// https://developers.line.biz/en/reference/messaging-api/#send-multicast-message
func (s *Messenger) Multicast(req *MulticastRequest) error {
	return nil
}

type SentMessage struct {
	Id         string `json:"id"`
	QuoteToken string `json:"quoteToken"`
}

type SendMessageResponse struct {
	SentMessages []SentMessage `json:"sentMessages"`
}

func (s *Messenger) SendMessage() (*SendMessageResponse, error) {
	header := map[string][]string{}

	hc := http.NewHttpClient("https://api.line.me/v2/bot/message/push", http.WithHeader(header), http.WithBody(ds.HttpBody{}))
	res, err := http.Do[SendMessageResponse](hc)
	if err != nil {
		return nil, err
	}

	return res, nil
}
