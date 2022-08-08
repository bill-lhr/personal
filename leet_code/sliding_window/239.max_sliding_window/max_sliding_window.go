package max_sliding_window

/**
   *Description给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。
	返回 滑动窗口中的最大值 。
	链接: https://leetcode.cn/problems/sliding-window-maximum/
*/
func maxSlidingWindow(nums []int, k int) []int {
	left, right := 0, 0
	res := make([]int, 0)
	queue := make([]int, 0) // 存放下标
	for right < len(nums) {
		// 队列中小于当前值的出队
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[right] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, right)
		right++

		if right-left == k {
			// 当前窗口的最大值
			res = append(res, nums[queue[0]])
			// 队头下标是left，窗口缩小时要出队
			if queue[0] == left {
				queue = queue[1:]
			}
			left++
		}
	}
	return res
}
