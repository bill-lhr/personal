package __19

func maxSlidingWindow(nums []int, k int) []int {
	left, right := 0, 0
	res := make([]int, 0)
	stack := make([]int, 0)
	for right < len(nums) {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[right] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, right)
		right++
		if right-left == k {
			res = append(res, stack[0])
			if left == stack[0] {
				stack = stack[1:]
			}
			left++
		}
	}
	return res
}
