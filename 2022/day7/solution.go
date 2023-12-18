package day7

import (
	"strconv"
	"strings"
)

//type DirectoryItem interface {
//	Directory | File
//}

type Directory struct {
	Items []any
}

type File struct {
}

type CommandType string

const (
	changeDirectory CommandType = "cd"
	listItems       CommandType = "ls"
)

type Command struct {
	CommandType CommandType
	CommandArgs []string
}

type ListItemsOutput struct {
	IsDirectory bool
	Size        int64
	Name        string
}

func Solution1(input []string) any {
	var (
		outermostDirectory = Directory{}
	)

	for _, line := range input {
		parseLine(line)
	}

	return ""
}

func Solution2(input []string) any {
	return ""
}

func parseLine(line string) (lineData any) {
	var (
		lineArray = strings.Split(line, " ")
	)
	if lineArray[0] == "$" {
		var commandData Command
		switch CommandType(lineArray[1]) {
		case changeDirectory:
			commandData.CommandType = changeDirectory
			commandData.CommandArgs = append(commandData.CommandArgs, lineArray[2])
		case listItems:
			commandData.CommandType = listItems
		}
		lineData = commandData
	} else {
		var listItemsOutput ListItemsOutput
		if lineArray[0] == "dir" {
			listItemsOutput.IsDirectory = true
		} else {
			listItemsOutput.Size, _ = strconv.ParseInt(lineArray[0], 10, 64)
		}
		listItemsOutput.Name = lineArray[1]
	}

	return
}
