package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func erR(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func get(url string) {
	istek, err := http.Get(url)
	erR(err)
	html, err := ioutil.ReadAll(istek.Body)
	erR(err)
	fmt.Print(string(html))
}

func argsRun(args []string) {
	if args[1] == "-h" || args[1] == "--help" || args[1] == "--manual" {
		homeHelp()
	} else {
		if strings.Contains(args[1], "http://") || strings.Contains(args[1], "https://") {
			get(args[1])
		} else {
			get("https://" + args[1])
		}
	}
}

func home() {
	fmt.Println("gurl: try 'gurl --help' or 'gurl --manual' for more information")
	os.Exit(0)
}

func homeHelp() {
	fmt.Println("Usage: gurl <url>")
	os.Exit(0)
}

func main() {
	args := os.Args
	switch {
	case len(args) == 2:
		argsRun(args)
	default:
		home()
	}

}
