/**
 * @Author: lihaoran
 * @Date: 2022/8/27 20:04
 */

package mid

/**
 * @Description: 你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。
	在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，表示如果要学习课程ai 则 必须 先学习课程 bi 。
	例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
	请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
	链接：https://leetcode.cn/problems/course-schedule
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	coursesList := make([]int, 0)        // 课程顺序列表
	relayNums := make([]int, numCourses) // 下边表示课程，值表示依赖的课程数量
	relayMap := make(map[int][]int)      // key: 选修的课程 value: 依赖key课程的课程列表
	queue := make([]int, 0)

	for _, courses := range prerequisites {
		relayNums[courses[0]]++
		relayMap[courses[1]] = append(relayMap[courses[1]], courses[0])
	}
	for course, relayNum := range relayNums {
		if relayNum == 0 {
			queue = append(queue, course)
		}
	}
	for len(queue) > 0 {
		course := queue[0]
		coursesList = append(coursesList, course)
		queue = queue[1:]
		for _, c := range relayMap[course] {
			relayNums[c]--
			if relayNums[c] == 0 {
				queue = append(queue, c)
			}
		}
	}

	return len(coursesList) == numCourses
}
