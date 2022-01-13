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

func main() {
	xs := []int{1, 2, 3, 4}
	fmt.Printf("%v\n", functionals.Sum(xs))
	fmt.Printf("%v\n", functionals.Product(xs))
	fmt.Printf("%v\n", binToDec("1011"))
	fmt.Printf("%v\n", functionals.Reverse(xs))
	fmt.Printf("%v\n", functionals.All([]int{2, 4, 6, 8}, func(x int) bool { return x%2 == 0 }))
	fmt.Printf("%v\n", functionals.Any([]int{2, 4, 5, 6, 8}, func(x int) bool { return x%2 == 1 }))
}
