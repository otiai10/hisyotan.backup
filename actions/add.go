package actions

import (
	"fmt"
	"hisyotan/conf"
	"hisyotan/models"
	"log"

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

	user, err := models.GetUserByTwitterID(status.User.IdStr)
	if err != nil {
		act.Tweet(newReply(status, "だれですかあなた"))
		return
	}

	errs := new(Errors)
	for _, word := range elms.List() {
		if err := user.Tasks.Add(models.NewTask(word)); err != nil {
			errs.Add(fmt.Errorf("`%s`は登録済みです", word))
		}
	}

	if err := user.Save(); err != nil {
		log.Println("Failed to save user")
		return
	}

	fmt.Println(act.Tweet(newReply(status, elms.ToText("", "を追加しました！"))))
}
