package actions

import "github.com/otiai10/botw"
import "strings"
import "hisyotan/models"

func ToWords(text string) model.Words {
	return model.NewWordSet(text)
}

func IsToMe(status botw.Status) bool {
	return status.InReplyToScreenName == "hisyotan"
}

func FilterMe(text string, filterWords ...string) string {
	res := strings.Replace(text, "@hisyotan", "", -1)
	for _, w := range filterWords {
		res = strings.Replace(res, w, "", -1)
	}
	return res
}
