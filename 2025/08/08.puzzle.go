package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)


func main() {
	color.Set(color.FgHiMagenta, color.Bold)
	fmt.Println("___________________")
	fmt.Println("07 - Advent Of Code")
	color.Unset()

	fileName := "08-example.txt"
	partOne(fileName)
}

type junction struct {
	result float64
	pointOne string
	pointTwo string	
}

func partOne(inputFile string) {
	fileArray := utils.ReadFile(inputFile)

	// Find the shortest connections (unique)
	var resultsArray []junction

	for i := 0; i < len(fileArray); i++ {
		for j := i+1; j < len(fileArray); j++ {
			dist := euclideanDistance(fileArray[i], fileArray[j])
			resultsArray = append(resultsArray, junction{
				result:   dist,
				pointOne: fileArray[i],
				pointTwo: fileArray[j],
			})
		}
	}
	
	// Sort the resultsArray from shortest to longest via result
	sort.Slice(resultsArray, func(i, j int) bool {
    return resultsArray[i].result < resultsArray[j].result
})

	fmt.Println(resultsArray)
}

// In mathematics, the Euclidean distance between two points in a Euclidean space
// is the length of the line segment between them [https://en.wikipedia.org/wiki/Euclidean_distance]
func euclideanDistance(pointOne string, pointTwo string) float64{
	// Break string into x,y,z
	p1Split := strings.Split(pointOne, ",")
	p2Split := strings.Split(pointTwo, ",")

	// Convert p1 to x,y,z
	p1X, _ := strconv.Atoi(p1Split[0])
	p1Y, _ := strconv.Atoi(p1Split[1])
	p1Z, _ := strconv.Atoi(p1Split[2])

	// Convert p2 to x,y,z
	p2X, _ := strconv.Atoi(p2Split[0])
	p2Y, _ := strconv.Atoi(p2Split[1])
	p2Z, _ := strconv.Atoi(p2Split[2])



	//d=(x2​−x1​)2+(y2​−y1​)2+(z2​−z1​)2
	equationPartOne := ((p2X - p1X) * (p2X - p1X)) + ((p2Y - p1Y) * (p2Y - p1Y)) + ((p2Z - p1Z) * (p2Z - p1Z))
	result := math.Sqrt(float64(equationPartOne))

	return result
}