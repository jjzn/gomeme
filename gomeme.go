package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printError()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "help":
		printHelp()
	case "get":
		getMeme()
	default:
		printError()
	}
}

const helpText = `
Usage: gomeme command [arguments]

Commands:
    help:
        Show this help text
    get url:
        Download the meme at url and save it to $HOME/memes
`
func printHelp() {
	fmt.Println(helpText)
}

func getMeme() {
}

func printError() {
	fmt.Println("Unknown command. Try gomeme help.")
}
