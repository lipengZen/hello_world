package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)

	if nums == nil || len(nums) < 3 {
		return ans
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// ans := make([][]int, 0)
	for i := 1; i < len(nums)-1; i++ {
		if i < len(nums)-2 && nums[i] == nums[i+1] && nums[i] == nums[i+2] {
			continue
		}
		l, r := 0, len(nums)-1
		for l < r {

			// 如何判断重复
			for l+1 < i && nums[l] == nums[l+1] {
				l++
			}
			for i < r-1 && nums[r] == nums[r-1] {
				r--
			}

			if nums[l] > 0 || nums[r] < 0 || l >= i || r <= i {
				break
			}

			sum := nums[l] + nums[i] + nums[r]
			if sum == 0 {
				slice := make([]int, 3, 3)
				slice[0] = nums[l]
				slice[1] = nums[i]
				slice[2] = nums[r]
				ans = append(ans, slice)
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}

func TestThreeSum(t *testing.T) {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
