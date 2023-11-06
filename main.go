package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

	findTopSubstrings(inputString)

}

func findTopSubstrings(inputString string) {
	substringCounts := make(map[string]int)
	maxLength := 0

	// Iterate through different substring lengths
	for length := 2; length <= len(inputString)/2; length++ {
		for i := 0; i <= len(inputString)-length; i++ {
			substr := inputString[i : i+length]
			substringCounts[substr]++
			if length > maxLength {
				maxLength = length
			}
		}
	}

	var topSubstrings []SubstringInfo

	// Calculate points for each substring
	for substr, count := range substringCounts {
		if count > 1 {
			occurrences := strings.Count(inputString, substr)
			points := count * occurrences
			topSubstrings = append(topSubstrings, SubstringInfo{
				Substring:      substr,
				CharacterCount: len(substr),
				Occurrences:    occurrences,
				Points:         points,
			})
		}
	}

	// Sort topSubstrings by Points in descending order
	sortByPoints(topSubstrings)

	// Print the top 5 substrings with the most Points
	topCount := 5
	if len(topSubstrings) < topCount {
		topCount = len(topSubstrings)
	}

	fmt.Printf("Top %d substrings with the most Points:\n", topCount)
	for i := 0; i < topCount; i++ {
		fmt.Printf("Substring: %s, Character count: %d, Occurrences: %d, Points: %d\n",
			topSubstrings[i].Substring,
			topSubstrings[i].CharacterCount,
			topSubstrings[i].Occurrences,
			topSubstrings[i].Points,
		)
	}
}

func sortByPoints(substrings []SubstringInfo) {
	for i := range substrings {
		maxIdx := i
		for j := i + 1; j < len(substrings); j++ {
			if substrings[j].Points > substrings[maxIdx].Points {
				maxIdx = j
			}
		}
		substrings[i], substrings[maxIdx] = substrings[maxIdx], substrings[i]
	}
}

type SubstringInfo struct {
	Substring      string
	CharacterCount int
	Occurrences    int
	Points         int
}
