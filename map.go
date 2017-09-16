package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const (
	MAP_FILE_PATH = "./config/map.json"
	SOURCE_DIR    = "./config"
)

func updateMap() {
	files, err := ioutil.ReadDir(SOURCE_DIR)
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
		data[key] = fmt.Sprintf("%s/%s", SOURCE_DIR, file.Name())
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
		fmt.Println("List avaiables gitignore:")
		var keys []string
		for key := range data {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for i, key := range keys {
			if i+1 == len(keys) {
				fmt.Println(key)
				continue
			}
			fmt.Printf("%s, ", key)
		}
	}
	return data
}
