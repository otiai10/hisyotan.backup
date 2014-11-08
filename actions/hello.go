package actions

import (
	"fmt"

	"github.com/otiai10/botw"
)

// Hello ...
type Hello struct {
	*botw.Action
}

// Match ...
func (act *Hello) Match(status botw.Status) bool {
	return false
	/*
		    if filtered(status) {
		        return false
		    }
			return models.ParseText(status.Text).Has(conf.My.Name(true), "hello")
	*/
}

// Execute ...
func (act *Hello) Execute(status botw.Status) {
	fmt.Println(status.Text)
}
