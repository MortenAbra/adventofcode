package main

import (
	"aoc/dec2023"
	"aoc/types"
)

func main() {
	res := types.Results{ResultMap: make(map[string]int)}

	puzzles := []dec2023.Dec{
		dec2023.GetDec1Instance(res),
		dec2023.GetDec2Instance(res),
		dec2023.GetDec3Instance(res),
	}

	for _, day := range puzzles {
		day.Calculate()
	}

	//log.Print(res)
}
