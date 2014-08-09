package model

import "strings"
import "regexp"

type Words struct {
	values []string
	dict   map[string]string
}

func NewWordSet(text string) (words Words) {
	text = strings.TrimSpace(text)
	exp := regexp.MustCompile("[ 　]+")
	words.values = exp.Split(text, -1)
	words.dict = map[string]string{}
	for _, v := range words.values {
		words.dict[v] = v
	}
	return words
}

func (w Words) refresh() Words {
	w.values = []string{}
	for _, v := range w.dict {
		w.values = append(w.values, v)
	}
	return w
}

func (w Words) Has(keyword string) bool {
	for _, v := range w.values {
		println(keyword, v)
		if keyword == v {
			return true
		}
	}
	return false
}

func (w Words) ToText(args ...string) string {
	var head, foot string
	if len(args) > 0 {
		head = args[0] + " "
	}
	if len(args) > 1 {
		foot = " " + args[1]
	}
	return head + strings.Join(w.values, " ") + foot
}

func (w Words) Filter(filterWords ...string) Words {
	for _, word := range filterWords {
		delete(w.dict, word)
	}
	return w.refresh()
}
