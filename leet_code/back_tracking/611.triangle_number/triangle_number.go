package triangle_number

import (
	"sort"
)

/**
 * @Description: 给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。
	链接: https://leetcode.cn/problems/valid-triangle-number/

	元素重复、不重复选择的组合问题
*/

// 1.回溯（超时）
func triangleNumber1(nums []int) int {
	cnt := 0
	list := make([]int, 0)
	//sort.Ints(nums)
	var backTracking func(int)

	backTracking = func(index int) {
		if len(list) == 3 {
			if check(list[0], list[1], list[2]) {
				cnt++
			}
			return
		}

		for i := index; i < len(nums); i++ {
			//if i > index && nums[i] == nums[i-1] {
			//	continue
			//}
			list = append(list, nums[i])
			backTracking(i + 1)
			list = list[:len(list)-1]
		}
	}
	backTracking(0)

	return cnt
}

func check(a, b, c int) bool {
	if a+b <= c {
		return false
	}
	if a+c <= b {
		return false
	}
	if b+c <= a {
		return false
	}
	return true
}

// 2.双指针
/*  1）确定一位下标为i
2) 下标j表示第二位，k表示能在i,j条件下能组成三角形的最大下标
*/
func triangleNumber2(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	cnt := 0
	for i := 0; i <= len(nums)-3; i++ {
		j, k := i+1, i+2
		for ; j <= len(nums)-2; j++ {
			for k < len(nums) && (k <= j || nums[i]+nums[j] > nums[k]) {
				k++
			}
			cnt += k - 1 - j
		}
	}
	return cnt
}
