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

	slices.Sort(sorted)

	for i, v := range elements {
		if v != sorted[i] && v != sorted[elements_len-i-1] {
			return false
		}
	}

	return true
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

	safeLines := 0
	for scanner.Scan() {
		line := scanner.Text()

		elements := convert_to_float64_slice(strings.Split(line, " "))

		if !is_adjacent_levels_safe(elements, adj_level_min, adj_level_max) {
			continue
		}

		if !is_sorted(elements) {
			fmt.Println("Not safe:", elements)
			continue
		}

		safeLines += 1
	}
	fmt.Println(safeLines)
}
