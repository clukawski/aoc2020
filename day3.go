package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func day3s1() {
	file, _ := ioutil.ReadFile("./input3")
	input := strings.Split(string(file), "\n")
	var x, y, tree int
	var lines [][]string

	for _, v := range input {
		lines = append(lines, strings.Split(v, ""))
	}

	for y < len(lines)-1 {
		x = (x + 3) % len(lines[y])
		y++

		if lines[y][x] == "#" {
			tree++
		}
	}

	fmt.Println(tree)
}

func day3s2() {
	file, _ := ioutil.ReadFile("./input3")
	input := strings.Split(string(file), "\n")
	var mult int
	var lines [][]string
	offsets := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, v := range input {
		lines = append(lines, strings.Split(v, ""))
	}

	for _, off := range offsets {
		var x, y, tree int
		for y < len(lines)-1 {
			x = (x + off[0]) % len(lines[y])
			y = y + off[1]

			if lines[y][x] == "#" {
				tree++
			}
		}
		if mult == 0 {
			mult = tree
		} else {
			mult = mult * tree
		}
	}
	fmt.Println(mult)
}
