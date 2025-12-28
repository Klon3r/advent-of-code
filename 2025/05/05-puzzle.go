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
	fmt.Println("05 - Advent Of Code")
	color.Unset()

	fileName := "05-input.txt"
	fileArray := utils.ReadFile(fileName)

	var freshRangeArray, idArray, freshIdArray []string

	beforeSpace := true

	// Split the array into fresh range/ids
	for _, value := range fileArray {
		if(value == "") {
			beforeSpace = false
		}

		if(beforeSpace) {
			freshRangeArray = append(freshRangeArray, value)
		} else {
			idArray = append(idArray, value)
		}
	}

	// Check the ids fall in fresh ranges
	for _, id := range idArray {
		idInt, _ := strconv.Atoi(id)
		idAdded := false

		for _, freshRange := range freshRangeArray {
			freshSplit := strings.Split(freshRange, "-")

			firstValue, _ := strconv.Atoi(freshSplit[0])
			secondValue, _ := strconv.Atoi(freshSplit[1])

			if(idInt >= firstValue && idInt <= secondValue && idAdded == false) {
				freshIdArray = append(freshIdArray, id)
				idAdded = true
			}

		}
	}

	//----------------------------
	// PART TWO
	//----------------------------
	var ranges [][]int
	beforeSpace = true
	
	// Split the array into fresh ranges
	for _, line := range fileArray {
    if line == "" {
        beforeSpace = false
        continue
    }
    if !beforeSpace {
        break
    }
    
    parts := strings.Split(line, "-")
    start, _ := strconv.Atoi(parts[0])
    end, _ := strconv.Atoi(parts[1])
    ranges = append(ranges, []int{start, end})
}

for i := 0; i < len(ranges); i++ {
	for j := i + 1; j < len(ranges); j++ {
			if ranges[i][0] > ranges[j][0] {
					ranges[i], ranges[j] = ranges[j], ranges[i]
			}
	}
}

var merged [][]int
if (len(ranges) > 0) {
	current := []int{ranges[0][0], ranges[0][1]}
	for i := 1; i < len(ranges); i++ {
		if(ranges[i][0] <= current[1]) {
			if (ranges[i][1] > current[1]) {
				current[1] = ranges[i][1]
			}
		}  else {
			merged = append(merged, current)
			current = []int{ranges[i][0], ranges[i][1]}
		}
	}
	merged = append(merged, current)
}

total := 0
for _, r := range merged {
	total += r[1] - r[0] + 1
}


	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(len(freshIdArray))))
	fmt.Println("Answer (Part Two):", color.GreenString(strconv.Itoa(total)))
}
