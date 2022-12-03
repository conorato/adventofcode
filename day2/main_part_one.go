package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Shape string

const (
	Rock    Shape = "Rock"
	Paper         = "Paper"
	Scissor       = "Scissor"
)

var stringToShape = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissor,
	"X": Rock,
	"Y": Paper,
	"Z": Scissor,
}

var shapeToScore = map[Shape]int{
	Rock:    1,
	Paper:   2,
	Scissor: 3,
}

type Round struct {
	myShape, opponentShape Shape
}

func readInput() ([]Round, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var rounds []Round

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		currentRound := Round{
			opponentShape: stringToShape[text[0]],
			myShape:       stringToShape[text[1]],
		}
		rounds = append(rounds, currentRound)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rounds, nil
}

func getRoundOutcomeScore(round Round) int {
	draw := round.myShape == round.opponentShape
	if draw {
		return 3
	}

	win := (round.myShape == Rock && round.opponentShape == Scissor) ||
		(round.myShape == Scissor && round.opponentShape == Paper) ||
		(round.myShape == Paper && round.opponentShape == Rock)
	if win {
		return 6
	}

	return 0
}

func calculateScoreAccordingToPlan(rounds []Round) int {
	score := 0

	for _, round := range rounds {
		currentScore := shapeToScore[round.myShape] + getRoundOutcomeScore(round)
		score += currentScore
	}

	return score
}

func main() {
	rounds, err := readInput()
	if err != nil {
		log.Fatal("Error happened!", err)
	}

	fmt.Println("Total number of rounds: ", len(rounds))
	fmt.Println("Round 1: myShape ", rounds[0].myShape, " vs ", rounds[1].opponentShape)

	score := calculateScoreAccordingToPlan(rounds)
	fmt.Println("Part 1: score is ", score)
}
