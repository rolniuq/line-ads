package notifier

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type service1 struct {
	token1 *string
}

func (s *service1) OnNotify(token string) {
	s.token1 = &token
}

type service2 struct {
	token2 *string
}

func (s *service2) OnNotify(token string) {
	s.token2 = &token
}

func Test_Notifier(t *testing.T) {
	notify := NotifierMod.Resolve()

	s1 := &service1{}
	s2 := &service2{}

	notify.Register("access_token", s1)
	notify.Register("access_token", s2)

	notify.Notify("new token")

	require.True(t, *s1.token1 == "new token")
	require.True(t, *s2.token2 == "new token")
}
