package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	distanceList := make([]int, 0)
	sumDistance := 0
	similarityScore := 0

	file, err := os.Open("./day1/data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// processing each line
		line := strings.Fields(scanner.Text())
		for i, v := range line {
			num, _ := strconv.Atoi(v)
			if i == 0 {
				leftList = append(leftList, num)
			} else {
				rightList = append(rightList, num)
			}
		}
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i := 0; i < len(leftList); i++ {
		sum := leftList[i] - rightList[i]
		if sum < 0 {
			sum = -sum
		}
		distanceList = append(distanceList, sum)
	}

	for _, v := range distanceList {
		sumDistance = sumDistance + v
	}

	for _, left := range leftList {
		similarCount := 0
		for _, right := range rightList {
			if left == right {
				similarCount++
			}
		}
		similarityScore = similarityScore + (left * similarCount)
	}

	fmt.Println("The total distance and the answer for day 1, part 1 is: ", sumDistance)
	fmt.Println("The similarity score and the answer for day 1, part 2 is: ", similarityScore)
}
