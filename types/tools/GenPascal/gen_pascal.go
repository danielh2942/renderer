package main

// gen_pascal.go
// Version 1
// 2022 - Daniel Hannon

import (
	"os"
	"strconv"
)

// This generates Pascals triangle because I was too lazy to do it by hand

func main() {
	pascal := [][]uint64{{1}, {1, 1}}
	var row uint64 = 2
	for row < 100 {
		pascal = append(pascal, []uint64{1})
		var col int = 0
		for col < (len(pascal[row-1]) - 1) {
			pascal[row] = append(pascal[row], pascal[row-1][col]+pascal[row-1][col+1])
			col++
		}
		pascal[row] = append(pascal[row], 1)
		row++
	}

	output := ""
	for _, v := range pascal {
		output += "{"
		for idx, x := range v {
			output += strconv.FormatUint(x, 10)
			if idx != len(v)-1 {
				output += ","
			}
		}
		output += "},\n"
	}
	f, err := os.Create("pascal.txt")
	if err != nil {
		panic(-1)
	}
	defer f.Close()
	f.WriteString(output)
}
