package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.Set(color.FgHiMagenta, color.Bold)
	fmt.Println("___________________")
	fmt.Println("02 - Advent Of Code")
	color.Unset()

	// puzzleExample := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	// puzzleExampleSplit := strings.Split(puzzleExample, ",")
	puzzleInput := "655-1102,2949-4331,885300-1098691,1867-2844,20-43,4382100-4484893,781681037-781860439,647601-734894,2-16,180-238,195135887-195258082,47-64,4392-6414,6470-10044,345-600,5353503564-5353567532,124142-198665,1151882036-1151931750,6666551471-6666743820,207368-302426,5457772-5654349,72969293-73018196,71-109,46428150-46507525,15955-26536,65620-107801,1255-1813,427058-455196,333968-391876,482446-514820,45504-61820,36235767-36468253,23249929-23312800,5210718-5346163,648632326-648673051,116-173,752508-837824"
	puzzleInputSplit := strings.Split(puzzleInput, ",")

	totalPuzzlePartOne := 0;
	totalPuzzlePartTwo := 0;


	// Part 1
	for _, value := range puzzleInputSplit {
		puzzleRange := strings.Split(value, "-")
		firstId := puzzleRange[0]
		lastId := puzzleRange[1]

		// Convert from string to int
		firstNumber, _ := strconv.Atoi(firstId)
		lastNumber, _ := strconv.Atoi(lastId)

		for i := firstNumber; i <= lastNumber; i++ {
			iString := strconv.Itoa(i)
			length := len(iString)
			isEven := length % 2

			if (isEven == 0) {
				iArray := strings.Split(iString, "")
				half := length / 2;
				firstHalf := strings.Join(iArray[0:half], "")
				secondHalf := strings.Join(iArray[half:length], "")

				if(firstHalf == secondHalf) {
					total := firstHalf + secondHalf
					number, _ := strconv.Atoi(total)
					totalPuzzlePartOne += number
				}
			}
		}
	}

	// Part 2
	for _, value := range puzzleInputSplit {
		puzzleRange := strings.Split(value, "-")
		firstId := puzzleRange[0]
		lastId := puzzleRange[1]

		// Convert from string to int
		firstNumber, _ := strconv.Atoi(firstId)
		lastNumber, _ := strconv.Atoi(lastId)

		for i := firstNumber; i <= lastNumber; i++ {
			iString := strconv.Itoa(i)
			length := len(iString)
			isInvalid := false

			for x := 1; x <= length / 2; x++ {
				if length % x == 0 {
					subString := iString[0:x]
					repeated := strings.Repeat(subString, length / x)

					if(repeated == iString) {
						isInvalid = true
						break
					}
				}
			}

			if isInvalid {
				totalPuzzlePartTwo += i
			}
		}		
	}
	

	fmt.Println("Answer (Part One):", color.GreenString(strconv.Itoa(totalPuzzlePartOne)))
	fmt.Println("Answer (Part Two):", color.GreenString(strconv.Itoa(totalPuzzlePartTwo)))

}
