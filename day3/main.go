package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readAndScanFile(fileLocation string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	return file, bufio.NewScanner(file)
}

func sliceAtoi(slice []string) []int {
	newSlice := make([]int, len(slice))
	for i, s := range slice {
		num, _ := strconv.Atoi(s)
		newSlice[i] = num
	}
	return newSlice
}

func findMuls(text string) []string {
	reg, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	if err != nil {
		fmt.Println("Error while trying to compile regex: ", err)
	}
	muls := reg.FindAllString(text, -1)
	return muls
}

func findMulsWithInstructions(text string, mulEnabled bool) ([]string, bool) {
	var muls []string
	reg, err := regexp.Compile(`don't\(\)|do\(\)|mul\([0-9]+,[0-9]+\)`)
	if err != nil {
		fmt.Println("Error while trying to compile regex: ", err)
	}

	instructions := reg.FindAllString(text, -1)
	for _, instruction := range instructions {
		if instruction == "don't()" {
			mulEnabled = false
		} else if instruction == "do()" {
			mulEnabled = true
		} else if mulEnabled {
			muls = append(muls, instruction)
		}
	}
	return muls, mulEnabled
}

func calcAllMuls(muls []string) int {
	answer := 0
	for _, mul := range muls {
		reg := regexp.MustCompile(`[0-9]+`)
		stringNums := reg.FindAllString(mul, -1)
		nums := sliceAtoi(stringNums)
		answer += nums[0] * nums[1]
	}
	return answer
}

func main() {
	_, scanner := readAndScanFile("C:\\Programming_Projects\\Personal\\advent-of-code-2024\\day3\\data.txt")
	resultPart1 := 0
	resultPart2 := 0
	mulEnabled := true

	for scanner.Scan() {
		var foundMulsPart2 []string
		line := scanner.Text()
		foundMulsPart1 := findMuls(line)
		resultPart1 += calcAllMuls(foundMulsPart1)
		foundMulsPart2, mulEnabled = findMulsWithInstructions(line, mulEnabled)
		resultPart2 += calcAllMuls(foundMulsPart2)
 	}

	fmt.Println("The result of all multiplications and the answer for day 3, part 1 is: ", resultPart1)
	fmt.Println("The result of all multiplications and the answer for day 3, part 2 is: ", resultPart2)
}