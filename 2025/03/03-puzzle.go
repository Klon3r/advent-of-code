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
	fmt.Println("03 - Advent Of Code")
	color.Unset()

	fileName := "03-input.txt"
	fileArray := utils.ReadFile(fileName)

	grandTotal := 0

	// Part One
	for _, value := range fileArray {
		valueArray := strings.Split(value, "")

		bigNum := 0

		for i := 0; i < len(valueArray) - 1; i++ {
			numOne := valueArray[i]
			numTwo := valueArray[i + 1]
			total := numOne + numTwo
			totalNum, _ := strconv.Atoi(total)		
			
			if totalNum > bigNum {
				bigNum = totalNum
			}

			// Loop through each of the remaining digits
			for x := i + 1; x < len(valueArray); x++ {
				numOne := valueArray[i]
				numTwo := valueArray[x]
				total := numOne + numTwo
				totalNum, _ := strconv.Atoi(total)		
			
				if totalNum > bigNum {
					bigNum = totalNum
				}
			}
		}
		grandTotal += bigNum
	}

	// Part Two (Greedy Stack Algorithm)
	totalResult := 0

	for _, value := range fileArray {
		stack := []rune{}
		toRemove := len(value) - 12

		// Process each digit
		for _, digit := range value {
			// While we can remove & stack has elements & top is smaller than current
			for toRemove > 0 && len(stack) > 0 && stack[len(stack) - 1] < digit {
				stack = stack[:len(stack) - 1] // Pop from stack
				toRemove--
			}

			stack = append(stack, digit)
		}

		// Take first 12 digits (should be exactly 12 after removals)
		result := string(stack[:12])
		resultNum, _ := strconv.Atoi(result)

		totalResult += resultNum
	}
		
	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(grandTotal)))
	fmt.Println("Answer (Part Two):", color.GreenString(strconv.Itoa(totalResult)))
}