package dec2023

import (
	"aoc/types"
	"aoc/utils"
	"regexp"
	"strconv"
	"strings"
)

type Dec2Impl struct {
	Data     []string
	GameReg  *regexp.Regexp
	ColorReg *regexp.Regexp
	Games    types.Games
	Results  types.Results
}

func GetDec2Instance(res types.Results) Dec {
	util := utils.GetUtil()
	return &Dec2Impl{
		Data:     util.LoadValues("dec2023/data/dec2.txt"),
		GameReg:  regexp.MustCompile(`^Game (\d+)`),
		ColorReg: regexp.MustCompile(`(\d+) (\w+)`),
		Games:    types.Games{Game: make(map[string]types.Cubes)},
		Results:  res,
	}
}

func (d *Dec2Impl) Calculate() types.Results {
	for _, s := range d.Data {
		d.Results.ResultMap["2023-02/1"] += d.partOne(s)
		d.Results.ResultMap["2023-02/2"] += d.partTwo(s)
	}

	return d.Results
}

func (d *Dec2Impl) partOne(in string) int {
	if in == "" {
		return 0
	}

	game := strings.Split(in, ":")
	gameId := strings.Split(game[0], " ")
	id, _ := strconv.Atoi(gameId[1])

	reds := 0
	greens := 0
	blues := 0

	for _, rounds := range strings.Split(game[1], ";") {
		for _, cubes := range strings.Split(rounds, ",") {
			m := d.ColorReg.FindStringSubmatch(cubes)

			switch m[2] {
			case "red":
				red, _ := strconv.Atoi(m[1])
				if red > reds {
					reds = red
				}
				break
			case "green":
				green, _ := strconv.Atoi(m[1])
				if green > greens {
					greens = green
				}
				break
			case "blue":
				blue, _ := strconv.Atoi(m[1])
				if blue > blues {
					blues = blue
				}
				break
			}

		}
	}

	//log.Printf("Game: %v\nRed: %v\nGreen: %v\nBlue: %v", id, reds, greens, blues)
	if reds <= 12 && greens <= 13 && blues <= 14 {

		return id
	}
	return 0
}

func (d *Dec2Impl) partTwo(in string) int {
	if in == "" {
		return 0
	}

	game := strings.Split(in, ":")

	reds := 0
	greens := 0
	blues := 0

	for _, rounds := range strings.Split(game[1], ";") {
		for _, cubes := range strings.Split(rounds, ",") {
			m := d.ColorReg.FindStringSubmatch(cubes)

			switch m[2] {
			case "red":
				red, _ := strconv.Atoi(m[1])
				if red > reds {
					reds = red
				}
				break
			case "green":
				green, _ := strconv.Atoi(m[1])
				if green > greens {
					greens = green
				}
				break
			case "blue":
				blue, _ := strconv.Atoi(m[1])
				if blue > blues {
					blues = blue
				}
				break
			}

		}
	}

	power := reds * greens * blues
	return power

}
