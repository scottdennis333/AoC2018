package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	total := 0
	puzz, err := ioutil.ReadFile("puzzle.txt")
	if err == nil {
		for _, line := range strings.Split(string(puzz), "\n") {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("error")
			}
			total += num
		}
	}
	fmt.Println(total)
}
