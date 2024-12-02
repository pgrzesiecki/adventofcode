package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	maxDiff := float64(3)

	safeLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		asc := true
		ascSet := false
		safe := true
		fault := 0

		parts := strings.Split(line, " ")

		elements := []int{}
		for _, v := range parts {
			current, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}

			elements = append(elements, current)
		}

		for i, v := range elements {
			if i == 0 {
				continue
			}

			if !ascSet && v != elements[i-1] {
				asc = v >= elements[i-1]
			}

			if math.Abs(float64(v-elements[i-1])) > maxDiff {
				safe = false
				break
			}

			if (asc && v <= elements[i-1]) || (!asc && v >= elements[i-1]) {
				fault += 1
				fmt.Println("Faulty COUNT:", line, " with ", fault)

				if fault > 2 {
					fmt.Println("Faulty line:", line, " with ", fault)
					safe = false
					break
				}
			}
		}

		if safe {
			safeLines += 1
		}

	}
	fmt.Println(safeLines)
}
