package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func convert_to_float64_slice(strings []string) []float64 {
	numbers := []float64{}
	for _, v := range strings {
		current, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}

		numbers = append(numbers, float64(current))
	}
	return numbers
}

func is_sorted(elements []float64) bool {
	elements_len := len(elements)
	sorted := make([]float64, elements_len)
	copy(sorted, elements)

	is_sorted_asc := true
	is_sorted_desc := true
	slices.Sort(sorted)

	for i, v := range elements {
		if v != sorted[i] {
			is_sorted_asc = false
		}

		if v != sorted[elements_len-i-1] {
			is_sorted_desc = false
		}
	}

	return is_sorted_asc || is_sorted_desc
}

func is_adjacent_levels_safe(elements []float64, min float64, max float64) bool {
	for i, v := range elements {
		if i == 0 {
			continue
		}

		abs_val := math.Abs(v - elements[i-1])
		if abs_val > max || abs_val < min {
			return false
		}
	}

	return true
}

func get_slice_without_element(slice []float64, index int) []float64 {
	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if the index is out of range
	}

	slice_without_element := make([]float64, len(slice)-1)
	copy(slice_without_element, slice[:index])
	copy(slice_without_element[index:], slice[index+1:])

	return slice_without_element
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name as an argument")
		return
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	adj_level_min := float64(1)
	adj_level_max := float64(3)

	safe_lines := 0
	dampener_safe_lines := 0

	for scanner.Scan() {
		line := scanner.Text()

		elements := convert_to_float64_slice(strings.Split(line, " "))

		if is_adjacent_levels_safe(elements, adj_level_min, adj_level_max) && is_sorted(elements) {
			safe_lines += 1
			continue
		}

		for i := range elements {
			dampener_elements := get_slice_without_element(elements, i)

			if is_adjacent_levels_safe(dampener_elements, adj_level_min, adj_level_max) && is_sorted(dampener_elements) {
				dampener_safe_lines += 1
				fmt.Println("Dampener Safe Line: ", dampener_elements)
				break
			}
		}
	}

	fmt.Println("RAW lines: ", safe_lines)
	fmt.Println("Dampener Safe Lines: ", dampener_safe_lines)
	fmt.Println("Total lines: ", dampener_safe_lines+safe_lines)
}

// 328
