package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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
