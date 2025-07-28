package notifier

import (
	"sync"

	"github.com/submodule-org/submodule.go/v2"
)

var NotifierMod = submodule.Make[*notifier](func() *notifier {
	return NewNotifier()
})

type Listener interface {
	OnNotify(token string)
}

type notifier struct {
	m *sync.Map
}

func NewNotifier() *notifier {
	return &notifier{
		m: &sync.Map{},
	}
}

func (n *notifier) Register(k string, l Listener) {
	n.m.Store(k, l)
}

func (n *notifier) Notify(token string) {
	n.m.Range(func(key, value any) bool {
		if l, ok := value.(Listener); ok {
			l.OnNotify(token)
		}

		return true
	})
}

func (n *notifier) Unregister(key string) {
	n.m.Delete(key)
}
