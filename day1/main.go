package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const INPUT_FILE_NAME = "input.txt"

func getCaloriesByElf() ([]int, error) {
	f, err := os.Open(INPUT_FILE_NAME)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var elfCalories []int
	currentElfCalories := 0

	for scanner.Scan() {
		currentCaloriesStr := scanner.Text()

		if currentCaloriesStr != "" {
			currentCalories, err := strconv.Atoi(currentCaloriesStr)

			if err != nil {
				return nil, err
			}

			currentElfCalories += currentCalories
		} else {
			elfCalories = append(elfCalories, currentElfCalories)
			currentElfCalories = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return elfCalories, nil
}

func getMaxIndex(elfCalories []int) int {
	maxCaloriesIdx := 0

	for i := range elfCalories {
		if elfCalories[i] > elfCalories[maxCaloriesIdx] {
			maxCaloriesIdx = i
		}
	}

	return maxCaloriesIdx
}

// part 1
func getBiggestCaloriesOneElf(elfCalories []int) int {
	index := getMaxIndex(elfCalories)

	return elfCalories[index]
}

func isDifferentFromTopElves(topElvesIndexes []int, index int) bool {
	for _, topElveIndex := range topElvesIndexes {
		if topElveIndex == index {
			return false
		}
	}
	return true
}

// part 2
func getBiggestCaloriesThreeElves(elfCalories []int) int {
	topElvesCaloryIndexes := []int{-1, -1, -1}

	for topElfIndex := range topElvesCaloryIndexes {

		currentTopElfIndex := 0
		for i := range elfCalories {
			isDifferentFromPreviousElves := isDifferentFromTopElves(topElvesCaloryIndexes, i)
			if elfCalories[i] > elfCalories[currentTopElfIndex] && isDifferentFromPreviousElves {
				currentTopElfIndex = i
			}
		}

		topElvesCaloryIndexes[topElfIndex] = currentTopElfIndex
	}

	totalCalories := 0
	for _, topElfCaloryIndex := range topElvesCaloryIndexes {
		totalCalories += elfCalories[topElfCaloryIndex]
	}

	return totalCalories
}

func main() {
	elfCalories, err := getCaloriesByElf()
	if err != nil {
		log.Fatal("Error happened!", err)
	}

	biggestCaloriesOneElf := getBiggestCaloriesOneElf(elfCalories)
	biggestCaloriesThreeElves := getBiggestCaloriesThreeElves(elfCalories)

	fmt.Println("Part 1, biggest number of calory by 1 elf: ", biggestCaloriesOneElf)
	fmt.Println("Part 2, biggest number of calory by 3 elves: ", biggestCaloriesThreeElves)
}
