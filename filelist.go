package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)
var linelist []string

func main() {
	str, _ := os.Getwd()
	filepath.Walk(str, walkfunc)
}

func walkfunc(path string, info os.FileInfo, err error) error {
	str, _ := os.Getwd()    //当前目录
	path = strings.Replace(path, str, "", -1)
	path = strings.Replace(path, "\\", "/", -1) //适配Windows目录
	fmt.Println(path)
	linelist = append(linelist,path)
	//写入文件
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	for _, line := range linelist {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
