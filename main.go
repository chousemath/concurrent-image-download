package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	log "github.com/Sirupsen/logrus"
)

var wg sync.WaitGroup
var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

const (
	dirMain = "images"
)

func main() {
	fmt.Println("v0.0.1")
	dirLists := filepath.Join(".", "image-lists")
	files, err := ioutil.ReadDir(dirLists)
	if err != nil {
		log.Fatalf("Error reading files in ./image-lists, %v\n", err)
	}

	dirMain := filepath.Join(".", "images")
	if err := os.MkdirAll(dirMain, os.ModePerm); err != nil {
		log.Fatalf("Error creating the directory %s\n", dirMain)
	}

	var fileName string
	var label string
	var dirSub string
	var dirTXT string

	for _, f := range files {
		fileName = f.Name()
		if fileName == ".gitignore" {
			continue
		}
		label = fileName[0 : len(fileName)-len(filepath.Ext(fileName))]
		dirSub = filepath.Join(dirMain, label)
		if err := os.MkdirAll(dirSub, os.ModePerm); err != nil {
			log.Fatalf("Error creating the directory %s\n", dirSub)
		}

		dirTXT = filepath.Join(dirLists, fileName)
		lines, err := os.Open(dirTXT)
		if err != nil {
			log.Fatalf("Error opening file %s\n", dirTXT)
		}
		defer lines.Close()

		scanner := bufio.NewScanner(lines)
		for scanner.Scan() {
			worker(scanner.Text(), dirSub)
		}

		if err = scanner.Err(); err != nil {
			log.Fatalf("Error with scanner %v\n", err)
		}
	}
}

func worker(url string, dirName string) {
	response, err := http.Get(url)
	if err != nil {
		log.Warningf("There was an error downloading a file, %v\n", err)
		return
	}
	defer response.Body.Close()
	// open a file for writing
	buf := make([]rune, 24)
	for i := range buf {
		buf[i] = chars[rand.Intn(len(chars))]
	}
	file, err := os.Create(filepath.Join(dirName, fmt.Sprintf("%s.jpg", string(buf))))
	if err != nil {
		log.Warningf("There was an error opening a file, %v\n", err)
		return
	}
	defer file.Close()
	// dump response body into a file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Warningf("There was an error copying the file, %v\n", err)
	}
}
