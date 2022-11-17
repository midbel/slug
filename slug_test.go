package slug

import (
	"testing"
)

func TestSlug(t *testing.T) {
	data := []struct {
		Input string
		Want  string
	}{
		{
			Input: "foobar",
			Want:  "foobar",
		},
		{
			Input: "foo-bar",
			Want:  "foo-bar",
		},
		{
			Input: "foo bar",
			Want:  "foo_bar",
		},
		{
			Input: " foo_bar ",
			Want:  "foo_bar",
		},
		{
			Input: "-foo-bar-",
			Want:  "foo-bar",
		},
		{
			Input: "foo, bar?",
			Want:  "foo_bar",
		},
	}
	for _, d := range data {
		got := Slug(d.Input)
		if got != d.Want {
			t.Errorf("unexpected result: want %s, got %s", d.Want, got)
		}
	}
}
