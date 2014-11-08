package actions

import (
	"fmt"
	"hisyotan/conf"
	"hisyotan/models"

	"github.com/otiai10/botw"
)

// Echo ...
type Echo struct {
	*botw.Action
}

// Match ...
func (act *Echo) Match(status botw.Status) bool {
	if filtered(status) {
		return false
	}
	return models.ParseText(status.Text).Has(conf.Name(true), "echo")
}

// Execute ...
func (act *Echo) Execute(status botw.Status) {
	elms := models.ParseText(status.Text)
	elms.Remove(conf.Name(true), "echo")
	reply := botw.Status{}
	reply.Text = elms.ToText("@"+status.User.ScreenName, "You said!")
	reply.InReplyToStatusIdStr = status.IdStr

	fmt.Println(act.Tweet(reply))
}
