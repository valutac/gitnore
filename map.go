package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const MAP_FILE_PATH = "./config/map.json"

func updateMap() {
	files, err := ioutil.ReadDir("./config")
	if err != nil {
		fmt.Printf("Error when updating map file: %s\n", err.Error())
		os.Exit(1)
	}
	var data = make(map[string]string)
	for _, file := range files {
		split := strings.Split(file.Name(), ".")
		if len(split) != 2 || split[1] != "gitignore" {
			continue
		}
		key := strings.ToLower(split[0])
		data[key] = fmt.Sprintf("config/%s", file.Name())
	}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error when updating map file: %s\n", err.Error())
		os.Exit(1)
	}
	if err := ioutil.WriteFile(MAP_FILE_PATH, b, 0644); err != nil {
		fmt.Printf("Error when updating map file: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Updating map file succed")
	os.Exit(0)
}

func listMap(print bool) map[string]string {
	var data map[string]string
	f, err := os.Open(MAP_FILE_PATH)
	if err != nil {
		fmt.Printf("Error when opening map file: %s\n", err.Error())
		os.Exit(1)
	}
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		fmt.Printf("Error when opening map file: %s\n", err.Error())
		os.Exit(1)
	}
	if print {
		for key := range data {
			fmt.Printf("%s ", key)
		}
	}
	return data
}
