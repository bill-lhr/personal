/**
 * @Author: lihaoran
 * @Date: 2022/7/24 11:40
 */

package reverse_words

import "strings"

/*
 * @Description:
	给你一个字符串 s ，颠倒字符串中 单词 的顺序。
	单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
	返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
	注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
	链接：https://leetcode.cn/problems/reverse-words-in-a-string
*/
func reverseWords(s string) string {
	list := strings.Split(s, " ")
	// 单个单次反转 去掉空字符串
	reverseList1 := make([]string, 0)
	for _, v := range list {
		if v == "" {
			continue
		}
		reverseList1 = append(reverseList1, reverse(0, len(s)-1, []byte(v)))
	}

	// 整体反转
	s = strings.Join(reverseList1, " ")
	b := []byte(s)
	return reverse(0, len(b)-1, b)
}

// 反转字符串 (i,j)之间的字符
func reverse(i, j int, b []byte) string {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
