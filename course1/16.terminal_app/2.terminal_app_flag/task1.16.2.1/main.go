package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path string, prefix string, isLast bool, depth int) {

	if isLast || depth <= 0 {
		return
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("failed to read dir: %s", err.Error())
	}

	for i, entry := range entries {
		fmt.Print(prefix)

		newPrefix := prefix

		if i == len(entries)-1 {
			fmt.Print("└── ")
			newPrefix += "    "
		} else {
			fmt.Print("│── ")
			newPrefix += "│   "
		}

		fmt.Println(entry.Name())

		newPath := filepath.Join(path, entry.Name())

		if !entry.IsDir() {
			isLast = true
		}

		printTree(newPath, newPrefix, isLast, depth-1)

	}
}

func main() {
	var depth int

	flag.IntVar(&depth, "n", -1, "tree depth")
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("no path to directory specified")
	}

	path := flag.Arg(0)

	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		path = filepath.Join(wd, path)
	}

	printTree(path, "", false, depth)
}
