package main

import (
	"fmt"
	"jaco/functionals"
)

func binToDec(bin string) int {
	return functionals.FoldLeft([]byte(bin), 0, func(acc int, x byte) int {
		return 2*acc + int(x-'0')
	})
}

func reverse(xs []int) []int {
	return functionals.FoldLeft(xs, []int{}, func(acc []int, x int) []int {
		return append([]int{x}, acc...)
	})
}

func main() {
	xs := []int{1, 2, 3, 4}
	fmt.Printf("%v\n", functionals.Sum(xs))
	fmt.Printf("%v\n", functionals.Product(xs))
	fmt.Printf("%v\n", binToDec("1011"))
	fmt.Printf("%v\n", reverse(xs))
}
