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

/**
 * ColorStr
 *
 * @desc: colorize `s` with `hex` color.
 *
 * @param: `s` string to colorize.
 * @param: `hex` hex color
 *
 * @return: `s` colored with `hex`.
 **/
func ColorStr(s string, hex string) termenv.Style {
	p := termenv.ColorProfile()
	return termenv.String(s).Foreground(p.Color(hex))
}

/**
 * getPath
 *
 * @desc: build the file path to the desired page
 *
 * @param: `page` - the note to check
 *
 * @return: path to paege
 **/
func GetPath(page string, DEBUG bool) string {
	path := os.Getenv("HOME")
	if DEBUG {
		return fmt.Sprintf("pages/%s", page)
	}
	// this should like something like '~/.config/summon/pagename
	return fmt.Sprintf("%s%s%s", path, NOTES_PATH, page)
}

/**
 * SmnDirExists
 *
 * @desc: check if `$HOME/.config/summon` exists
 *
 * @return - `true` if it exists, else `false`.
 */
func SmnDirExists() bool {
	//var smnDir string
	smnDir := fmt.Sprintf("%s%s", os.Getenv("HOME"), NOTES_PATH)
	if DEBUG {
		smnDir = "pages"
	}
	_, err := os.Stat(smnDir)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

/**
 * CreatePagesDir
 *
 * @desc: create `$HOME/.config/summon`
 *
 * @return: error else nil
 */
func CreatePagesDir() error {
	smnDir := fmt.Sprintf("%s%s", os.Getenv("HOME"), NOTES_PATH)
	if DEBUG {
		smnDir = "pages"
	}
	return os.Mkdir(smnDir, 0755)
}

/**
 * ListNotes
 *
 * @desc: list all summon pages in summon directory
 *
 * @param: `DEBUG` - global debug constant.
 *
 * @return: `true` if successful, else `false` if error reading directory.
 */
func ListNotes(DEBUG bool) bool {
	p := termenv.ColorProfile()
	path := fmt.Sprintf("%s%s", os.Getenv("HOME"), NOTES_PATH)

	if DEBUG {
		path = "pages"
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
