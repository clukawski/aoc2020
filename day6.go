package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func day6s1() {
	file, _ := ioutil.ReadFile("./input6")
	groups := strings.Split(string(file), "\n\n")
	var groupsAns []int
	for _, group := range groups {
		total := 0
		ansMap := make(map[string]bool)
		persons := strings.Split(group, "\n")
		for _, person := range persons {
			answers := strings.Split(person, "")
			for _, answer := range answers {
				if _, ok := ansMap[answer]; ok {
					continue
				}
				ansMap[answer] = true
				total++
			}
		}
		groupsAns = append(groupsAns, total)
	}
	fmt.Println(sum(groupsAns))
}
