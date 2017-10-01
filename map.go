package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/google/go-github/github"
)

var sourceDir string

func updateMap() {
	var rawDir = fmt.Sprintf("%s/raw", sourceDir)
	if _, err := os.Stat(sourceDir); err != nil && os.IsNotExist(err) {
		if err = os.Mkdir(sourceDir, 0700); err != nil {
			fmt.Printf("Error when updating map file: %s\n", err.Error())
			os.Exit(1)
		}
		if err = os.Mkdir(rawDir, 0700); err != nil {
			fmt.Printf("Error when updating map file: %s\n", err.Error())
			os.Exit(1)
		}
	}
	ctx := context.Background()
	client := github.NewClient(nil)
	_, dirContents, _, err := client.Repositories.GetContents(ctx, "valutac", "gitnore", "/config", nil)
	if err != nil {
		fmt.Printf("Error when updating map file: %s\n", err.Error())
		os.Exit(1)
	}

	dl := grab.NewClient()
	for _, content := range dirContents {
		req, _ := grab.NewRequest(rawDir, content.GetDownloadURL())
		resp := dl.Do(req)
		t := time.NewTicker(500 * time.Millisecond)
		defer t.Stop()
	Loop:
		for {
			select {
			case <-t.C:
				fmt.Printf("\ttransferred %v / %v bytes (%.2f%%)\n",
					resp.BytesComplete(),
					resp.Size,
					100*resp.Progress())

			case <-resp.Done:
				break Loop
			}
		}
		if err := resp.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Download failed: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Download saved to %s \n", resp.Filename)
	}

	files, err := ioutil.ReadDir(rawDir)
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
		data[key] = fmt.Sprintf("%s/%s", rawDir, file.Name())
	}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error when updating map file: %s\n", err.Error())
		os.Exit(1)
	}
	var mapFilePath = fmt.Sprintf("%s/map.json", sourceDir)
	if err := ioutil.WriteFile(mapFilePath, b, 0644); err != nil {
		fmt.Printf("Error when updating map file: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Updating map file succeeded")
	os.Exit(0)
}

func listMap() map[string]string {
	var (
		data        map[string]string
		mapFilePath = fmt.Sprintf("%s/map.json", sourceDir)
	)
	f, err := os.Open(mapFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Map file not available, please run gitnore update")
			os.Exit(1)
		}
		fmt.Printf("Error when opening map file: %s\n", err.Error())
		os.Exit(1)
	}
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		fmt.Printf("Error when opening map file: %s\n", err.Error())
		os.Exit(1)
	}
	return data
}

func printMap(data map[string]string) {
	fmt.Println("Available gitignore configurations:")
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
