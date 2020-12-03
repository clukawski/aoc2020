package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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
