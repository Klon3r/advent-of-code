package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func getSlice(array []string, bottom int, top int) []string{
	return array[bottom:top+1]
}

func getArrayValues(array []string, index int) []string {
	if(len(array) <= 0) {
		return []string{}
	}

	if(index - 1 < 0) {
		return getSlice(array, index, index + 1)
	} else if (index + 1 >= len(array)) {
		return getSlice(array, index-1, index)
	} else {
		return getSlice(array, index-1, index+1)
	}
}

func checkAdjacentValues(index int, currentArray []string,  prevArray []string, nextArray []string) int {
	count := 0

	// Top row
	if len(prevArray) > 0 {
		topValues := getArrayValues(prevArray, index)
		for _,value := range topValues {
			if value == "@" {
				count++
			}
		}
	}

	// Middle Row
	if index > 0 && currentArray[index-1] == "@" {
		count++
	}
	if index < len(currentArray)-1 && currentArray[index+1] == "@" {
		count++
	}

	// Bottom Row
	if len(nextArray) > 0 {
		bottomValues := getArrayValues(nextArray, index)
		for _, value := range bottomValues {
			if value == "@" {
				count++
			}
		}
	}

	return count
}

func main() {
	color.Set(color.FgHiMagenta, color.Bold)
	fmt.Println("___________________")
	fmt.Println("04 - Advent Of Code")
	color.Unset()

	fileName := "04-input.txt"
	fileArray := utils.ReadFile(fileName)

	partOneTotal := 0

	// Part One
	for i := 0; i < len(fileArray); i++ {
		currentArray := strings.Split(fileArray[i], "")
		var nextArray []string 
		var prevArray []string

		// Get prev/next array if applicable
		if i+1 < len(fileArray) {
			nextArray = strings.Split(fileArray[i+1], "")
		}

		if i-1 >= 0 {
			prevArray = strings.Split(fileArray[i-1], "")
		}

		for j := 0; j < len(currentArray); j++ {
			currentValue := currentArray[j]
			
			if(currentValue == "@") {
				// Check adjacent value function
				numOfRollsAround := checkAdjacentValues(j, currentArray, prevArray, nextArray)
				if(numOfRollsAround < 4) {
					partOneTotal += 1
				} 
			}
		}
	}


	// Part Two
	// Create a deep copy of the 2D array
	copyArray := make([][]string, len(fileArray))
	for i := range fileArray {
		copyArray[i] = strings.Split(fileArray[i], "")
	}

	partTwoTotal := 0
	removedInPass := 0

	for {
		removedInPass = 0

		for x := 0; x < len(copyArray); x++ {
			var prevArray []string
			var nextArray []string

			if x+1 < len(copyArray) {
				nextArray = copyArray[x+1]
			}
			if x-1 >= 0 {
					prevArray = copyArray[x-1]
			}

			for z := 0; z < len(copyArray[x]); z++ {
				currentValue := copyArray[x][z]

				if currentValue == "@" {
					numOfRollsAround := checkAdjacentValues(z, copyArray[x], prevArray, nextArray)
					if numOfRollsAround < 4 {
						// Remove this roll
						copyArray[x][z] = "."
						partTwoTotal++
						removedInPass++
					}
				}
			}
		}

		if removedInPass == 0 {
			break
		}
	}

	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(partOneTotal)))
	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(partTwoTotal)))
}