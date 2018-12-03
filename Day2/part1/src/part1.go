package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var sum, sum2, sum3 int
	puzz, err := ioutil.ReadFile("puzzle.txt")
	if err == nil {
		for _, line := range strings.Split(string(puzz), "\n") {
			two := false
			three := false
			characters := make(map[string]int)
			for _, letter := range strings.Split(line, "") {
				if _, ok := characters[letter]; ok == false {
					characters[letter] = 1
				} else {
					characters[letter]++
				}
			}
			for i, _ := range characters {
				if two && three {
					break
				}
				if characters[i] == 2 && two == false {
					sum2++
					two = true
				} else if characters[i] == 3 && three == false {
					sum3++
					three = true
				}
			}
		}
		sum = sum2 * sum3
		fmt.Println(sum)
	}
}
