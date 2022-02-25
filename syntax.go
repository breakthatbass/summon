package main

import "os"

// default syntax highlight colors
// adjust with env vars
const rose = "#F295B1"
const gold = "#DBAB79"
const gray = "#B9BFCA"
const cyan = "#71BEF2"

// strings that will get highlighted

var Header string
var Desc string
var Cmd string
var Flag string

var errStr string

func GetColors() {
	head := os.Getenv("SUMMON_HEADER_COLOR")
	description := os.Getenv("SUMMON_DESC_COLOR")
	command := os.Getenv("SUMMON_CMD_COLOR")
	flagClr := os.Getenv("SUMMON_FLAG_COLOR")

	if head == "" {
		Header = rose
	} else {
		Header = head
	}

	if description == "" {
		Desc = gold
	} else {
		Desc = description
	}

	if command == "" {
		Cmd = gray
	} else {
		Cmd = command
	}

	if flagClr == "" {
		Flag = cyan
	} else {
		Flag = flagClr
	}
}
