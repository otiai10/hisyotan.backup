package models

import (
	"testing"

	"hisyotan/conf"

	. "github.com/otiai10/mint"
)

func init() {
	conf.SetEnv("test")
}

func TestNewUser(t *testing.T) {
	user := NewUser("otiai10", "140021552")
	Expect(t, user).TypeOf("*models.User")
}

func TestUser_Save(t *testing.T) {
	user := NewUser("otiai10", "140021552")
	err := user.Save()
	Expect(t, err).ToBe(nil)
}

func TestGetUserByTwitterID(t *testing.T) {
	user := NewUser("otiai10", "140021552")
	Expect(t, user.Save()).ToBe(nil)

	found, err := GetUserByTwitterID("140021552")
	Expect(t, err).ToBe(nil)

	Expect(t, found.ScreenName).ToBe("otiai10")
}
