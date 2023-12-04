package dec2023

import (
	"aoc/types"
	"aoc/utils"
	"log"
	"regexp"
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
	//var part int
	for _, s := range d.Data {
		lines := d.NumberPattern.FindAllStringIndex(s, -1)
		//log.Print(lines)

		log.Printf("Line: %v", lines[1][1])

		break
	}

	return 0
}

func (d *Dec3Impl) partTwo(in string) int {
	return 0
}
