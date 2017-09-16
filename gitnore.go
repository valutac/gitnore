package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	osuser "os/user"
)

var (
	src, dest *string
	user      *osuser.User
)

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

	var err error
	user, err = osuser.Current()
	if err != nil {
		fmt.Printf("Error when checking current user: %s\n", err.Error())
		os.Exit(1)
	}

	// setup source dir
	source_dir = fmt.Sprintf("%s/.gitnore", user.HomeDir)

	cmd := os.Args[1]

	if cmd == "update" {
		updateMap()
	}

	mapping = listMap()
	if cmd == "list" {
		printMap(mapping)
		os.Exit(0)
	}

	if *src == "" || *dest == "" {
		fmt.Println("src / dest is required")
		usage()
	}

	var (
		b    []byte
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
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("\tupdate \t: Update gitignore dictionary")
	fmt.Println("\tlist \t: List available gitignore")
	fmt.Println()
	fmt.Println("Parameters:")
	fmt.Println("\t-i \t: Select Language (-i python)")
	fmt.Println("\t-o \t: Output filename (default .gitignore)")
	fmt.Println()
	fmt.Println("Example usage:")
	fmt.Println("\t$ gitnore -i python \t\t: default set to .gitignore")
	fmt.Println("\t$ gitnore -i go -u .gitmodule \t: set output file to .gitmodule")
	os.Exit(1)
}
