package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var src, dest *string

var mapping map[string]string

func init() {
	src = flag.String("i", "", "Source File")
	dest = flag.String("o", ".gitignore", "Destination File")
}

func main() {
	flag.Parse()
	flag.Usage = usage

	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]

	if cmd == "update" {
		updateMap()
	}

	if cmd == "list" {
		listMap(true)
		os.Exit(0)
	}

	mapping = listMap(false)

	if *src == "" || *dest == "" {
		fmt.Println("src / dest is required")
		usage()
	}

	var (
		b    []byte
		err  error
		ok   bool
		path string
	)

	if path, ok = mapping[*src]; !ok {
		fmt.Printf("unknown source file of: %s\n", *src)
		os.Exit(1)
	}

	if b, err = ioutil.ReadFile(path); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err = ioutil.WriteFile(*dest, b, 0644); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Writing %s into %s.\n", path, *dest)
}

func usage() {
	fmt.Println("Usage: ./gitnore -i language -o .gitignore")
	flag.PrintDefaults()
	os.Exit(1)
}
