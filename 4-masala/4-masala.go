package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func generateRandomFileName(filePath string) string {
	return filePath + strconv.Itoa(rand.Intn(100)) + ".txt"
}

func suggestAvailableFileNames(basePath string, numSuggestions int) {
	for i := 0; i < numSuggestions; i++ {
		filePath := generateRandomFileName(basePath)
		if !checkFileExists(filePath) {
			fmt.Printf("%v ", filePath)
		}
	}
	fmt.Println()
}

func main() {
	var filePath string
	fmt.Print("Enter file path: ")
	fmt.Scan(&filePath)

	if checkFileExists(filePath + ".txt") {
		fmt.Println("File already exists. You can create files with names like these:")
		suggestAvailableFileNames(filePath, 3)
		return
	}

	_, err := os.Create(filePath + ".txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File created successfully:", filePath+".txt")
}
