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
		
	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(grandTotal)))
}
