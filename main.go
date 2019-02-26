package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during reading file: %v\n", err)
	}
	var sumOfAllAreas int
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			break
		}
		vString := strings.Split(line, "x")

		d := []int{toInt(vString[0]), toInt(vString[1]), toInt(vString[2])}

		area := d[0]*d[1]*2 + d[1]*d[2]*2 + d[0]*d[2]*2
		smallestSideArea := min(d[0]*d[1], d[1]*d[2], d[0]*d[2])
		fmt.Printf("%v -> %d + %d = %d\n", d, area, smallestSideArea, (area + smallestSideArea))

		sumOfAllAreas += area + smallestSideArea
	}

	fmt.Printf("You need %d dm^2 of paper\n", sumOfAllAreas)
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during reading file: %v\n", err)
	}
	return v
}

func min(args ...int) int {
	minValue := args[0]
	for _, v := range args {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}
