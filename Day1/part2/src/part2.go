package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	total := 0
	totals := map[int]bool{
		0: true,
	}
loop:
	for {
		puzz, err := ioutil.ReadFile("puzzle.txt")
		if err == nil {
			for _, line := range strings.Split(string(puzz), "\n") {
				num, err := strconv.Atoi(line)
				if err != nil {
					fmt.Println(err)
				}
				total += num
				if _, ok := totals[total]; ok == false {
					totals[total] = true
				} else {
					fmt.Println(total)
					break loop
				}
			}
		} else {
			fmt.Println(err)
		}
	}
}
