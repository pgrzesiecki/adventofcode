package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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

	re := regexp.MustCompile(`\s+`)
	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = re.ReplaceAllString(line, " ")

		parts := strings.Split(line, " ")

		for i, part := range parts {
			trimmed := strings.TrimSpace(part)
			if trimmed != "" {
				number, err := strconv.Atoi(trimmed)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					continue
				}
				if i == 0 {
					left = append(left, number)
				}
				if i == 1 {
					right = append(right, number)
				}
			}
		}
	}

	// Sort the left and right slices
	sort.Ints(left)
	sort.Ints(right)

	result := 0

	for i := range left {
		result = result + int(math.Abs(float64(left[i])-float64(right[i])))
	}

	fmt.Println(result)
}
