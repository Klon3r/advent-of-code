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
	fmt.Println("08 - Advent Of Code")
	color.Unset()

	fileName := "08-input.txt"
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

	// fmt.Println(resultsArray)

	// Loop through the results array and calculate the connections
	var connections [][]string
	
	for i, row := range resultsArray {
		// Only calculate the 10 closest connections
		if(i >= 1000) {
			break
		}

		p1 := row.pointOne
		p2 := row.pointTwo

		// Find which connections each point is in (-1 if not found)
		p1Connection := -1
		p2Connection := -1
		for i, connection := range connections {
			for _, point := range connection {
				if point == p1 {
					p1Connection = i
				}
				if point == p2 {
					p2Connection = i
				}
			}
		}

		if p1Connection != -1 && p1Connection == p2Connection {
			continue // Skip this connection, don't count it
		}

		if p1Connection == -1 && p2Connection == -1 {
			connections = append(connections, []string{p1,p2})
		} else if p1Connection == -1 {
			// p1 is new, add the p2's connection
			connections[p2Connection] = append(connections[p2Connection], p1)
		} else if p2Connection == -1 {
			// p1 is new, add the p2's connection
			connections[p1Connection] = append(connections[p1Connection], p2)
		} else if p1Connection != p2Connection {
			// Both exist in different circuits - merge them
			connections[p1Connection] = append(connections[p1Connection], connections[p2Connection]...)
			// Remove the now-empty p2Connection circuit
			connections = append(connections[:p2Connection], connections[p2Connection+1:]...)
	}

	}

	fmt.Println("Connections:", connections)
	fmt.Println("------------------------------------------------------------")

	// Get circuit sizes
	var sizes []int
	for _, circuit := range connections {
			sizes = append(sizes, len(circuit))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	fmt.Println("Circuit sizes:", sizes)
	fmt.Println("Answer:", sizes[0]*sizes[1]*sizes[2])
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