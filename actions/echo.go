package actions

import "fmt"
import "github.com/otiai10/botw"
import "hisyotan/models"

type Echo struct {
	*botw.Action
}

// See https://godoc.org/github.com/otiai10/twistream#Status to know Statsu structure
func (act *Echo) Match(status botw.Status) bool {
	if !IsToMe(status) {
		return false
	}
	words := model.NewWordSet(status.Text)
	if words.Has("/echo") {
		return true
	}
	return false
}

func (act *Echo) Execute(status botw.Status) {
	fmt.Printf("------------------ IN SAMPLE -------------------\n%+v\n", status)
	act.echo(status)
}

func (act *Echo) echo(status botw.Status) error {
	text := model.NewWordSet(status.Text).Filter(
		"@hisyotan",
		"/echo",
	).ToText(
		"@" + status.User.ScreenName,
	)
	reply := botw.Status{}
	reply.Text = text
	reply.InReplyToStatusIdStr = status.IdStr
	return act.Tweet(reply)
}
