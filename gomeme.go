package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
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
        case "view":
                viewMeme()
	default:
		printError()
	}
}

var memefolder = os.Getenv("HOME") + "/memes/"

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
		defer fmt.Println("Common error causes include missing internet connection, server errors or wrong URLs.")
		fmt.Println("Error: network error (cannot get resource).")
		os.Exit(1)
	}
	defer res.Body.Close()

	mime := res.Header.Get("Content-Type")
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: read error (cannot read HTTP response)")
		os.Exit(1)
	}
	if mime != "image/jpeg" && mime != "image/png" {
		fmt.Println("Error: bad response content-type (must be image/jpeg or image/png)")
		os.Exit(1)
	}

	fmt.Printf("Sucess! (get %d bytes)\n", len(data))
	filename := memefolder + time.Now().UTC().Format("020106_030405") + "." + strings.Split(mime, "/")[1]
	fmt.Printf("Saving file to %v\n", filename)

	err = ioutil.WriteFile(filename, data, 0664)
	if err != nil {
		fmt.Println("Error: write error (cannot write to file)")
		os.Exit(1)
	}
}

func viewMeme() {
        if len(os.Args) < 3 {
                fmt.Println("Error: no local or remote meme ar  gument passed")
                os.Exit(1)
        }

        os.exec.Command("xdg-open", os.Args[2]).Run()
}

func printError() {
	fmt.Println("Unknown command. Try gomeme help.")
}
