package goutil

import "testing"

func TestStrGenerateRandom(t *testing.T) {
	allowedChar := []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	for i := 0; i < 10; i++ {
		t.Log(StrGenerateRandom(allowedChar, 8))
	}
}
