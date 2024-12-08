package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to check if a report is safe without modification
func isSafe(report []int) bool {
	if len(report) < 2 {
		return true // A single number is trivially safe
	}

	// Determine if the sequence is increasing or decreasing based on the first pair
	firstDiff := report[1] - report[0]
	if firstDiff == 0 {
		return false // If the first two elements are equal, it's unsafe
	}

	// If the first pair is increasing, ensure the rest of the sequence increases
	if firstDiff > 0 {
		for i := 1; i < len(report); i++ {
			if report[i] <= report[i-1] {
				return false // Found a decrease or equal pair in an increasing sequence
			}
			diff := report[i] - report[i-1]
			if diff < 1 || diff > 3 {
				return false // The difference is outside the allowed range
			}
		}
	} else { // If the first pair is decreasing
		for i := 1; i < len(report); i++ {
			if report[i] >= report[i-1] {
				return false // Found an increase or equal pair in a decreasing sequence
			}
			diff := report[i-1] - report[i]
			if diff < 1 || diff > 3 {
				return false // The difference is outside the allowed range
			}
		}
	}
	return true
}

// Function to check if a report can become safe by removing one level
func canBecomeSafeByRemovingOne(report []int) bool {
	for i := 0; i < len(report); i++ {
		// Create a new report by removing the level at index i
		newReport := append([]int{}, report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		// Check if the new report is safe
		if isSafe(newReport) {
			return true
		}
	}
	return false
}

func readFile(fileLocation string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	return file, bufio.NewScanner(file)
}

func main() {
	file, scanner := readFile("C:\\Programming_Projects\\Personal\\advent-of-code-2024\\day2\\data.txt")
	defer file.Close()


	// Variable to count safe reports
	safeCount := 0
	for scanner.Scan() {
		// Parse the report levels into an integer slice
		reportStrs := strings.Fields(scanner.Text())
		report := make([]int, len(reportStrs))
		for i, s := range reportStrs {
			num, _ := strconv.Atoi(s)
			report[i] = num
		}

		fmt.Printf("Report=%v", report)
		// Check if the report is safe with or without removing one level
		if isSafe(report) || canBecomeSafeByRemovingOne(report) {
			safeCount++
			fmt.Printf(" is SAFE => Total Safe: %d\n", safeCount)
		} else {
			fmt.Printf("is UNSAFE => Total Safe: %d\n", safeCount)
		}
	}

	// Print the number of safe reports
	fmt.Println("Number of safe reports:", safeCount)
}
