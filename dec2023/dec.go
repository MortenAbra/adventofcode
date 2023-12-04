package dec2023

import "aoc/types"

type Dec interface {
	Calculate() types.Results
	partOne(in string) int
	partTwo(in string) int
}
