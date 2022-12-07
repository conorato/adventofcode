package main

import (
	"bufio"
	"log"
	"os"
)

func readInput() (string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	datastreamBuffer := scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return datastreamBuffer, nil
}

func isMarker(sequence string) bool {
	set := make(map[rune]struct{})

	for _, character := range sequence {
		if _, exist := set[character]; exist {
			return false
		} else {
			set[character] = struct{}{}
		}
	}
	return true
}

func getNumberCharactersBeforeFirstStartOfPacket(datastreamBuffer string, sequenceSize int) int {
	for i := sequenceSize; i < len(datastreamBuffer); i++ {
		sequence := datastreamBuffer[i-sequenceSize : i]

		if isMarker(sequence) {
			return i
		}
	}
	return 0
}

func main() {
	datastreamBuffer, err := readInput()
	if err != nil {
		log.Fatal("Error happened!", err)
	}

	nbCharacters := getNumberCharactersBeforeFirstStartOfPacket(datastreamBuffer, 4)
	println("Part 1: ", nbCharacters)

	nbCharacters = getNumberCharactersBeforeFirstStartOfPacket(datastreamBuffer, 14)
	println("Part 2: ", nbCharacters)
}
