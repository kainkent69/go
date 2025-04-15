package main

import (
	"fmt"
	"sort"
)

func interleaveDistinct(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	counts := make(map[int]int)
	for _, val := range input {
		counts[val]++
	}

	uniqueVals := make([]int, 0, len(counts))
	for val := range counts {
		uniqueVals = append(uniqueVals, val)
	}

	// Sort unique values by frequency in descending order
	sort.Slice(uniqueVals, func(i, j int) bool {
		return counts[uniqueVals[i]] > counts[uniqueVals[j]]
	})

	mostFrequent := uniqueVals[0]
	mostFrequentCount := counts[mostFrequent]

	result := make([]int, 0, len(input))
	remainingCounts := make(map[int]int)
	for k, v := range counts {
		remainingCounts[k] = v
	}

	for i := 0; i < mostFrequentCount; i++ {
		// Append the most frequent element first
		if remainingCounts[mostFrequent] > 0 {
			result = append(result, mostFrequent)
			remainingCounts[mostFrequent]--
		}

		// Append other unique elements
		for _, val := range uniqueVals {
			if val != mostFrequent && remainingCounts[val] > 0 {
				result = append(result, val)
				remainingCounts[val]--
			}
		}
	}

	// Append any remaining elements (shouldn't happen ideally if interleaving is perfect)
	for _, val := range uniqueVals {
		for j := 0; j < remainingCounts[val]; j++ {
			result = append(result, val)
		}
	}

	// A final pass to try and ensure no adjacent duplicates (might require more sophisticated logic
	// if the frequency difference is too large)
	finalResult := make([]int, 0, len(result))
	if len(result) > 0 {
		finalResult = append(finalResult, result[0])
		for i := 1; i < len(result); i++ {
			if result[i] != finalResult[len(finalResult)-1] {
				finalResult = append(finalResult, result[i])
			} else {
				// If a duplicate is found, try to insert a different element if available
				for j := 0; j < len(uniqueVals); j++ {
					otherVal := uniqueVals[j]
					if otherVal != result[i] && remainingCounts[otherVal] > 0 {
						finalResult = append(finalResult, otherVal)
						remainingCounts[otherVal]--
						finalResult = append(finalResult, result[i])
						break
					}
					if j == len(uniqueVals)-1 {
						// If no other element is available, we have a problem.
						// In a perfectly balanced scenario, this shouldn't happen.
						finalResult = append(finalResult, result[i]) // Force append as a last resort
					}
				}
			}
		}
	}

	return finalResult
}

func main() {
	input1 := []int{1, 1, 1, 2, 2, 2}
	output1 := interleaveDistinct(input1)
	fmt.Println("Input:", input1, "Output:", output1) // Expected: [1 2 1 2 1 2] or similar

	input2 := []int{1, 1, 1, 1, 2, 2}
	output2 := interleaveDistinct(input2)
	fmt.Println("Input:", input2, "Output:", output2) // Expected: [1 2 1 2 1 1] or similar

	input3 := []int{5, 5, 3, 3, 8}
	output3 := interleaveDistinct(input3)
	fmt.Println("Input:", input3, "Output:", output3) // Expected: [5 3 5 3 8] or similar

	input4 := []int{7, 7, 7, 8, 9}
	output4 := interleaveDistinct(input4)
	fmt.Println("Input:", input4, "Output:", output4) // Expected: [7 8 7 9 7] or similar
}
