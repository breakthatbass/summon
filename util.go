package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/muesli/termenv"
)

const VERSION = "0.0.1"

// hold all the pages to the notes here
const NOTES_PATH = "/.config/summon/"

func GetErrColor() termenv.Style {
	p := termenv.ColorProfile()
	return termenv.String("error").Foreground(p.Color("#D38384"))
}

func ColorStr(s string, hex string) termenv.Style {
	p := termenv.ColorProfile()
	return termenv.String(s).Foreground(p.Color(hex))
}

/**
 * getPath
 *
 * get the file path to the desired note
 *
 * noteFile: the note to check
 *
 * returns: path to noteFile
 **/
func GetPath(page string, DEBUG bool) string {
	if DEBUG {
		return fmt.Sprintf("pages/%s", page)
	} else {
		path := os.Getenv("HOME")

		// this should like something like '~/.config/summon/note-to-view
		return fmt.Sprintf("%s%s%s", path, NOTES_PATH, page)
	}
}

// check for directory
func SmnDirExists() bool {
	var smnDir string
	if DEBUG {
		smnDir = "pages"
	} else {
		smnDir = fmt.Sprintf("%s%s", os.Getenv("HOME"), NOTES_PATH)
	}
	_, err := os.Stat(smnDir)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// create directory
func CreatePagesDir() error {
	var smnDir string
	if DEBUG {
		smnDir = "pages"
	} else {
		smnDir = fmt.Sprintf("%s%s", os.Getenv("HOME"), NOTES_PATH)
	}
	return os.Mkdir(smnDir, 0755)
}

func ListNotes(DEBUG bool) bool {
	var path string
	p := termenv.ColorProfile()

	if DEBUG {
		path = "pages"
	} else {
		//home := os.Getenv("HOME")
		path = fmt.Sprintf("%s%s", os.Getenv("HOME"), NOTES_PATH)
	}
	pages, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}

	for _, f := range pages {
		s := termenv.String(f.Name()).Foreground(p.Color(Cmd))
		fmt.Printf("%s\n", s)
	}
	return true
}

func PrintVersion() {
	p := termenv.ColorProfile()
	summon := termenv.String("summon").Foreground(p.Color(Flag)).Bold()
	fmt.Printf("%s %s\n", summon, VERSION)
}

func PrintUsage() {
	p := termenv.ColorProfile()
	summon := termenv.String("summon").Foreground(p.Color(Flag)).Bold()
	usage := termenv.String("USAGE:").Foreground(p.Color(Cmd))
	commands := termenv.String("COMMANDS:").Foreground(p.Color(Cmd))

	fmt.Printf("%s %s\n\n", summon, VERSION)
	fmt.Printf("Personalized man pages for things you want to remember\n\n")
	fmt.Println(usage)
	fmt.Printf("\tsummon [page/command]\n\n")

	fmt.Println(commands)
	fmt.Printf("\tlist - list all available pages\n")
	fmt.Printf("\thelp - print this list and quit\n")
	fmt.Printf("\tversion - print current version\n")
	fmt.Printf("\tinit - create pages directory at `$HOME/.config/summon/`\n")
}
