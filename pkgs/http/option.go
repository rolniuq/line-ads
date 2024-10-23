package http

import "line-ads/internal"

type Option interface {
	Apply(*internal.DialSettings)
}
