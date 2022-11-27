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

func main() {

    dat, err := os.ReadFile("input.txt")
    check(err)
    whole := string(dat)
	lines := strings.Split(whole, "\n")

	inc := 0
	cur,_ := strconv.Atoi(lines[0])
	for _, line := range lines{
		val,_ := strconv.Atoi(line)
		if cur < val{
			inc++;
		}
		cur = val
	}
	fmt.Println(inc)

}