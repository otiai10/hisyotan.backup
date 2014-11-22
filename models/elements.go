package models

import (
	"regexp"
	"strings"
)

// Elements is parsed status text
type Elements struct {
	raw   string
	dict  map[string]string
	joint string
}

var splitExp = regexp.MustCompile("[ |　]+")

// ParseText parses status text and returns Elements
func ParseText(text string) *Elements {
	dict := map[string]string{}
	for _, w := range splitExp.Split(text, -1) {
		dict[w] = w
	}
	return &Elements{
		raw:   text,
		dict:  dict,
		joint: ",",
	}
}

// Has ...
func (elm *Elements) Has(words ...string) bool {
	for _, w := range words {
		if _, ok := elm.dict[w]; !ok {
			return false
		}
	}
	return true
}

// Remove ...
func (elm *Elements) Remove(words ...string) *Elements {
	for _, w := range words {
		delete(elm.dict, w)
	}
	return elm
}

// Add ...
func (elm *Elements) Add(words ...string) *Elements {
	for _, w := range words {
		elm.dict[w] = w
	}
	return elm
}

// List ...
func (elm *Elements) List() []string {
	res := []string{}
	for _, w := range elm.dict {
		res = append(res, w)
	}
	return res
}

// ToText ...
func (elm *Elements) ToText(args ...string) string {
	args = append(args, "", "")
	return args[0] + " " + strings.Join(elm.List(), elm.joint) + " " + args[1]
}
