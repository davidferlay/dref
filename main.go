package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func main() {
	// Check if a filename argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	// Get file information
	file, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if file.IsDir() {
		log.Fatalf("Error: %s is a directory, not a file.", file.Name())
	}

	// Read the file into a byte slice
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string
	content := string(data)

	// Calculate the character count
	charCount := len(content)

	println("----------------------------------------------------------------------------------")
	fmt.Println("File Information:")
	fmt.Printf("- File Name: %s\n", file.Name())
	fmt.Printf("- Mode: %s\n", file.Mode())
	fmt.Printf("- Size: %s\n", formatSize(file.Size()))
	fmt.Printf("- Character count: %d\n", charCount)
	println("----------------------------------------------------------------------------------")
	//spew.Dump(file)

	//fmt.Println("\nFile Content:")
	//fmt.Println(content)

	inputString := content
	substrings := make(map[string]*substringInfo)

	// Create all possible substrings from the inputString
	for i := 0; i < len(inputString); i++ {
		for j := i + 2; j <= len(inputString); j++ { // Substrings with more than 1 character
			substr := inputString[i:j]
			if info, exists := substrings[substr]; exists {
				info.Occurrences++
			} else {
				substrings[substr] = &substringInfo{
					Substring:      substr,
					CharacterCount: len(substr),
					Occurrences:    1,
					Points:         len(substr),
				}
			}
		}
	}

	// Convert the map of substrings to a slice for sorting
	substringSlice := make([]*substringInfo, 0, len(substrings))
	for _, info := range substrings {
		substringSlice = append(substringSlice, info)
	}

	// Sort the substrings by Points in descending order
	sort.Slice(substringSlice, func(i, j int) bool {
		return substringSlice[i].Points > substringSlice[j].Points
	})

	// Print the top 5 substrings with the most Points
	fmt.Println("Top 5 Substrings:")
	for i, info := range substringSlice {
		if i >= 5 {
			break
		}
		fmt.Printf("%d. Substring: %s\n   Character Count: %d\n   Occurrences: %d\n   Points: %d\n",
			i+1, info.Substring, info.CharacterCount, info.Occurrences, info.Points)
	}

	println("\n-end-")
}

type substringInfo struct {
	Substring      string
	CharacterCount int
	Occurrences    int
	Points         int
}

func formatSize(size int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}

	var i int
	fSize := float64(size)
	for fSize >= 1024 && i < len(sizes)-1 {
		fSize /= 1024
		i++
	}

	return fmt.Sprintf("%.2f %s", fSize, sizes[i])
}
