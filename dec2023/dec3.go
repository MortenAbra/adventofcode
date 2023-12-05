package dec2023

import (
	"aoc/types"
	"aoc/utils"
	"log"
	"regexp"
	"strings"
	"unicode"
)

type Dec3Impl struct {
	Data          []string
	Symbols       string
	NumberPattern regexp.Regexp
	Results       types.Results
}

func GetDec3Instance(res types.Results) Dec {
	util := utils.GetUtil()
	return &Dec3Impl{
		Data:          util.ScannerLoadValues("dec2023/data/dec3-example.txt"),
		Symbols:       "*#+$",
		NumberPattern: *regexp.MustCompile(`(\d+)`),
		Results:       res,
	}
}

func (d *Dec3Impl) Calculate() types.Results {

	d.Results.ResultMap["2023-02/1"] += d.partOne("")
	//d.Results.ResultMap["2023-02/2"] += d.partTwo(s)

	return d.Results
}

func (d *Dec3Impl) partOne(in string) int {
	//var lines [][]int

	test := strings.Join(d.Data, "\n")
	log.Print(test)
	nums := d.NumberPattern.FindAllStringIndex(test, -1)
	log.Print(nums)

	for j, val := range d.Data {
		for i := 0; i < len(val); i++ {
			if ok := unicode.IsDigit(rune(val[i])); ok {
				log.Printf("Index: %d, Val: %v, OK: %v, J: %d", i, val, ok, j)

				// Find symbols only below first line
				if j == 0 {
					exist := strings.ContainsAny(d.Data[j+1], d.Symbols)
					if exist {
						log.Printf("Exist: %v - %v", exist, d.Data[j+1])
					}
					// Find symbols above last line
				} else if j == len(d.Data)-1 {
					exist := strings.ContainsAny(d.Data[j-1], d.Symbols)
					if exist {
						log.Printf("Exist: %v - %v", exist, d.Data[j-1])
					}
				}

			}
		}

	}
	/*
		for _, s := range d.Data {
			lines = d.NumberPattern.FindAllStringIndex(s, -1)
			log.Print(lines)
		}
	*/

	return 0
}

func (d *Dec3Impl) partTwo(in string) int {
	return 0
}
