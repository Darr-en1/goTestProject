package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func twoSum(nums []int, target int) []int {
	var temp = make(map[int]int)
	for idx, val := range nums {
		if v, ok := temp[val]; ok == true {
			return []int{v, idx}
		} else {
			temp[target-val] = idx
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 8))

	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n ", i, a(i))
	}
}
