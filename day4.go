package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var reqFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

func day4s1() {
	file, _ := ioutil.ReadFile("./input4")
	rawInput := strings.Split(string(file), "\n\n")
	var validPass int
	var lines []map[string]string
	for _, line := range rawInput {
		keyMap := make(map[string]string)
		values := strings.Split(strings.ReplaceAll(line, "\n", " "), " ")
		for _, value := range values {
			kv := strings.Split(value, ":")
			keyMap[kv[0]] = kv[1]
		}
		lines = append(lines, keyMap)
	}

	for _, fields := range lines {
		_, ok := fields["cid"]
		var valid int
		for _, req := range reqFields {
			if _, ok := fields[req]; ok {
				valid++
			}
		}

		if !ok && valid == (len(reqFields)-1) {
			valid++
		}
		if valid == len(reqFields) {
			validPass++
		}
	}

	fmt.Println(validPass)
}

func day4s2() {
	hclRegexp := regexp.MustCompile(`^[#]{1}[0-9a-f]{6}$`)
	pidRegexp := regexp.MustCompile("^[0-9]{9}$")

	file, _ := ioutil.ReadFile("./input4")
	rawInput := strings.Split(string(file), "\n\n")
	var validPass int
	var lines []map[string]string
	for _, line := range rawInput {
		keyMap := make(map[string]string)
		values := strings.Split(strings.ReplaceAll(line, "\n", " "), " ")
		for _, value := range values {
			kv := strings.Split(value, ":")
			keyMap[kv[0]] = kv[1]
		}
		lines = append(lines, keyMap)
	}

	for _, fields := range lines {
		var valid int
		var thing []string
		for _, req := range reqFields {
			if val, ok := fields[req]; ok {
				switch req {
				case "byr":
					if !rangeValid(val, 1920, 2002) {
						continue
					}
				case "iyr":
					if !rangeValid(val, 2010, 2020) {
						continue
					}
				case "eyr":
					if !rangeValid(val, 2020, 2030) {
						continue
					}
				case "hgt":
					if h := strings.Split(val, "cm"); len(h) != 1 {
						if !rangeValid(h[0], 150, 193) {
							continue
						}
					} else if h := strings.Split(val, "in"); len(h) != 1 {
						if !rangeValid(h[0], 59, 76) {
							continue
						}
					} else {
						continue
					}
				case "hcl":
					if !hclRegexp.MatchString(val) {
						continue
					}
				case "ecl":
					things := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
					inThing := false
					for _, thing := range things {
						if val == thing {
							inThing = true
						}
					}
					if !inThing {
						continue
					}
				case "pid":
					if len(val) != 9 {
						continue
					}
					if !pidRegexp.MatchString(val) {
						continue
					}
				case "cid":
					continue
				}
				if req != "cid" {
					thing = append(thing, req)

				}
				valid++
			}
		}

		if valid == len(reqFields)-1 {
			validPass++
		}
	}

	fmt.Println(validPass)
}

func rangeValid(val string, min, max int) bool {
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if valInt < min || valInt > max {
		return false
	}

	return true
}
