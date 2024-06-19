package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Destination int = iota
	Source
	Length
)

type Range [3]uint64

type Data struct {
	Seeds                 []uint64
	SeedToSoil            []Range
	SoilToFertilizer      []Range
	FertilizerToWater     []Range
	WaterToLight          []Range
	LightToTemperature    []Range
	TemperatureToHumidity []Range
	HumidityToLocation    []Range
}

func parseRange(lines []string) []Range {
	ranges := make([]Range, len(lines))
	for i, line := range lines {
		tokens := strings.Split(strings.TrimSpace(line), " ")
		for j, token := range tokens {
			r, _ := strconv.ParseUint(token, 10, 64)
			ranges[i][j] = r
		}
	}
	return ranges
}

func parse(lines []string) Data {
	data := Data{}

	seeds := strings.Split(lines[0], " ")
	for _, seed := range seeds[1:] {
		seed := strings.TrimSpace(seed)
		if seed == "" {
			continue
		}
		seedUint, _ := strconv.ParseUint(seed, 10, 64)
		data.Seeds = append(data.Seeds, seedUint)
	}

	start, end := 3, 3
	for lines[end] != "" {
		end++
	}
	data.SeedToSoil = parseRange(lines[start:end])

	start = end + 2
	end += 1
	for lines[end] != "" {
		end++
	}
	data.SoilToFertilizer = parseRange(lines[start:end])

	start = end + 2
	end += 1
	for lines[end] != "" {
		end++
	}
	data.FertilizerToWater = parseRange(lines[start:end])

	start = end + 2
	end += 1
	for lines[end] != "" {
		end++
	}
	data.WaterToLight = parseRange(lines[start:end])

	start = end + 2
	end += 1
	for lines[end] != "" {
		end++
	}
	data.LightToTemperature = parseRange(lines[start:end])

	start = end + 2
	end += 1
	for lines[end] != "" {
		end++
	}
	data.TemperatureToHumidity = parseRange(lines[start:end])

	start = end + 2
	end += 1
	for end < len(lines) && lines[end] != "" {
		end++
	}
	data.HumidityToLocation = parseRange(lines[start:end])

	return data
}

func findInRanges(rnge []Range, value uint64) uint64 {
	for _, r := range rnge {
		if value >= r[Source] && value < r[Source]+r[Length] {
			return r[Destination] + (value - r[Source])
		}
	}
	return value
}

func findLocation(data Data, seed uint64) uint64 {
	soil := findInRanges(data.SeedToSoil, seed)
	fertilizer := findInRanges(data.SoilToFertilizer, soil)
	water := findInRanges(data.FertilizerToWater, fertilizer)
	light := findInRanges(data.WaterToLight, water)
	temperature := findInRanges(data.LightToTemperature, light)
	humidity := findInRanges(data.TemperatureToHumidity, temperature)
	location := findInRanges(data.HumidityToLocation, humidity)
	// fmt.Println(seed, soil, fertilizer, water, light, temperature, humidity, location)
	return location
}

func main() {
	lines := []string{}
	input, _ := os.Open("./5/input.txt")
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	minLocation := uint64(math.MaxUint64)

	data := parse(lines)
	for _, seed := range data.Seeds {
		if location := findLocation(data, seed); location < minLocation {
			minLocation = location
		}
	}
	fmt.Println("minimum location", minLocation)

	minLocation = uint64(math.MaxUint64)

	for i := 0; i < len(data.Seeds); i += 2 {
		for seed := data.Seeds[i]; seed < data.Seeds[i]+data.Seeds[i+1]; seed++ {
			if location := findLocation(data, seed); location < minLocation {
				minLocation = location
			}
		}
	}

	fmt.Println("minimum location for seeds range", minLocation)
}
