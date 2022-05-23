package main

import (
	"fmt"
	"strings"
	"unicode"
	"os"
	"os/exec"

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


// open a text editor to edit a page
// TODO add template page to open up to for new page
func AddEditPage(page string, todo string, DEBUG bool) error {

	// get path to page
	p := GetPath(page, DEBUG)
	var err error
	err = nil

	// does that page exist?
	if _, err := os.Stat(p); err == nil {
		// page exists
		if todo == "add" {
			// if page exists already but we're trying to add it...
			return fmt.Errorf("page '%s' doesn't exist.\nTry adding it instead.\n\n\t%s %s\n",
								page, 
								ColorStr("summon edit", Cmd),
								ColorStr(page, Cmd))
		}
	} else {
		if todo == "edit" {
			// if page doesn't exist but we're trying to edit it...
			return fmt.Errorf("page '%s' doesn't exist.\nTry adding it instead.\n\n\t%s %s\n",
								page, 
								ColorStr("summon add", Cmd),
								ColorStr(page, Cmd))
		}
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		err = fmt.Errorf("EDITOR global variable not found. Defaulting to nano")
		editor = "usr/bin/nano"
	}

	// TODO when template page is implemented,
	// copy that page to this new page
	// then open it

	// run exec on opening that file with editor
	cmd := exec.Command(editor, p)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = cmd.Stderr
	cmd.Run()
	return err
}