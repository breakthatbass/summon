package main

import "os"

// default syntax highlight colors
// adjust with env vars
const rose = "#F295B1"
const gold = "#DBAB79"
const gray = "#B9BFCA"
const cyan = "#71BEF2"
const green = "#5F874D"

// strings that will get highlighted

var Header string
var Desc string
var Cmd string
var Flag string
var Comment string

var errStr string

func GetColors() {
	head := os.Getenv("SUMMON_HEADER_COLOR")
	description := os.Getenv("SUMMON_DESC_COLOR")
	command := os.Getenv("SUMMON_CMD_COLOR")
	flagClr := os.Getenv("SUMMON_FLAG_COLOR")
	CommClr := os.Getenv("SUMMON_COMMENT_COLOR")

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

	if CommClr == "" {
		Comment = green
	} else {
		Comment = CommClr
	}
}
