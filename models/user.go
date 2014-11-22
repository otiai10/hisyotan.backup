package models

import (
	"hisyotan/conf"

	"github.com/otiai10/rodeo"
)

var vaquero *rodeo.Vaquero

func init() {
	var err error
	vaquero, err = rodeo.NewVaquero(conf.RedisHost(), conf.RedisPort())
	if err != nil {
		panic(err)
	}
}

// User ...
type User struct {
	TwitterID  string
	ScreenName string
	Tasks      Tasks
}

// GetUserByTwitterID ...
func GetUserByTwitterID(twitterID string) (*User, error) {
	user := new(User)
	if err := vaquero.Cast(conf.RedisNS("user."+twitterID), user); err != nil {
		return nil, err
	}
	return user, nil
}

// NewUser ...
func NewUser(screenName, twitterID string) *User {
	return &User{
		TwitterID:  twitterID,
		ScreenName: screenName,
	}
}

// Save ...
func (user *User) Save() error {
	return vaquero.Store(conf.RedisNS("user."+user.TwitterID), user)
}
