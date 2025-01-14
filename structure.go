package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Base folder name for hash map problems
	baseFolder := "4-hash-map"

	// Define subfolders and categorized files
	categories := map[string][]string{
		"01_Basic": {
			"hash_map_implementation.go",
			"count_character_frequency.go",
			"check_anagram_using_map.go",
			"two_sum.go",
			"find_duplicate_elements.go",
			"intersection_of_two_arrays.go",
		},
		"02_Intermediate": {
			"group_anagrams.go",
			"longest_consecutive_sequence.go",
			"subarray_sum_equals_k.go",
			"find_all_duplicates.go",
			"top_k_frequent_elements.go",
			"word_pattern_matching.go",
		},
		"03_Advanced": {
			"minimum_window_substring.go",
			"substring_with_concatenation_of_all_words.go",
			"four_sum_count.go",
			"find_all_anagrams_in_string.go",
			"smallest_subarray_covering_set.go",
			"randomized_set.go",
		},
	}

	// Create the base folder if it doesn't exist
	if err := os.Mkdir(baseFolder, os.ModePerm); err != nil {
		if !os.IsExist(err) {
			fmt.Printf("Error creating base folder: %v\n", err)
			return
		}
	}

	// Loop through categories to create subfolders and files
	for subfolder, files := range categories {
		// Create the subfolder path
		subfolderPath := filepath.Join(baseFolder, subfolder)

		// Create the subfolder
		if err := os.Mkdir(subfolderPath, os.ModePerm); err != nil {
			if !os.IsExist(err) {
				fmt.Printf("Error creating subfolder %s: %v\n", subfolderPath, err)
				continue
			}
		}

		// Create files in the subfolder
		for _, file := range files {
			// Construct the file path
			filePath := filepath.Join(subfolderPath, file)

			// Create the file
			f, err := os.Create(filePath)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", filePath, err)
				continue
			}

			// Write a basic Go template to the file
			template := fmt.Sprintf(`package main

import "fmt"

// %s: Implement your solution here
func solution() {
    // Add your logic here
}

func main() {
    fmt.Println("Running: %s")
    solution()
}
`, file[:len(file)-3], file[:len(file)-3])

			_, err = f.WriteString(template)
			if err != nil {
				fmt.Printf("Error writing to file %s: %v\n", filePath, err)
			}

			// Close the file
			if err := f.Close(); err != nil {
				fmt.Printf("Error closing file %s: %v\n", filePath, err)
			}
		}
	}

	fmt.Println("All subfolders and files created successfully inside the '2-hash-map' folder!")
}
