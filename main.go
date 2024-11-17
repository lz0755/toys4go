package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{16, 17, 18}

	fmt.Println(numFriendRequests(ages))
}

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	res := 0
	for x, y := 1, 0; x < len(ages); x++ {
		for ages[y] <= (ages[x]/2+7) && y < x {
			y++
		}
		res += x - y
	}

	return res
}
