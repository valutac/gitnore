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

var source_dir string

func updateMap() {
	var raw_dir = fmt.Sprintf("%s/raw", source_dir)
	if _, err := os.Stat(source_dir); err != nil && os.IsNotExist(err) {
		if err = os.Mkdir(source_dir, 0700); err != nil {
			fmt.Printf("Error when updating map file: %s\n", err.Error())
			os.Exit(1)
		}
		if err = os.Mkdir(raw_dir, 0700); err != nil {
			fmt.Printf("Error when updating map file: %s\n", err.Error())
			os.Exit(1)
		}
	}
	ctx := context.Background()
	client := github.NewClient(nil)
	_, dir_contents, _, err := client.Repositories.GetContents(ctx, "valutac", "gitnore", "/config", nil)
	if err != nil {
		fmt.Printf("Error when updating map file: %s\n", err.Error())
		os.Exit(1)
	}

	dl := grab.NewClient()
	for _, content := range dir_contents {
		req, _ := grab.NewRequest(raw_dir, content.GetDownloadURL())
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

	files, err := ioutil.ReadDir(raw_dir)
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
		data[key] = fmt.Sprintf("%s/%s", raw_dir, file.Name())
	}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error when updating map file: %s\n", err.Error())
		os.Exit(1)
	}
	var map_file_path = fmt.Sprintf("%s/map.json", source_dir)
	if err := ioutil.WriteFile(map_file_path, b, 0644); err != nil {
		fmt.Printf("Error when updating map file: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Updating map file succed")
	os.Exit(0)
}

func listMap() map[string]string {
	var (
		data          map[string]string
		map_file_path = fmt.Sprintf("%s/map.json", source_dir)
	)
	f, err := os.Open(map_file_path)
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
