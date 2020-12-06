package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func day6s2() {
	file, _ := ioutil.ReadFile("./input6")
	groups := strings.Split(string(file), "\n\n")
	var groupsAns []int
	for _, group := range groups {
		ans := []string{}
		persons := strings.Split(group, "\n")
		for i, person := range persons {
			var ansPerson []string
			answers := strings.Split(person, "")
			ansPerson = append(ansPerson, answers...)

			if i == 0 {
				ans = ansPerson
				continue
			}

			ans = intersect(ans, ansPerson)
		}
		groupsAns = append(groupsAns, len(ans))
	}
	fmt.Println(sum(groupsAns))
}

func intersect(as, bs []string) []string {
	i := make([]string, 0, int(math.Max(float64(len(as)), float64(len(bs)))))
	for _, a := range as {
		for _, b := range bs {
			if a == b {
				i = append(i, a)
			}
		}
	}
	return i
}
