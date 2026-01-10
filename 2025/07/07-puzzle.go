package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.Set(color.FgHiMagenta, color.Bold)
	fmt.Println("___________________")
	fmt.Println("07 - Advent Of Code")
	color.Unset()

	fileName := "07-input.txt"
	partOneResult := partOne(fileName)
	partTwoResult := partTwo(fileName)

	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(partOneResult)))
	fmt.Println("Answer (Part Two):", color.GreenString(strconv.Itoa(partTwoResult)))
}

func partOne(inputFile string) int{
	fileArray := utils.ReadFile(inputFile)

	array := make([][]string, len(fileArray))
	numArray := make([][]int, len(fileArray))

	for index, row := range fileArray {
		rowSplit := strings.Split(row, "")
		array[index] = rowSplit
		numArray[index] = make([]int, len(rowSplit)) 
	}

	// Find the start cell
	currentCell := -1
	for i := 0; i < len(array[0]); i++ {
		if(array[0][i] == "S") {
			currentCell = i
			numArray[0][i] = 1
		}
	}

	timesSplit := 0
	partOneMoveBeam(array, 1, currentCell, &timesSplit)

	return timesSplit
}

func partOneMoveBeam(array [][]string, index int, currentCell int, timeSplit *int) {
	loop:
		for x := index; x < len(array); x++ {
			switch array[x][currentCell] {
			case ".":
				array[x][currentCell] = "|"
			case "^":
				array[x][currentCell-1] = "|"
				array[x][currentCell+1] = "|"

				*timeSplit++

				array[x][currentCell] = "X"

				partOneMoveBeam(array, x, currentCell-1, timeSplit)
				partOneMoveBeam(array, x, currentCell+1, timeSplit)
				break loop

			case "X":
				break loop

			default:
				continue
			}
		}
}

func partTwo(inputFile string) int {
	fileArray := utils.ReadFile(inputFile)

	array := make([][]string, len(fileArray))
	numArray := make([][]int, len(fileArray))

	for index, row := range fileArray {
		rowSplit := strings.Split(row, "")
		array[index] = rowSplit
		numArray[index] = make([]int, len(rowSplit)) 
	}

	// Find the start cell
	currentCell := -1
	for i := 0; i < len(array[0]); i++ {
		if(array[0][i] == "S") {
			currentCell = i
			numArray[0][i] = 1
		}
	}

	partTwoTotal := countTimelines(array, currentCell)

	return partTwoTotal
}

func countTimelines(array [][]string, startCol int) int {
	height := len(array)
	width := len(array[0])

	dp := make([][]int, height)
	for i := range dp {
		// Allocate the [][]int array
		dp[i] = make([]int, width)
	}

	// Bottom row (Set each cell as 1)
	for col := 0; col < width; col++ {
		dp[height-1][col] = 1
	}

	// Bottom to top
	for row := height - 2; row >= 0; row -- {
		for col := 0; col < width; col++ {
			switch array[row][col] {
			case ".", "S": 
				dp[row][col] = dp[row+1][col]
			case "^":
				left := 1
				right := 1

				if col > 0 {
					left = dp[row+1][col-1]
				}

				if col < width-1 {
					right = dp[row+1][col+1]
				}
				dp[row][col] = left + right
			}
		}
	}

	return dp[0][startCol]
}
