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
		v := strings.Split(line, "x")

		if len(v) == 1 {
			break
		}

		i := []int{toInt(v[0]), toInt(v[1]), toInt(v[2])}

		minValue := min(i[0]*i[1], i[1]*i[2], i[0]*i[2])
		area := i[0]*i[1]*2 + i[1]*i[2]*2 + i[0]*i[2]*2
		sumOfAllAreas += area + minValue
		fmt.Printf("%v -> %d + %d = %d\n", v, area, minValue, (area + minValue))
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
