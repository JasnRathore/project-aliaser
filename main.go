package main

import (
	"fmt"
	"os"
	ui "fa/ui"
	api "fa/api"
)

func main() {
	size := len(os.Args)
	if (size < 2) {
		err := ui.RunBubbleTeaApp()
		if err != nil {
			fmt.Printf("Error running Bubble Tea: %v\n", err)
			os.Exit(1)
		}
		return
	}
	args := os.Args[1:]
	switch args[0] {
		case "add","ad":
			api.InsertData(args[1], args[2])
		case "delete","dl":
			api.DeleteData(args[1])
		case "check","ch":
			api.CheckAlias(args[1])
		case "list","ls":
			api.ListAliases()
		case "cd":
			api.GetAlias(args[1])
	}
}
