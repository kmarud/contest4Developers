// main.go
package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	"os"

	// "math"
	"strconv"
	"strings"
)

func main() {

	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "print: %v\n", err)
	}
	var sumAll int
	for _, line := range strings.Split(string(data), "\n") {

		v := strings.Split(line, "x")
		// fmt.Println(len(v))
		if len(v) == 1 {
			break
		}
		var i [3]int
		i[0], _ = strconv.Atoi(v[0])
		i[1], _ = strconv.Atoi(v[1])
		i[2], _ = strconv.Atoi(v[2])
		// math.Min(1, 2, 3)
		mins := minim(i[0]*i[1], i[1]*i[2], i[0]*i[2])
		// fmt.Println("mins", mins)
		sum := i[0]*i[1]*2 + i[1]*i[2]*2 + i[0]*i[2]*2
		sumAll += sum + mins
		fmt.Println(v, sum, "+", mins, "=", (sum + mins))
	}

	fmt.Println("sum: ", sumAll)
}

func minim(args ...int) int {
	// var mint int
	mint := args[0]
	for _, v := range args {
		if v < mint {
			mint = v
		}
	}
	return mint
}

// func m(a, b int) int {
// 	if a <= b {
// 		return a
// 	}
// 	return b
// }
