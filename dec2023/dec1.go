package dec2023

import (
	"aoc/types"
	"aoc/utils"
	"strconv"
	"strings"
)

type Dec1Impl struct {
	Data            []string
	IntValues       string
	IntStringValues map[string]string
	Results         types.Results
}

func GetDec1Instance(res types.Results) Dec {
	util := utils.GetUtil()
	return &Dec1Impl{
		Data:      util.LoadValues("dec2023/data/dec1-1.txt"),
		IntValues: "1234567890",
		IntStringValues: map[string]string{
			"one":   "o1one",
			"two":   "t2two",
			"three": "t3three",
			"four":  "f4four",
			"five":  "f5five",
			"six":   "s6six",
			"seven": "s7seven",
			"eight": "e8eight",
			"nine":  "n9nine",
			"zero":  "z0zero",
		},
		Results: res,
	}
}

func (d *Dec1Impl) Calculate() types.Results {
	for _, s := range d.Data {
		d.Results.ResultMap["2023-01/1"] += d.partOne(s)
		d.Results.ResultMap["2023-01/2"] += d.partTwo(s)
	}

	return d.Results
}

func (d *Dec1Impl) partOne(in string) int {
	var firstNum, lastNum string
	return findIntInString(in, d.IntValues, firstNum, lastNum)
}

func (d *Dec1Impl) partTwo(in string) int {
	var firstNum, lastNum string
	data := in

	for key, r := range d.IntStringValues {
		data = strings.ReplaceAll(data, key, r)
	}

	return findIntInString(data, d.IntValues, firstNum, lastNum)
}

func findIntInString(s string, containsValues string, firstNum, lastNum string) int {
	for _, r := range s {
		if found := strings.Contains(containsValues, string(r)); found {
			if len(firstNum) == 0 {
				firstNum = string(r)
			} else {
				lastNum = string(r)
			}

		}
	}
	if len(lastNum) == 0 {
		lastNum = firstNum
	}

	res, _ := strconv.Atoi(firstNum + lastNum)
	return res
}
