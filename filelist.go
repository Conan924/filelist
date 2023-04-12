package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	linelist  []string
	blacklist []string
)

func main() {
	readConfig()

	str, _ := os.Getwd()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		writeToFile("filelist_output.txt")
	}()

	filepath.Walk(str, func(path string, info os.FileInfo, err error) error {
		if isFile(path) {
			if !isBlackListed(path) {
				path = strings.Replace(path, str, "", -1)
				path = strings.Replace(path, "\\", "/", -1)
				fmt.Println(path)
				linelist = append(linelist, path)
				writeToFile(path)
			}
		}
		return nil
	})

	wg.Wait()
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func isBlackListed(path string) bool {
	for _, suffix := range blacklist {
		if strings.HasSuffix(strings.ToLower(path), strings.ToLower(suffix)) {
			return true
		}
	}
	return false
}

func readConfig() {
	file, err := os.Open("filelist_config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		blacklist = append(blacklist, strings.ToLower(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func writeToFile(path string) {
	f, err := os.OpenFile("filelist_output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(path + "\n"); err != nil {
		log.Fatal(err)
	}
}
