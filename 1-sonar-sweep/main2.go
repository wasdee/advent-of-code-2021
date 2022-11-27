package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func zip[T any](lists ...[]T) func() []T {
	zip := make([]T, len(lists))
	i := 0
	return func() []T {
		for j := range lists {
			if i >= len(lists[j]) {
				return nil
			}
			zip[j] = lists[j][i]
		}
		i++
		return zip
	}
}

func main() {

	dat, err := os.ReadFile("input.txt")
	check(err)
	whole := string(dat)
	lines := strings.Split(whole, "\n")
	lines_int := make([]int, len(lines))
	for i, l := range lines {
		val, _ := strconv.Atoi(l)
		lines_int[i] = val
	}

	inc := 0
	var cur int
	just_start := true

	iter := zip(lines_int, lines_int[1:], lines_int[2:])
	for tuple := iter(); tuple != nil; tuple = iter() {
		sum := 0
		for _, v := range tuple {
			sum += v
		}
		if just_start {
			cur = sum
			just_start = false
		} else {
			if cur < sum {
				inc++
			}
			cur = sum
		}
	}
	// for it{
	// 	val,_ := strconv.Atoi(line)
	// 	if cur < val{
	// 		inc++;
	// 	}
	// 	cur = val
	// }
	fmt.Println(inc)

}
