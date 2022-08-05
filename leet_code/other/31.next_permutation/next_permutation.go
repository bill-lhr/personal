/**
 * @Author: lihaoran
 * @Date: 2022/8/5 22:07
 */

package next_permutation

/**
 * @Description: 给你一个整数数组 nums ，找出 nums 的下一个排列。
	链接: https://leetcode.cn/problems/next-permutation/

	下一个排列是比当前排列更大的最小的排列，要增大需要把后边大的元素往前换，但这个大元素要尽量小
	从右向左找第一对升序(i,i+1) 那么 [i+1,n)一定是降序的
	i就是要被向后换的小数，此时要找到右侧最小的比他大的数，即右侧第一个比i大的数，记作j
	交换(i,j) 此时i后边一定是降序，反转变成升序
*/
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] > nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		if j > i {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	reverse(nums, i+1, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
