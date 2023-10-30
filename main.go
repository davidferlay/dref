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
	topN := 5

	/////////////////////////////////////////////////////////////////////////

	characterCounts := make(map[string]int)

	// Process the string and count character combinations
	for i := 0; i < len(inputString); i++ {
		for j := i + 1; j <= len(inputString); j++ {
			substring := inputString[i:j]
			characterCounts[substring]++
		}
	}

	// Create a slice of character combinations and their counts
	var characterCountSlice []CharCount

	for combination, count := range characterCounts {
		points := count * len(combination)
		characterCountSlice = append(characterCountSlice, CharCount{Combination: combination, Count: count, Points: points})
	}

	// Sort the slice by Points in descending order
	sort.Slice(characterCountSlice, func(i, j int) bool {
		return characterCountSlice[i].Points > characterCountSlice[j].Points
	})

	// Print the top N combinations of characters with Points
	fmt.Printf("Top %d combinations of characters that occur the most:\n", topN)
	for i := 0; i < topN && i < len(characterCountSlice); i++ {
		fmt.Printf("%s: Character count %d, Count %d, Points %d\n",
			characterCountSlice[i].Combination, len(characterCountSlice[i].Combination),
			characterCountSlice[i].Count, characterCountSlice[i].Points)
	}

	println("-end-")
}

type CharCount struct {
	Combination string
	Count       int
	Points      int
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
