package helpers

import (
	"github.com/lithammer/shortuuid/v3"
)

func GenerateUUID() string {
	return shortuuid.New()
}
