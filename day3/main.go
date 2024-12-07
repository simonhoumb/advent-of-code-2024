package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func readAndScanFile(fileLocation string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	return file, bufio.NewScanner(file)
}

func main() {
	_, scanner := readAndScanFile("C:\\Programming_Projects\\Personal\\advent-of-code-2024\\day3\\data.txt")

	for scanner.Scan() {
		line := scanner.Text()
		reg, err := regexp.Compile("mul(%d,%d)")
		if err != nil {
			fmt.Println("Error when trying to compile regex string: ", err)
		}
		foundMatches := reg.FindAllString(line, -1)
		if len(foundMatches) == 0 {
			fmt.Println("No matches found in line.")
		} else {
			fmt.Printf("Found matches: %v", foundMatches)
		}
 	}
}