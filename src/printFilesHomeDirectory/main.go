package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	filesInHomeDir, err := ioutil.ReadDir(homeDir)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(filesInHomeDir))

	fmt.Println("Count of files and dirs is", len(filesInHomeDir))

	for _, file := range filesInHomeDir {
		go func(f os.FileInfo) {
			fmt.Println(f.Name())

			defer wg.Done()
		}(file)
	}
	wg.Wait()
}
