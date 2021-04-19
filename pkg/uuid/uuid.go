package uuid

import (
	"github.com/lithammer/shortuuid/v3"
	"github.com/teris-io/shortid"
)

func NewShortUUID() string {
	return shortuuid.New()
}
func NewShortID() string {
	_shortid, _ := shortid.Generate()
	return _shortid
}
