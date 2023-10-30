package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// Check if a filename argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	// Get file information
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if fileInfo.IsDir() {
		log.Fatalf("Error: %s is a directory, not a file.", fileInfo.Name())
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("File Information:")
	fmt.Printf("- File Name: %s\n", fileInfo.Name())
	fmt.Printf("- Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("- Mode: %s\n", fileInfo.Mode())
	fmt.Printf("- Is Directory: %t\n", fileInfo.IsDir())
	fmt.Printf("- ModTime: %v\n", fileInfo.ModTime())
	println("====")
	spew.Dump(fileInfo)
}
