package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Calc(mass int) (fuel int) {
	return mass/3 - 2
}

func Calc2(mass int) (totalFuel int) {
	for {
		fuel := Calc(mass)
		if fuel <= 0 {
			break
		}
		totalFuel += fuel
		mass = fuel
	}
	return
}

func main() {
	var total int
	var total2 int
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		mass, err := strconv.Atoi(input.Text())
		if err != nil {
			log.Fatal(err)
		}
		total += Calc(mass)
		total2 += Calc2(mass)
	}
	fmt.Println("Total fuel:", total)
	fmt.Println("Total fuel 2:", total2)
}
