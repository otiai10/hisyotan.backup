package actions

import (
	"fmt"
	"hisyotan/conf"
	"hisyotan/models"

	"github.com/otiai10/botw"
)

// Add ...
type Add struct {
	*botw.Action
}

// Match ...
func (act *Add) Match(status botw.Status) bool {
	if filtered(status) {
		return false
	}
	// TODO: check user
	return models.ParseText(status.Text).Has(conf.Name(true), "add")
}

// Execute ...
func (act *Add) Execute(status botw.Status) {
	elms := models.ParseText(status.Text)
	elms.Remove(conf.Name(true), "add")
	// TODO: add tasks

	reply := botw.Status{}
	reply.Text = elms.ToText(
		"@"+status.User.ScreenName,
		"を追加しました！",
	)
	reply.InReplyToStatusIdStr = status.IdStr

	fmt.Println(act.Tweet(reply))
}
