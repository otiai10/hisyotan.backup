package models

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestParseText(t *testing.T) {
	elms := ParseText("some strings")
	Expect(t, elms).TypeOf("*models.Elements")
}

func TestElements_Has(t *testing.T) {
	elms := ParseText("@xxxx aaa       bbb　　　ccccc")
	Expect(t, elms.Has("@xxxx")).ToBe(true)
	Expect(t, elms.Has("ccccc")).ToBe(true)
	Expect(t, elms.Has("zzzzz")).ToBe(false)
}

func TestElements_ToText(t *testing.T) {
	elms := ParseText("@xxxx aaa       bbb　　　ccccc")
	Expect(t, elms.ToText()).ToBe(" @xxxx,aaa,bbb,ccccc ")
	Expect(t, elms.ToText("pref", "suf")).ToBe("pref @xxxx,aaa,bbb,ccccc suf")
}

func TestElements_Add(t *testing.T) {
	elms := ParseText("aaa bbb ccc")
	Expect(t, elms.ToText()).ToBe(" aaa,bbb,ccc ")
	elms.Add("ddd")
	Expect(t, elms.ToText()).ToBe(" aaa,bbb,ccc,ddd ")
}

func TestElements_Remove(t *testing.T) {
	elms := ParseText("aaa bbb ccc")
	Expect(t, elms.ToText()).ToBe(" aaa,bbb,ccc ")
	elms.Remove("bbb")
	Expect(t, elms.ToText()).ToBe(" aaa,ccc ")
}
