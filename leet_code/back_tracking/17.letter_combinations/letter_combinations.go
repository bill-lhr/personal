package letter_combinations

/**
*Description: 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
链接: https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
*/

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	list := make([]byte, 0, len(digits))
	m := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var backTracking func(int)

	backTracking = func(idx int) {
		if len(list) == len(digits) {
			res = append(res, string(list))
			return
		}
		digit := digits[idx]
		for _, letter := range m[digit] {
			list = append(list, letter)
			backTracking(idx + 1)
			list = list[:len(list)-1]
		}
	}
	backTracking(0)

	return res
}
