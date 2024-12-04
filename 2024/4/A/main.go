package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var word = []string{"X", "M", "A", "S"}

func count_file_lines(file *os.File) (int, error) {
	defer file.Seek(0, 0)

	line_count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line_count++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error reading file: %s", err)
	}
	return line_count, nil
}

func prepare_matrix(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	file_lines, err := count_file_lines(file)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	matrix := make([][]string, file_lines)
	line_count := 0

	for scanner.Scan() {
		line := scanner.Text()
		matrix[line_count] = strings.Split(line, "")

		line_count++
	}

	return matrix, nil
}

func count_words_from(matrix [][]string, x, y int) int {

	found := 0

	y_len := len(matrix)
	x_len := len(matrix[y])

	// Right
	if 3 <= (x_len-x-1) && matrix[y][x+1] == word[1] && matrix[y][x+2] == word[2] && matrix[y][x+3] == word[3] {
		found++
	}

	// Left
	if 3 <= x && matrix[y][x-1] == word[1] && matrix[y][x-2] == word[2] && matrix[y][x-3] == word[3] {
		found++
	}

	// Top
	if 3 <= (y_len-y-1) && matrix[y+1][x] == word[1] && matrix[y+2][x] == word[2] && matrix[y+3][x] == word[3] {
		found++
	}

	// Bottom
	if 3 <= y && matrix[y-1][x] == word[1] && matrix[y-2][x] == word[2] && matrix[y-3][x] == word[3] {
		found++
	}

	// Diagonal (right, bottom)
	if 3 <= (x_len-x-1) && 3 <= (y_len-y-1) && matrix[y+1][x+1] == word[1] && matrix[y+2][x+2] == word[2] && matrix[y+3][x+3] == word[3] {
		found++
	}

	// Diagonal (right, top)
	if 3 <= (x_len-x-1) && 3 <= y && matrix[y-1][x+1] == word[1] && matrix[y-2][x+2] == word[2] && matrix[y-3][x+3] == word[3] {
		found++
	}

	// Diagonal (left, bottom)
	if 3 <= x && 3 <= (y_len-y-1) && matrix[y+1][x-1] == word[1] && matrix[y+2][x-2] == word[2] && matrix[y+3][x-3] == word[3] {
		found++
	}

	// Diagonal (left, top)
	if 3 <= x && 3 <= y && matrix[y-1][x-1] == word[1] && matrix[y-2][x-2] == word[2] && matrix[y-3][x-3] == word[3] {
		found++
	}

	return found
}

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("Please provide a file name as an argument"))
	}

	matrix, err := prepare_matrix(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("Can not prepare matrix: %s", err))
	}

	found_words := 0
	for y, yv := range matrix {
		for x, xv := range yv {
			if xv == word[0] {
				found_words += count_words_from(matrix, x, y)
			}
		}
	}

	fmt.Println(found_words)
}
