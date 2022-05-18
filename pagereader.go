package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/muesli/termenv"
)

// check the command string to change the color of certain words
func checkCmd(line string) []termenv.Style {
	p := termenv.ColorProfile()
	splitWords := strings.Split(line, " ")
	n := len(splitWords)
	termarr := make([]termenv.Style, n)
	inComment := false

	for i := 0; i < n; i++ {
		if splitWords[i] == "#" || strings.HasPrefix(splitWords[i], "#") || inComment {
			// comment - gray out the rest of the line
			termarr[i] = termenv.String(splitWords[i]).Foreground(p.Color(Comment))
			inComment = true
		} else if strings.HasPrefix(splitWords[i], "-") || strings.HasPrefix(splitWords[i], "<") {
			termarr[i] = termenv.String(splitWords[i]).Foreground(p.Color(Flag))
		} else {
			termarr[i] = termenv.String(splitWords[i]).Foreground(p.Color(Cmd))
		}
	}
	return termarr
}

// firgure out if a line is a header, description, or a command and format it properly
func CheckLine(line string) {
	p := termenv.ColorProfile()
	r := []rune(line)
	var s termenv.Style

	if unicode.IsLetter(r[0]) {
		// first line, note title
		s = termenv.String(line).Foreground(p.Color(Header)).Bold()
		fmt.Println(s)
	} else if r[0] == '-' {
		// description line
		s = termenv.String(line).Foreground(p.Color(Desc))
		fmt.Println(s)
	} else if r[0] == '\t' || r[0] == ' ' {
		// the actual command
		// whether it's spaces or tabs, format this consistently with 4 spaces
		if r[0] == '\t' {
			r[0] = ' '
			fmt.Printf("   ")
		}
		// look for flags in the command
		cmd := checkCmd(string(r))
		for i := 0; i < len(cmd); i++ {
			fmt.Printf("%s ", cmd[i])
		}
		fmt.Printf("\n")
	}
}
