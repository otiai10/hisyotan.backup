package actions

import (
	"hisyotan/conf"

	"github.com/otiai10/botw"
)

// Errors ...
type Errors struct {
	list []error
}

// Add ...
func (errs *Errors) Add(err error) {
	errs.list = append(errs.list, err)
}

func (errs *Errors) Error() string {
	msg := ""
	for _, e := range errs.list {
		msg += e.Error()
	}
	return msg
}

func filtered(status botw.Status) bool {
	// fmt.Printf("%+v\n\n", status)
	// 自分によるツイートはフィルターする
	if status.User.ScreenName == conf.Name() {
		return true
	}
	// 自分へのツイートでないならフィルターする
	if status.InReplyToScreenName != conf.Name() {
		return true
	}
	return false
}

// TODO: botwの機能に持って行きたい
// act.Reply(status, text) 的な
// 本当は act.Reply(text)の方がいい
func newReply(status botw.Status, text string) botw.Status {
	reply := botw.Status{}
	reply.InReplyToStatusIdStr = status.IdStr
	reply.InReplyToUserIdStr = status.User.IdStr
	reply.Text = "@" + status.User.ScreenName + " " + text
	return reply
}
