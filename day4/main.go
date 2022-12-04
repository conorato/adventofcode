package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CleaningAssignment struct {
	sectionIdStart int
	sectionIdEnd   int
}

type PairCleaningAssignment []CleaningAssignment

func readInput() ([]PairCleaningAssignment, error) {
	f, err := os.Open("input_test.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var pairCleaningAssignments []PairCleaningAssignment

	for scanner.Scan() {
		pairOfRanges := strings.Split(scanner.Text(), ",")
		var pairCleaningAssignment PairCleaningAssignment

		for _, sectionRange := range pairOfRanges {
			sectionDelimitersStr := strings.Split(sectionRange, "-")
			var sectionDelimiters []int
			for _, sectionDelimiterStr := range sectionDelimitersStr {
				sectionDelimiter, err := strconv.Atoi(sectionDelimiterStr)
				if err != nil {
					return nil, err
				}
				sectionDelimiters = append(sectionDelimiters, sectionDelimiter)
			}

			pairCleaningAssignment = append(pairCleaningAssignment,
				CleaningAssignment{
					sectionIdStart: sectionDelimiters[0],
					sectionIdEnd:   sectionDelimiters[1],
				})
		}
		pairCleaningAssignments = append(pairCleaningAssignments, pairCleaningAssignment)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return pairCleaningAssignments, nil
}

func isAssignment1ContainingAssignment2(assignment1 CleaningAssignment, asssignment2 CleaningAssignment) bool {
	return assignment1.sectionIdStart <= asssignment2.sectionIdStart &&
		assignment1.sectionIdEnd >= asssignment2.sectionIdEnd
}

func getNumberPairsOneElfFullyContainsOther(pairCleaningAssignments []PairCleaningAssignment) int {
	total := 0

	for _, pairCleaningAssignment := range pairCleaningAssignments {
		elfOne := pairCleaningAssignment[0]
		elfTwo := pairCleaningAssignment[1]
		elfOneFullyContainsElfTwoSection := isAssignment1ContainingAssignment2(elfOne, elfTwo)
		elfTwoFullyContainsElfoneSection := isAssignment1ContainingAssignment2(elfTwo, elfOne)
		if elfOneFullyContainsElfTwoSection || elfTwoFullyContainsElfoneSection {
			total++
		}
	}

	return total
}

func areAssignmentsOverlapping(assignment1 CleaningAssignment, asssignment2 CleaningAssignment) bool {
	if assignment1.sectionIdStart <= asssignment2.sectionIdStart {
		return assignment1.sectionIdEnd >= asssignment2.sectionIdStart
	}

	return asssignment2.sectionIdEnd >= assignment1.sectionIdStart
}

func getNumberPairsOneElfOverlapsOther(pairCleaningAssignments []PairCleaningAssignment) int {
	total := 0

	for _, pairCleaningAssignment := range pairCleaningAssignments {
		elfOne := pairCleaningAssignment[0]
		elfTwo := pairCleaningAssignment[1]
		if areAssignmentsOverlapping(elfOne, elfTwo) {
			total++
		}
	}

	return total
}

func main() {
	pairCleaningAssignments, err := readInput()
	if err != nil {
		log.Fatal("Error happened!", err)
	}
	fmt.Println("Total number of pairs: ", len(pairCleaningAssignments))

	numberPairsOneElfFullyContainsOther := getNumberPairsOneElfFullyContainsOther(pairCleaningAssignments)
	fmt.Println("Part 1: ", numberPairsOneElfFullyContainsOther)

	numberPairsOneElfOverlapsOther := getNumberPairsOneElfOverlapsOther(pairCleaningAssignments)
	fmt.Println("Part 2: ", numberPairsOneElfOverlapsOther)
}
