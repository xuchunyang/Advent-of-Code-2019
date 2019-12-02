package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func Calc(s []int) []int {
	c := make([]int, len(s))
	if n := copy(c, s); n != len(s) {
		panic(fmt.Sprintf("copy(%v, %v) = %d, want %d", c, s, n, len(s)))
	}

	for i := 0; ; {
		switch c[i] {
		case 1:
			ix, iy, iz := c[i+1], c[i+2], c[i+3]
			c[iz] = c[ix] + c[iy]
			i += 4
		case 2:
			ix, iy, iz := c[i+1], c[i+2], c[i+3]
			c[iz] = c[ix] * c[iy]
			i += 4
		case 99:
			return c
		}
	}
}

func Input() []int {
	dat, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var input []int
	for _, s := range strings.Split(strings.TrimSpace(string(dat)), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}
	return input
}

// part 1
func main1() {
	input := Input()
	input[1] = 12
	input[2] = 2
	fmt.Println(Calc(input))
}

func main() {
	input := Input()
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			input[1] = noun
			input[2] = verb
			if Calc(input)[0] == 19690720 {
				fmt.Println(noun, verb)
				return
			}
		}
	}
}
