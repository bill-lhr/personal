/**
 * @Author: lihaoran
 * @Description:
 * @Date: 2022/7/19 00:13
 */

package permutation

func permutation(s string) []string {
	res := make([]string, 0)
	needVisit := make(map[byte]int)
	list := make([]byte, 0)
	var backTracking func()

	for i := range s {
		needVisit[s[i]]++
	}

	backTracking = func() {
		if len(list) == len(s) {
			tmp := make([]byte, len(s))
			copy(tmp, list)
			res = append(res, string(list))
			return
		}

		tmpVisit := make(map[byte]struct{})
		for i := range s {
			b := s[i]
			if needVisit[b] == 0 {
				continue
			}
			if _, ok := tmpVisit[b]; ok {
				continue
			}
			tmpVisit[b] = struct{}{}
			list = append(list, b)
			needVisit[b]--
			backTracking()
			needVisit[b]++
			list = list[0 : len(list)-1]
		}
	}

	backTracking()
	return res
}
