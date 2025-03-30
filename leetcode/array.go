package leetcode

import "fmt"

func TwoSum(nums []int, target int) []int {
	l := 0
	r := nums[len(nums)-2]
	res := make([]int, 2)
	for {
		if (nums[l] + nums[r]) == target {
			res[0] = l
			res[1] = r
			break
		}
		if (nums[l] + nums[r]) > target {
			r--
		}
		if (nums[l] + nums[r]) < target {
			l++
		}
		if l == r {
			break
		}
	}
	return res
}

func LeetcodeArrayPrint() {
	a := TwoSum([]int{1, 2, 3}, 5)
	fmt.Println(a)
}
