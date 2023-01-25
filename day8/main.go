package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const TOTAL_SPACE = 70000000
const MINIMUM_AVAILABLE_SPACE = 30000000
const MAXIMUM_TOTAL_SPACE_USED = TOTAL_SPACE - MINIMUM_AVAILABLE_SPACE

func readInput() ([]int, int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	trees := make([]int, 0)
	nbTreesPerRow := 0

	for scanner.Scan() {
		text := scanner.Text()
		treeRow := make([]int, 0)
		nbTreesPerRow = len(text)

		for _, char := range text {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, 0, err
			}
			treeRow = append(treeRow, value)
		}
		trees = append(trees, treeRow...)
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}

	return trees, nbTreesPerRow, nil
}

func isInvisibleFromTop(trees []int, isTreeInvisible []bool, nbTreesPerRow int) []bool {
	currentMaximumPerColumn := make([]int, nbTreesPerRow)
	for i := 0; i < nbTreesPerRow; i++ {
		currentMaximumPerColumn[i] = trees[i]
	}
	for row := 1; row <= nbTreesPerRow-2; row++ {
		for col := 1; col <= nbTreesPerRow-2; col++ {
			if trees[nbTreesPerRow*row+col] > currentMaximumPerColumn[col] {
				isTreeInvisible[nbTreesPerRow*row+col] = false
				currentMaximumPerColumn[col] = trees[nbTreesPerRow*row+col]
			}
		}
	}

	return isTreeInvisible
}
func isInvisibleFromBottom(trees []int, isTreeInvisible []bool, nbTreesPerRow int) []bool {
	currentMaximumPerColumn := make([]int, nbTreesPerRow)
	for i := 0; i < nbTreesPerRow; i++ {
		currentMaximumPerColumn[i] = trees[i+len(trees)-nbTreesPerRow]
	}

	for row := nbTreesPerRow - 2; row > 0; row-- {
		for col := 1; col <= nbTreesPerRow-2; col++ {
			if trees[nbTreesPerRow*row+col] > currentMaximumPerColumn[col] {
				isTreeInvisible[nbTreesPerRow*row+col] = false
				currentMaximumPerColumn[col] = trees[nbTreesPerRow*row+col]
			}
		}
	}

	return isTreeInvisible
}
func isInvisibleFromLeft(trees []int, isTreeInvisible []bool, nbTreesPerRow int) []bool {
	currentMaximumPerRow := make([]int, nbTreesPerRow)
	for i := 0; i < nbTreesPerRow; i++ {
		currentMaximumPerRow[i] = trees[i*nbTreesPerRow]
	}

	for row := 1; row <= nbTreesPerRow-2; row++ {
		for col := 1; col <= nbTreesPerRow-2; col++ {
			if trees[nbTreesPerRow*row+col] > currentMaximumPerRow[row] {
				isTreeInvisible[nbTreesPerRow*row+col] = false
				currentMaximumPerRow[row] = trees[nbTreesPerRow*row+col]
			}
		}
	}

	return isTreeInvisible
}
func isInvisibleFromRight(trees []int, isTreeInvisible []bool, nbTreesPerRow int) []bool {
	currentMaximumPerRow := make([]int, len(trees))
	for i := 0; i < nbTreesPerRow; i++ {
		currentMaximumPerRow[i] = trees[i*nbTreesPerRow+nbTreesPerRow-1]
	}

	for row := nbTreesPerRow - 2; row > 0; row-- {
		for col := nbTreesPerRow - 2; col > 0; col-- {
			if trees[nbTreesPerRow*row+col] > currentMaximumPerRow[row] {
				isTreeInvisible[nbTreesPerRow*row+col] = false
				currentMaximumPerRow[row] = trees[nbTreesPerRow*row+col]
			}
		}
	}

	return isTreeInvisible
}

func getNbTreesVisibleFromOutside(trees []int, nbTreesPerRow int) int {
	// A tree is invisible if it is invisible from all directions (top, left, right, bottom).
	// We assume all trees are invisible from the start and see if there's one condition that
	//  contradicts it.
	isTreeInvisible := make([]bool, len(trees))
	for row := 0; row < nbTreesPerRow; row++ {
		for col := 0; col < nbTreesPerRow; col++ {
			// trees on the outskirt of the forest are by definition visible.
			isTreeInvisible[nbTreesPerRow*row+col] = row != 0 && col != 0 && row != nbTreesPerRow-1 && col != nbTreesPerRow-1
		}
	}

	// Instead of defining four times the function, we could have changes the "trees" input
	// by transposing it, doing a 90 degree and a 180 degree
	isTreeInvisible = isInvisibleFromTop(trees, isTreeInvisible, nbTreesPerRow)
	isTreeInvisible = isInvisibleFromBottom(trees, isTreeInvisible, nbTreesPerRow)
	isTreeInvisible = isInvisibleFromLeft(trees, isTreeInvisible, nbTreesPerRow)
	isTreeInvisible = isInvisibleFromRight(trees, isTreeInvisible, nbTreesPerRow)

	nbTreesVisible := 0
	for _, tree := range isTreeInvisible {
		if !tree {
			nbTreesVisible++
		}
	}
	return nbTreesVisible
}

func getNumberTreesFromUp(trees []int, nbTreesPerRow int, currentTreeIndex int) int {
	nbTrees := 0
	i := currentTreeIndex - nbTreesPerRow

	for i >= 0 {
		nbTrees++
		if trees[i] >= trees[currentTreeIndex] {
			return nbTrees
		}
		i -= nbTreesPerRow
	}

	return nbTrees
}

func getNumberTreesFromDown(trees []int, nbTreesPerRow int, currentTreeIndex int) int {
	nbTrees := 0
	i := currentTreeIndex + nbTreesPerRow

	for i < len(trees) {
		nbTrees++
		if trees[i] >= trees[currentTreeIndex] {
			return nbTrees
		}
		i += nbTreesPerRow
	}

	return nbTrees
}

func getNumberTreesFromLeft(trees []int, nbTreesPerRow int, currentTreeIndex int) int {
	nbTrees := 0
	i := currentTreeIndex - 1

	for i%nbTreesPerRow != (nbTreesPerRow-1) && i >= 0 {
		nbTrees++
		if trees[i] >= trees[currentTreeIndex] {
			return nbTrees
		}
		i -= 1
	}

	return nbTrees
}

func getNumberTreesFromRight(trees []int, nbTreesPerRow int, currentTreeIndex int) int {
	nbTrees := 0
	i := currentTreeIndex + 1

	for i%nbTreesPerRow != 0 && i < len(trees) {
		nbTrees++
		if trees[i] >= trees[currentTreeIndex] {
			return nbTrees
		}
		i += 1
	}

	return nbTrees
}

func getMaxScenicScore(trees []int, nbTreesPerRow int) int {
	maxScore := -1

	for index := range trees {
		scoreUp := getNumberTreesFromUp(trees, nbTreesPerRow, index)
		scoreDown := getNumberTreesFromDown(trees, nbTreesPerRow, index)
		scoreLeft := getNumberTreesFromLeft(trees, nbTreesPerRow, index)
		scoreRight := getNumberTreesFromRight(trees, nbTreesPerRow, index)

		scenicScore := scoreUp * scoreDown * scoreLeft * scoreRight

		// println("Tree ", index, " scenic score: ", scoreUp, ", ", scoreDown, ", ", scoreLeft, ", ", scoreRight)
		if scenicScore > maxScore {
			maxScore = scenicScore
		}
	}

	return maxScore
}

func main() {
	trees, nbTreesPerRow, err := readInput()
	if err != nil {
		log.Fatal("Error happened!", err)
	}
	println("Number of trees in row: ", nbTreesPerRow)

	nbTreesVisible := getNbTreesVisibleFromOutside(trees, nbTreesPerRow)
	println("Part 1: ", nbTreesVisible) // 1679

	score := getMaxScenicScore(trees, nbTreesPerRow)
	println("Part 2: ", score) // 536625
}
