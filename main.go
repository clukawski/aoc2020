package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
}

func day1s1() {
	file, _ := ioutil.ReadFile("./input")
	input := strings.Split(string(file), "\n")
	for _, i := range input {
		for _, j := range input {
			is, _ := strconv.Atoi(i)
			js, _ := strconv.Atoi(j)
			if (is + js) == 2020 {
				fmt.Printf("is*js = %d\n", is*js)
				return
			}
		}
	}
}

func day1s2() {
	file, _ := ioutil.ReadFile("./input")
	input := strings.Split(string(file), "\n")
	for _, i := range input {
		for _, j := range input {
			for _, k := range input {
				is, _ := strconv.Atoi(i)
				js, _ := strconv.Atoi(j)
				ks, _ := strconv.Atoi(k)
				if (is + js + ks) == 2020 {
					fmt.Printf("is*js*ks = %d\n", is*js*ks)
					return
				}
			}
		}
	}

}

func day2s1() {
	file, _ := ioutil.ReadFile("./input2")
	input := strings.Split(string(file), "\n")
	var valid int
	for _, i := range input {
		entry := strings.Split(i, ": ")
		rule := strings.Split(entry[0], " ")

		password := entry[1]
		letter := rule[1]
		index := strings.Split(rule[0], "-")
		index1, _ := strconv.Atoi(index[0])
		index2, _ := strconv.Atoi(index[1])

		count := strings.Count(password, letter)
		if count >= index1 && count <= index2 {
			valid++
		}
	}

	fmt.Println(valid)
}

func day2s2() {
	file, _ := ioutil.ReadFile("./input2")
	input := strings.Split(string(file), "\n")
	var valid int
	for _, i := range input {
		entry := strings.Split(i, ": ")
		rule := strings.Split(entry[0], " ")

		password := entry[1]
		letter := rule[1]
		index := strings.Split(rule[0], "-")
		index1, _ := strconv.Atoi(index[0])
		index2, _ := strconv.Atoi(index[1])

		pwArray := strings.Split(password, "")
		if (pwArray[index1-1] == letter) != (pwArray[index2-1] == letter) {
			valid++
		}
	}

	fmt.Println(valid)
}

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
