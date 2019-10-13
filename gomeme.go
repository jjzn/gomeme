package main

import (
	"fmt"
	"os"
        "net/http"
        "io/ioutil"
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
        if len(os.Args) < 3 {
                fmt.Println("Error: no URL argument passed")
                os.Exit(1)
        }
        res, err := http.Get(os.Args[2])
        if err != nil {
                defer fmt.Println("Common error causes include missing internet connection, server errrors or wrong URLs.")
                fmt.Println("Error: network error (cannot get resource).")
                os.Exit(1)
        }
        defer res.Body.Close()

        data, err := ioutil.ReadAll(res.Body)
        if err != nil {
                fmt.Println("Error: read error (cannot read HTTP response)")
                os.Exit(1)
        }
        fmt.Println(len(data))
}

func printError() {
	fmt.Println("Unknown command. Try gomeme help.")
}
