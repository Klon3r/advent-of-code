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


func main() {
	color.Set(color.FgHiMagenta, color.Bold)
	fmt.Println("___________________")
	fmt.Println("05 - Advent Of Code")
	color.Unset()

	fileName := "06-example.txt"
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
				switch operator {
				case "+":
					count = count + num
				case "*":
					count = count * num
			}

		}
	}

	total += count
	}
	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(total)))

}
