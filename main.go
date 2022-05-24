package main

import (
	"bufio"
	"fmt"
	"os"
)

const DEBUG = false

// color for printing "error"
const errorColor = "#D38384"

/**
 * callPage
 *
 * @desc: attempt to read and print a summon page.
 *
 * @param: `page` the name of the summon page to call.
 *
 * @return: `true` is page is found. `false` if page is not found.
 */
func callPage(page string) bool {

	pagePath := GetPath(page, DEBUG)

	f, err := os.Open(pagePath)
	if err != nil {
		return false
	}
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		if len(line) > 0 {
			CheckLine(line)
		} else {
			fmt.Println(line)
		}
	}
	f.Close()
	return true
}

func main() {

	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}

	// get syntax highlight colors
	GetColors()
	errStr := ColorStr("error", errorColor)

	switch os.Args[1] {

	case "list":
		e := SmnDirExists()
		if !e {
			fmt.Fprintf(os.Stderr, "%s: summon pages directory not found\n", errStr)
			fmt.Printf("\nrun %s to create it\n", ColorStr("summon init", Cmd))
			os.Exit(1)
		}
		ListNotes(DEBUG)

	case "init":
		e := SmnDirExists()
		if !e {
			err := CreatePagesDir(DEBUG)
			if err != nil {
				fmt.Printf("%s: problem creating pages directory", errStr)
				os.Exit(1)
			} else {
				fmt.Printf("Summon pages directory created successfully at %s%s\n\n", ColorStr(os.Getenv("HOME"), Cmd), ColorStr(NOTES_PATH, Cmd))
				fmt.Println("Now create some of your own pages!")
			}
		}
	
	case "add":
		if len(os.Args) >= 3 {
			err := AddEditPage(os.Args[2], "add", DEBUG)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Fprintf(os.Stderr, 
				"%s: provide a page name to add\n\n\t%s\n", 
				errStr, 
				ColorStr("summon add <page-name>", Cmd))
		}

	case "edit":
		if len(os.Args) >= 3 {
			err := AddEditPage(os.Args[2], "edit", DEBUG)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Fprintf(os.Stderr, 
				"%s: provide a page name to edit\n\n\t%s\n", 
				errStr, 
				ColorStr("summon edit <page-name>", Cmd))
		}


	case "version":
		PrintVersion()

	case "help":
		PrintUsage()

	default:
		// search for a page
		e := SmnDirExists()
		if !e {
			fmt.Fprintf(os.Stderr, "%s: summon page directory not found\n", errStr)
			fmt.Println("run %s to create it", ColorStr("summon init", Cmd))
			os.Exit(1)
		}
		err := callPage(os.Args[1])
		if !err {
			fmt.Fprintf(os.Stderr, "%s: %s is not an avaialble page\n\n", errStr, os.Args[1])
			fmt.Printf("add it as %s%s if you want to create it as a summon page\n", ColorStr("$HOME/.config/summon/", Cmd), ColorStr(os.Args[1], Cmd))
			os.Exit(1)
		}

	}
}
