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
	fileArray := utils.ReadFile(fileName)

	array := make([][]string, len(fileArray))

	for index, row := range fileArray {
		rowSplit := strings.Split(row, "")
		array[index] = rowSplit
	}

	// Find the start cell
	currentCell := -1
	for i := 0; i < len(array[0]); i++ {
		if(array[0][i] == "S") {
			currentCell = i
		}
	}

	timesSplit := 0
	moveBeam(array, 1, currentCell, &timesSplit)

	// Print Array
	// for _, row := range array {
	// 	fmt.Println(row)
	// }

	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(timesSplit)))
}

func moveBeam(array [][]string, index int, currentCell int, timeSplit *int) {
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

				moveBeam(array, x, currentCell-1, timeSplit)
				moveBeam(array, x, currentCell+1, timeSplit)
				break loop

			case "X":
				break loop
			default:
				continue
			}
		}
}