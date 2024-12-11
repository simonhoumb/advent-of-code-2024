package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Load the grid from input
	file, scanner := readAndScanFile("C:\\Programming_Projects\\Personal\\advent-of-code-2024\\day4\\data.txt")
	defer file.Close()
	
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	word := "XMAS"
	count := countWordOccurrences(grid, word)
	fmt.Println("Total occurrences of \"XMAS\" and the answer to day 1, part1 is: ", count)
}

func readAndScanFile(fileLocation string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	return file, bufio.NewScanner(file)
}

func countWordOccurrences(grid []string, word string) int {
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},  // Vertical and horizontal
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Diagonal
	}
	wordLen := len(word)
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if matches(grid, word, r, c, dir, wordLen, rows, cols) {
					count++
				}
			}
		}
	}
	return count
}

func matches(grid []string, word string, r, c int, dir [2]int, wordLen, rows, cols int) bool {
	for i := 0; i < wordLen; i++ {
		nr, nc := r+i*dir[0], c+i*dir[1]
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != word[i] {
			return false
		}
	}
	return true
}
