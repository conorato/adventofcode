package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Shape string
type Outcome string

const (
	Rock    Shape = "Rock"
	Paper         = "Paper"
	Scissor       = "Scissor"
)

const (
	Lose Outcome = "Lose"
	Draw         = "Draw"
	Win          = "Win"
)

var stringToShape = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissor,
}

var stringToOutcome = map[string]Outcome{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

var shapeToScore = map[Shape]int{
	Rock:    1,
	Paper:   2,
	Scissor: 3,
}

var outcomeToScore = map[Outcome]int{
	Lose: 0,
	Draw: 3,
	Win:  6,
}

type Round struct {
	opponentShape Shape
	outcome       Outcome
}

var opponentShapeToWinShape = map[Shape]Shape{
	Rock:    Paper,
	Paper:   Scissor,
	Scissor: Rock,
}

var opponentShapeToLoseShape = map[Shape]Shape{
	Paper:   Rock,
	Rock:    Scissor,
	Scissor: Paper,
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
			outcome:       stringToOutcome[text[1]],
		}
		rounds = append(rounds, currentRound)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rounds, nil
}

func getShape(round Round) Shape {
	if round.outcome == Draw {
		return round.opponentShape
	}

	if round.outcome == Win {
		return opponentShapeToWinShape[round.opponentShape]
	}

	return opponentShapeToLoseShape[round.opponentShape]
}

func calculateScore(rounds []Round) int {
	score := 0

	for _, round := range rounds {
		myShape := getShape(round)
		currentScore := shapeToScore[myShape] + outcomeToScore[round.outcome]
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
	fmt.Println("Round 1: opponent ", rounds[1].opponentShape, " and outcome: ", rounds[0].outcome)

	score := calculateScore(rounds)
	fmt.Println("Part 2: score is ", score)
}
