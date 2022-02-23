package main

import (
	"bufio"
	"fmt"
	"os"
)

const DEBUG = false

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
		// printUsage()
		fmt.Println("usage: summon [command] <option>")
		os.Exit(1)
	}

	// get syntax highlight colors
	GetColors()

	switch os.Args[1] {

	case "list":
		ListNotes(DEBUG)

	case "help":
		PrintUsage()

	case "init":
		e := SmnDirExists()
		if !e {
			err := CreatePagesDir()
			if err != nil {
				fmt.Printf("error: problem creating pages directory")
				os.Exit(1)
			} else {
				fmt.Printf("Summon pages directory created successfully at %s%s\b", os.Getenv("HOME"), NOTES_PATH)
				fmt.Println("Now create some of your own pages!")
			}
		}

	case "version":
		PrintVersion()

	default:
		err := callPage(os.Args[1])
		if !err {
			fmt.Fprintf(os.Stderr, "error: %s is not an avaialble page\n\n\trun `summon add %s` to create it\n", os.Args[1], os.Args[1])
			os.Exit(1)
		}
	}
}
