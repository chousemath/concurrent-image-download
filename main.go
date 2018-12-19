package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"
)

var wg sync.WaitGroup
var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

const (
	dirMain = "images"
)

func main() {
	fmt.Println("v0.0.1")
	rand.Seed(time.Now().UnixNano())
	dirName := os.Args[1]

	dirMainPath := filepath.Join(".", "images")
	if err := os.MkdirAll(dirMainPath, os.ModePerm); err != nil {
		fmt.Println(err)
	}

	dirSubPath := filepath.Join(".", "images", dirName)
	if err := os.MkdirAll(dirSubPath, os.ModePerm); err != nil {
		fmt.Println(err)
	}

	newPath := path.Join(".", "images", dirName)
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		fmt.Println("There was an error creating the directory...")
		return
	}

	file, err := os.Open("./images.txt")
	if err != nil {
		fmt.Println("Could not open images.txt file...")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		worker(scanner.Text(), dirName)
	}
}

// worker represents each go routine
func worker(url string, dirName string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("There was an error downloading a file...")
		return
	}
	defer response.Body.Close()
	// open a file for writing
	buf := make([]rune, 24)
	for i := range buf {
		buf[i] = chars[rand.Intn(len(chars))]
	}
	file, err := os.Create(fmt.Sprintf("./images/%s/%s.jpg", dirName, string(buf)))
	if err != nil {
		fmt.Println("There was an error opening a file...")
		return
	}
	defer file.Close()
	// dump response body into a file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("There was an error copying the file...")
	}
}
