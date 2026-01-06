package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func parseIntoColumns(file []string) [][]string{
	// Find length of row
	maxCols := 0
	for _, line := range file {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) > maxCols {
			maxCols = len(fields)
		}
	}

	var columnArray = make([][]string, maxCols) 

	// Place columns into array
	for _, line := range file {
		valueSplit := strings.Fields(line)
		for index, val := range valueSplit {
			columnArray[index] = append(columnArray[index], val)
		}
	}

	return columnArray
}

func parseIntoColumnsPartTwo(file []string) [][]string {
	maxLineLength := 0
	for _, line := range file {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}
	
	// Initialize column array - each column will be a slice of strings
	columnArray := make([][]string, maxLineLength)
	
	// Get the last row (operators row)
	lastRowIndex := len(file) - 1
	
	// Process each column position from left to right
	for col := 0; col < maxLineLength; col++ {
		// Check if there's an operator in the last row at this column
		var operator string
		if lastRowIndex >= 0 && col < len(file[lastRowIndex]) {
			char := string(file[lastRowIndex][col])
			if char == "+" || char == "*" {
				operator = char
				columnArray[col] = append(columnArray[col], operator)
			}
		}
		
		// Process data rows from top to bottom (excluding the last row)
		for row := 0; row < lastRowIndex; row++ {
			// Check if this column position exists in this row
			if col < len(file[row]) {
				char := string(file[row][col])
				// Only add non-space characters
				if char != " " {
					columnArray[col] = append(columnArray[col], char)
				}
			}
		}
	}
	
	// Filter out empty columns and return only non-empty ones
	var result [][]string
	for _, col := range columnArray {
		if len(col) > 0 {
			result = append(result, col)
		}
	}
	
	return result
}
	
func calc(operator string, count int, num int) int {
	switch operator {
	case "+":
		return count + num
	case "*":
		if num == 0 {
			return count
		} else {
		return count * num
		}
	}
	return count
}


func main() {
	color.Set(color.FgHiMagenta, color.Bold)
	fmt.Println("___________________")
	fmt.Println("06 - Advent Of Code")
	color.Unset()

	fileName := "06-input.txt"
	fileArray := utils.ReadFile(fileName)
	rowLength := len(fileArray)
	fileCols := parseIntoColumns(fileArray)

	total := 0

	for _, cols:= range fileCols {
		operator := cols[rowLength-1]
		count := 0

		for i := 0; i < len(cols)-1; i++ {
			num, _ := strconv.Atoi(cols[i])

			if i == 0 {
				count = num
			} else {
				count = calc(operator, count, num)
		}
	}

	total += count
	}
	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(total)))

	// Part Two
	total = 0
	file := utils.ReadFile(fileName)
	result := parseIntoColumnsPartTwo(file)
	var operator string
	count := 0
	// var countArray []int

	for _, array := range result {
		var numArray []string
		array = append(array, "-")
		for _, number := range array {
			if(number == "*" || number == "+") {
				operator = number
				total += count
				count = 0
			} else {
				if(number != "-") {
					numArray = append(numArray, number)
				}
				
			}
		}
		join := strings.Join(numArray,"")
		num, _ := strconv.Atoi(join)

		if (count == 0) {
			count = num
		} else {
			count = calc(operator, count, num)
		}
}
total += count

	fmt.Println("Answer (Part Two):", color.GreenString(strconv.Itoa(total)))
}
