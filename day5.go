package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const rows = 128

func day5s1() {
	file, _ := ioutil.ReadFile("./input5")
	input := strings.Split(string(file), "\n")

	var rows []int
	var seatIDs []int
	for i := 0; i < 128; i++ {
		rows = append(rows, i+1)
	}

	var columns []int
	for i := 0; i < 8; i++ {
		columns = append(columns, i+1)
	}

	for _, pass := range input {
		rowrange := rows
		colrange := columns
		parts := strings.Split(pass, "")[:7]

		for _, part := range parts {
			if len(rowrange) == 1 {
				break
			}
			half := len(rowrange) / 2
			switch part {
			case "F":
				rowrange = rowrange[:half]
			case "B":
				rowrange = rowrange[half:]
			}
			fmt.Printf("rowrange (%s): %v\n", part, rowrange)
		}

		parts2 := strings.Split(pass, "")[7:]
		for _, part := range parts2 {
			if len(colrange) == 1 {
				break
			}
			half := len(colrange) / 2

			switch part {
			case "L":
				colrange = colrange[:half]
			case "R":
				colrange = colrange[half:]
			}
			fmt.Printf("colrange (%s): %v\n", part, colrange)
		}

		seatIDs = append(seatIDs, (rowrange[0]-1)*8+(colrange[0]-1))
	}

	sort.Ints(seatIDs)

	fmt.Printf("list: %v\n", seatIDs)
	fmt.Printf("highest: %d\n", seatIDs[len(seatIDs)-1])
}

func day5s2() {
	file, _ := ioutil.ReadFile("./input5")
	input := strings.Split(string(file), "\n")

	var rows []int
	var seatIDs []int
	var seatMap = make(map[struct {
		Row    int
		Column int
	}]int)
	var seatIDMap = make(map[int][]int)
	for i := 0; i < 128; i++ {
		rows = append(rows, i+1)
	}

	var columns []int
	for i := 0; i < 8; i++ {
		columns = append(columns, i+1)
	}

	for _, pass := range input {
		rowrange := rows
		colrange := columns
		parts := strings.Split(pass, "")[:7]

		for _, part := range parts {
			if len(rowrange) == 1 {
				break
			}
			half := len(rowrange) / 2
			switch part {
			case "F":
				rowrange = rowrange[:half]
			case "B":
				rowrange = rowrange[half:]
			}
		}

		parts2 := strings.Split(pass, "")[7:]
		for _, part := range parts2 {
			if len(colrange) == 1 {
				break
			}
			half := len(colrange) / 2
			// half := (len(colrange) / 2) + (len(colrange) % 2)
			switch part {
			case "L":
				colrange = colrange[:half]
			case "R":
				colrange = colrange[half:]
			}
		}

		id := (rowrange[0]-1)*8 + (colrange[0] - 1)
		seatIDs = append(seatIDs, id)
		seatMap[struct {
			Row    int
			Column int
		}{rowrange[0] - 1, colrange[0] - 1}] = id
		seatIDMap[id] = []int{rowrange[0] - 1, colrange[0] - 1}
	}

	sort.Ints(seatIDs)

	var rownums []int
	for _, v := range seatIDMap {
		rownums = append(rownums, v[0])
	}
	sort.Ints(rownums)
	uniqueRows := unique(rownums)
	var mySeat int
	for _, row := range uniqueRows[1 : len(uniqueRows)-1] {
		for i := 0; i < 8; i++ {
			if _, ok := seatMap[struct {
				Row    int
				Column int
			}{Row: row, Column: i}]; !ok {
				mySeat = row*8 + i
			}
		}
	}

	fmt.Printf("my seat: %d\n", mySeat)
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
