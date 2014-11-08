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
