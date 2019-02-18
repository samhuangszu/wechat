package util

import "github.com/samhuangszu/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
