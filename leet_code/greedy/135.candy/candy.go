/**
 * @Author: lihaoran
 * @Date: 2022/7/24 16:51
 */

package candy

import "fmt"

/*
 * @Description:
	n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
	你需要按照以下要求，给这些孩子分发糖果：
	每个孩子至少分配到 1 个糖果。
	相邻两个孩子评分更高的孩子会获得更多的糖果。
	请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
	链接：https://leetcode.cn/problems/candy
*/
func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}

	// 正向遍历 保证每个孩子比左边分高时糖果多一
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	fmt.Println(candies)

	// 从右往左 保证比右边孩子分高时 糖果也多
	for j := len(ratings) - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] && candies[j] <= candies[j+1] {
			candies[j] = candies[j+1] + 1
		}
	}
	fmt.Println(candies)

	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}
