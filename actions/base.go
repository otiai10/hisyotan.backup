package actions

import (
	"hisyotan/conf"

	"github.com/otiai10/botw"
)

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
func newReplay(status botw.Status, text string) botw.Status {
	reply := botw.Status{}
	reply.InReplyToStatusIdStr = status.IdStr
	reply.InReplyToUserIdStr = status.User.IdStr
	reply.Text = "@" + status.User.ScreenName + " " + text
	return reply
}
