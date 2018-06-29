package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: changelog2json <changelog_filepath>")
		os.Exit(1)
	}

	text, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadFile() failed: %v\n", err)
	}

	clogs, err := parse(string(text))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %v\n", err)
	}

	jsonBytes, err := json.Marshal(clogs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json marshal failed: %v\n", err)
	}
	fmt.Println(string(jsonBytes))
}
