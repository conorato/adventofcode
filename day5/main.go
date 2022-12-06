package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type Crate rune

type RearrangementStep struct {
	numberCrates     int
	stackOrigin      int
	stackDestination int
}

func listsToStacks(stacksSlices [][]Crate) []stack.Stack {
	stacks := make([]stack.Stack, len(stacksSlices))

	for stackIndex, stackSlice := range stacksSlices {
		for crateIndex := len(stackSlice) - 1; crateIndex >= 0; crateIndex-- {
			stacks[stackIndex].Push(stackSlice[crateIndex])
		}
	}

	return stacks
}

func readStacks(scanner *bufio.Scanner) []stack.Stack {
	var stacks [][]Crate
	var numberOfStacks int
	isInitializing := true

	for scanner.Scan() {
		text := scanner.Text()
		if text[1] == '1' {
			break
		}

		if isInitializing {
			numberOfStacks = (len(text) + 1) / 4
			stacks = make([][]Crate, numberOfStacks)
			isInitializing = false
		}

		for i := 0; i < numberOfStacks; i++ {
			crateId := text[4*i+1]
			if crateId != ' ' {
				stacks[i] = append(stacks[i], Crate(crateId))
			}
		}
	}

	return listsToStacks(stacks)
}

func readRearrengmentProcedure(scanner *bufio.Scanner) ([]RearrangementStep, error) {
	var rearrangementProcedure []RearrangementStep
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		splittedText := strings.Split(text, " ")
		numberCrates, err := strconv.Atoi(splittedText[1])
		if err != nil {
			return nil, err
		}
		stackOrigin, err := strconv.Atoi(splittedText[3])
		if err != nil {
			return nil, err
		}
		stackDestination, err := strconv.Atoi(splittedText[5])
		if err != nil {
			return nil, err
		}

		rearrangementProcedure = append(rearrangementProcedure, RearrangementStep{
			numberCrates:     numberCrates,
			stackOrigin:      stackOrigin,
			stackDestination: stackDestination,
		})
	}
	return rearrangementProcedure, nil
}

func readInput() ([]stack.Stack, []RearrangementStep, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	stacks := readStacks(scanner)
	rearrangementProcedure, err := readRearrengmentProcedure(scanner)
	if err != nil {
		return nil, nil, err
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return stacks, rearrangementProcedure, nil
}

func rearrangeStacksCrateMove9000(stacks []stack.Stack, rearrangementProcedure []RearrangementStep) string {
	for _, step := range rearrangementProcedure {
		for i := 0; i < step.numberCrates; i++ {
			crate := stacks[step.stackOrigin-1].Pop().(Crate)
			stacks[step.stackDestination-1].Push(crate)
		}
	}

	topLevelCrate := ""
	for _, stack := range stacks {
		topLevelCrate = topLevelCrate + string(stack.Peek().(Crate))
	}
	return topLevelCrate
}

func rearrangeStacksCrateMove9001(stacks []stack.Stack, rearrangementProcedure []RearrangementStep) string {
	// ...... I'm not redoing that readInput function, so here's this monstruosity

	for _, step := range rearrangementProcedure {
		crates := make([]Crate, step.numberCrates)
		for i := 0; i < step.numberCrates; i++ {
			crates[step.numberCrates-i-1] = stacks[step.stackOrigin-1].Pop().(Crate)
		}
		for _, crate := range crates {
			stacks[step.stackDestination-1].Push(crate)
		}
	}

	topLevelCrate := ""
	for _, stack := range stacks {
		topLevelCrate = topLevelCrate + string(stack.Peek().(Crate))
	}
	return topLevelCrate
}

func main() {
	stacks, rearrangementProcedure, _ := readInput()

	topLevelCrates9000 := rearrangeStacksCrateMove9000(stacks, rearrangementProcedure)
	println("Part 1: ", topLevelCrates9000)

	stacks, rearrangementProcedure, _ = readInput()
	topLevelCrates9001 := rearrangeStacksCrateMove9001(stacks, rearrangementProcedure)
	println("Part 2: ", topLevelCrates9001)
}
