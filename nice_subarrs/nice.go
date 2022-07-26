package main

import (
	"fmt"
)

const debug = false

func main() {
	in := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	fmt.Println(getNiceSubArraysCount(in, 2))
}

func getNiceSubArraysCount(nums []int, k int) int {
	res := 0

	oddCount := 0
	niceSubArrsCount := 0

	startPos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			oddCount++
			if debug {
				fmt.Printf("кол-во нечетных (oddCount)++ : %v\n", oddCount)
			}
			niceSubArrsCount = 0 // нашли нечетное число - сбрасываем счетчик
		}

		for oddCount == k {
			niceSubArrsCount++
			if nums[startPos]%2 != 0 {
				oddCount--
			}
			startPos++ // двигаем окно слева
		}
		if niceSubArrsCount != 0 {
			res += niceSubArrsCount
		}
	}
	return res
}
