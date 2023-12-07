package goutil

import (
	"testing"
)

func TestStrFormatHumpToUnder(t *testing.T) {
	keys := []string{
		"userName",
		"_user_Name",
		"UUUsername",
		"_User_name",
	}
	for _, key := range keys {
		t.Log(StrFormatHumpToUnder(key))
	}
}

func TestStrFormatUnderToHump(t *testing.T) {
	keys := []string{
		"user_name",
		"userName",
		"_user_Name",
		"__user_Name",
		"Username",
		"_User_1name",
	}
	for _, key := range keys {
		t.Log(StrFormatUnderToHump(key))
	}
}
