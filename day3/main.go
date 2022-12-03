package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const NUMBER_OF_ELF_PER_GROUP = 3

type Rucksack string

func readInput() ([]Rucksack, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var rucksacks []Rucksack

	for scanner.Scan() {
		text := scanner.Text()
		rucksacks = append(rucksacks, Rucksack(text))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rucksacks, nil
}

func getItemInTwoCompartment(rucksack Rucksack) rune {
	firstCompartmentItems := rucksack[0 : len(rucksack)/2]
	secondCompartmentItems := rucksack[len(rucksack)/2:]
	for _, firstCompartmentItem := range firstCompartmentItems {
		for _, secondCompartmentItem := range secondCompartmentItems {
			if firstCompartmentItem == secondCompartmentItem {
				return firstCompartmentItem
			}
		}
	}
	panic("No duplicate item found :(")
}

func getBadge(rucksacks []Rucksack) rune {
	// Complexity O(n+m+p) where n,m,p are the number of items in the elves' rucksacks.
	// Could be improved by finding list of common runes between 2 first rucksacks and then
	// reusing the common runes between 1 and 2 with the third -> kind off O(2n^2)
	// Considering the size of a rucksack is relatively small, improving performance doesn't really matter here ;)
	for _, item1 := range rucksacks[0] {
		for _, item2 := range rucksacks[1] {
			for _, item3 := range rucksacks[2] {
				if item1 == item2 && item2 == item3 {
					return item1
				}
			}
		}
	}
	panic("No duplicate item found :(")
}

func getPriority(item rune) int {
	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}

	return int(item - 'a' + 1)
}

func sumPrioritiesItemTwoCompartment(rucksacks []Rucksack) int {
	sum := 0

	for _, rucksack := range rucksacks {
		item := getItemInTwoCompartment(rucksack)
		priority := getPriority(item)
		sum += priority
	}

	return sum
}

func sumPrioritiesBadges(rucksacks []Rucksack) int {
	sum := 0

	for i := 0; i < len(rucksacks); i += NUMBER_OF_ELF_PER_GROUP {
		item := getBadge(rucksacks[i : i+NUMBER_OF_ELF_PER_GROUP])
		priority := getPriority(item)
		sum += priority
	}

	return sum
}

func main() {
	rucksacks, err := readInput()
	if err != nil {
		log.Fatal("Error happened!", err)
	}
	fmt.Println("Total number of rucksacks: ", len(rucksacks))

	sumPriorities := sumPrioritiesItemTwoCompartment(rucksacks)
	fmt.Println("Part 1: ", sumPriorities)

	sumPrioritiesBadge := sumPrioritiesBadges(rucksacks)
	fmt.Println("Part 2: ", sumPrioritiesBadge)
}
