package main

import (
	"fmt"
	"slices"
)

func plusOne(digits []int) (r []int) {
	o := 0

	for _, d := range digits {
		o *= 10
		o += d
	}

	o += 1

	r = make([]int, 0)

	for {
		nd := o % 10
		r = append(r, nd)

		o -= nd
		o/=10

		if o == 0 {
			break
		}

	}


	slices.Reverse(r)
    return 
}

func main() {
	fmt.Println(plusOne([]int{1,2,3}))
}