package notifier

import "sync"

type notifier struct {
	m sync.Map
}

func NewNotifier() *notifier {
	return &notifier{
		m: sync.Map{},
	}
}

func (n *notifier) Register(k, v any) {
	n.m.Store(k, v)
}

func (n *notifier) Notify() {}
