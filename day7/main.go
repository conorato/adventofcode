package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const TOTAL_SPACE = 70000000
const MINIMUM_AVAILABLE_SPACE = 30000000
const MAXIMUM_TOTAL_SPACE_USED = TOTAL_SPACE - MINIMUM_AVAILABLE_SPACE

type INode interface {
	getName() string
	getSize() int
}

type File struct {
	name string
	size int
}

type Directory struct {
	name     string
	children *[]INode
	parent   INode
}

func (f File) getName() string {
	return f.name
}

func (d Directory) getName() string {
	return d.name
}

func (f File) getSize() int {
	return f.size
}

func (d Directory) getSize() int {
	size := 0

	for _, child := range *d.children {
		size += child.getSize()
	}

	return size
}

func changeDirectory(currentDirectory Directory, root Directory, destination string) Directory {
	if destination == "/" {
		return root
	} else if destination == ".." {
		return currentDirectory.parent.(Directory)
	}
	for _, child := range *currentDirectory.children {
		if child.getName() == destination {
			return (child).(Directory)
		}
	}
	panic("")
}

func readInput() (*Directory, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	children := make([]INode, 0)

	root := Directory{
		name:     "/",
		parent:   nil,
		children: &children,
	}
	currentDirectory := root

	for scanner.Scan() {
		text := scanner.Text()
		if text[:len("$ cd")] == "$ cd" {
			destination := text[len("$ cd "):]
			currentDirectory = changeDirectory(currentDirectory, root, destination)
		}

		if text[0] != '$' {
			var currentNode INode
			if text[:len("dir")] == "dir" {
				children := make([]INode, 0)
				currentNode = Directory{
					name:     text[len("dir "):],
					children: &children,
					parent:   currentDirectory,
				}
			} else {
				sizeAndName := strings.Split(text, " ")
				size, err := strconv.Atoi(sizeAndName[0])
				if err != nil {
					return nil, err
				}
				currentNode = File{
					size: size,
					name: sizeAndName[1],
				}
			}

			*currentDirectory.children = append(*currentDirectory.children, currentNode)
		}
		println(text, "current state:  ", currentDirectory.name, len(*currentDirectory.children), "child ref: ", currentDirectory.children)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &root, nil
}

func getSum(directory Directory) int {
	curentDirectorySize := 0
	totalSmallDirectorySize := 0

	for _, child := range *directory.children {
		curentDirectorySize += child.getSize()

		if _, ok := child.(Directory); ok {
			totalSmallDirectorySize += getSum(child.(Directory))
		}
	}

	if curentDirectorySize > 100000 {
		return totalSmallDirectorySize
	}
	return curentDirectorySize + totalSmallDirectorySize
}

func getSmallestDirectory(directory Directory, currentSpaceToFreeUp int, spaceToFreeUp int) int {
	for _, child := range *directory.children {
		if childDirectory, ok := child.(Directory); ok {
			size := childDirectory.getSize()
			if size >= spaceToFreeUp && size < currentSpaceToFreeUp {
				currentSpaceToFreeUp = size
			}

			currentSpaceToFreeUp = getSmallestDirectory(childDirectory, currentSpaceToFreeUp, spaceToFreeUp)
		}
	}

	return currentSpaceToFreeUp
}

func getSmallestDirectoryToFreeUpEnoughSpace(root Directory) int {
	totalUsedSpace := root.getSize()
	if totalUsedSpace < MAXIMUM_TOTAL_SPACE_USED {
		println("No need to delete directory")
		return 0
	}

	spaceToFreeUp := totalUsedSpace - MAXIMUM_TOTAL_SPACE_USED
	println("Total used space: ", totalUsedSpace, " Total space to free up: ", spaceToFreeUp)

	return getSmallestDirectory(root, totalUsedSpace, spaceToFreeUp)
}

func main() {
	root, err := readInput()
	if err != nil {
		log.Fatal("Error happened!", err)
	}
	println(root.getName(), len(*root.children))

	sum := getSum(*root)
	println("Part 1: ", sum)

	size := getSmallestDirectoryToFreeUpEnoughSpace(*root)
	println("Part 2: ", size)
}
