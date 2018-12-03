package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	puzz, err := ioutil.ReadFile("puzzle.txt")
	if err == nil {
		lines := strings.Split(string(puzz), "\n")
	loop:
		for i := 0; i < len(lines); i++ {
			for j := i + 1; j < len(lines)-1; j++ {
				answer := diffCheck(lines[i], lines[j])
				if answer != "" {
					fmt.Println(answer)
					break loop
				}
			}
		}
	}
}

func diffCheck(line1 string, line2 string) string {
	diff := 0
	j := 0
	var answer [25]byte
	for i := 0; i < len(line1); i++ {
		if line1[i] == line2[i] {
			answer[j] = line1[i]
			j++
		} else {
			diff++
		}
	}
	if diff == 1 {
		return string(answer[:])
	} else {
		return ""
	}
}
