package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define the directory structure in sequential order
	categories := []struct {
		name  string
		files []string
	}{
		{
			"01_Basic",
			[]string{
				"01_ReverseArray.go",
				"02_FindMaxMin.go",
				"03_SumAndAverage.go",
				"04_CheckSorted.go",
				"05_MergeSortedArrays.go",
				"06_CountEvenOdd.go",
				"07_FindIndex.go",
				"08_FrequencyOfElements.go",
				"09_RemoveElement.go",
				"10_CopyArray.go",
			},
		},
		{
			"02_Intermediate",
			[]string{
				"01_RotateArray.go",
				"02_SecondLargest.go",
				"03_PairsWithGivenSum.go",
				"04_RemoveDuplicates.go",
				"05_LongestSubarrayEqual01.go",
				"06_FindMissingNumber.go",
				"07_Sort012.go",
				"08_IntersectionOfArrays.go",
				"09_UnionOfArrays.go",
				"10_MajorityElement.go",
			},
		},
		{
			"03_Advanced",
			[]string{
				"01_MaxProductSubarray.go",
				"02_SubarrayWithGivenSum.go",
				"03_DistinctElementsInWindow.go",
				"04_MissingAndRepeating.go",
				"05_TrappingRainwater.go",
				"06_LongestConsecutiveSequence.go",
				"07_SmallestSubarrayWithSum.go",
				"08_MaxSumSubarrayOfSizeK.go",
				"09_MergeOverlappingIntervals.go",
				"10_MedianOfSortedArrays.go",
			},
		},
	}

	// Create directories and files in sequence
	for _, category := range categories {
		// Create category directory
		if err := os.MkdirAll(category.name, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", category.name, err)
			continue
		}

		// Create files within the directory
		for _, file := range category.files {
			filePath := filepath.Join(category.name, file)
			// Create file
			if _, err := os.Create(filePath); err != nil {
				fmt.Printf("Error creating file %s: %v\n", filePath, err)
				continue
			}
			fmt.Printf("Created file: %s\n", filePath)
		}
	}

	fmt.Println("Sequential directory and file structure created successfully!")
}
