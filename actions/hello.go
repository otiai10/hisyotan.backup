package actions

import (
	"fmt"
	"hisyotan/conf"
	"hisyotan/models"

	"github.com/otiai10/botw"
)

// Hello ...
type Hello struct {
	*botw.Action
}

// Match ...
func (act *Hello) Match(status botw.Status) bool {
	if filtered(status) {
		return false
	}
	return models.ParseText(status.Text).Has(conf.Name(true), "hello")
}

// Execute ...
func (act *Hello) Execute(status botw.Status) {
	_, err := models.GetUserByTwitterID(status.User.IdStr)
	if err == nil {
		// found
		// TODO: feedback
		return // do nothing
	}
	user := models.NewUser(status.User.ScreenName, status.User.IdStr)
	if err := user.Save(); err != nil {
		// TODO: do nothing
	}
	fmt.Println("Hello: ", act.Tweet(newReply(status, "よろしくお願いします！")))
}
