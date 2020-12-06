package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var seatIDs = *getSeatIDList()

// Note: this is the my second implementation, go back in history to see previous, uglier solution, thanks to @benharri for inspo

func day5s1() {
	fmt.Printf("highest: %d\n", seatIDs[len(seatIDs)-1])
}

func day5s2() {
	seat := (len(seatIDs)+1)*(seatIDs[0]+seatIDs[len(seatIDs)-1])/2 - sum(seatIDs)
	fmt.Printf("my seat: %d\n", seat)
}

func getSeatIDList() *[]int {
	file, _ := ioutil.ReadFile("./input5")
	input := strings.Split(string(file), "\n")

	var seatIDs []int
	for _, pass := range input {
		id := getSeatID(pass)
		seatIDs = append(seatIDs, int(id))
	}
	sort.Ints(seatIDs)

	return &seatIDs
}

func getSeatID(input string) int64 {
	mapReplace := func(r rune) rune {
		switch r {
		case 'F', 'L':
			return '0'
		case 'B', 'R':
			return '1'
		}
		return 'x' // invalid
	}
	i2, _ := strconv.ParseInt(strings.Map(mapReplace, input), 2, 64)
	return i2
}

func unique(slice []int) []int {
	keys := make(map[int]bool)
	valid := []int{}
	for _, value := range slice {
		if _, ok := keys[value]; !ok {
			keys[value] = true
			valid = append(valid, value)
		}
	}

	return valid
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
