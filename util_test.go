package main

import (
	"os"
	"testing"
)

func TestGetPath(t *testing.T) {
	home := os.Getenv("HOME")
	var tests = []struct {
		input string
		debug bool
		want  string
	}{
		{"", true, "pages/"},
		{"ssh", true, "pages/ssh"},
		{"", false, home + "/.config/summon/"},
		{"ssh", false, home + "/.config/summon/ssh"},
	}
	for _, test := range tests {
		if got := GetPath(test.input, test.debug); got != test.want {
			t.Errorf("GetPath(%q, %t) = %v", test.input, test.debug, got)
		}
	}
}
