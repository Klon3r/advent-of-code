package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.Set(color.FgHiMagenta, color.Bold)
	fmt.Println("___________________")
	fmt.Println("01 - Advent Of Code")
	color.Unset()

	fileName := "01-input.txt";
	
	fileArray := utils.ReadFile(fileName)

	dialNumber := 50
	timesOnZero := 0
	timesPastZero := 0

	for _, value := range fileArray {
		splitValue := strings.Split(value, "")
		operator := splitValue[0]
		number := strings.Join(splitValue[1:], "") 

		integer, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal("Failed to convert string to integer", err)
		}
		
		if operator == "R" {
			dialNumber = operation(true, dialNumber, integer, &timesPastZero, &timesOnZero)
		} else {
			dialNumber = operation(false, dialNumber, integer, &timesPastZero, &timesOnZero)
		}
	}

	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(timesOnZero)))
	fmt.Println("Answer (Part Two):", color.GreenString(strconv.Itoa(timesPastZero)))
}

// Operation simulates turning the dial left (subtract) or right (add) by the given value.
// It wraps the dial from 0 to 99 (for add) or 99 to 0 (for subtract)
func operation(add bool, dialNumber int, value int, timesPastZero *int, timesOnZero *int) int {
	result := dialNumber

	for i := 1; i <= value; i++ {
		if add {
			result = result + 1;
			if result == 100 {
				result = 0
				*timesPastZero++
			}
		} else {
			result = result - 1;
			if result == -1 {
				result = 99
			}

			if result == 0 {
				*timesPastZero++
			}
		}
	}

	if(result == 0) {
		*timesOnZero++
	}

	return result
}
