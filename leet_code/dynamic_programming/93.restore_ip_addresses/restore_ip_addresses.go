/**
 * @Author: lihaoran
 * @Date: 2022/8/18 22:14
 */

package restore_ip_addresses

import (
	"strconv"
	"strings"
)

/**
 * @Description:
	有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
	例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
	给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入'.' 来形成。你 不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
	链接：https://leetcode.cn/problems/restore-ip-addresses
*/

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	list := make([]string, 0)
	var backTracking func(int)

	backTracking = func(idx int) {
		// 正常退出条件
		if len(list) == 4 && idx == len(s) {
			res = append(res, strings.Join(list, "."))
			return
		}
		// 剪枝 1.非法字符 2.剩余的长度超长 3.数字组合范围不在0-255
		for i := idx; i < idx+3 && i < len(s); i++ {
			// 非法字符
			b := s[i]
			if b < '0' || b > '9' {
				break
			}
			// 合法数字及范围
			numStr := s[idx : i+1]
			if strings.HasPrefix(numStr, "0") && len(numStr) > 1 {
				break
			}
			num, err := strconv.Atoi(numStr) // 问题00能不能解析成功
			if err != nil {
				continue
			}
			if num > 255 {
				break
			}
			// 是否超长
			if len(s)-i-1 > (3-len(list))*3 {
				continue
			}
			list = append(list, numStr)
			backTracking(i + 1)
			list = list[:len(list)-1]
		}
	}
	backTracking(0)

	return res
}
