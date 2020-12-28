package leetcode

import (
	"math"
	"sort"
	"strings"
)

func NumJewelsInStones(J string, S string) int {
	lj := len(J)
	ls := len(S)
	count := 0
	jewels := make(map[byte]bool, lj)
	for i := 0; i < lj; i++ {
		jewels[J[i]] = true
	}
	for i := 0; i < ls; i++ {
		if jewels[S[i]] {
			count++
		}
	}
	return count
}

func IsToeplitzMatrix(m [][]int) bool {
	lm := len(m)    // 行数
	ln := len(m[0]) // 列数
	for i := 0; i < lm-1; i++ {
		for j := 0; j < ln; j++ {
			if i+1 <= lm-1 && j+1 <= ln-1 && m[i][j] != m[i+1][j+1] {
				return false
			}
		}

	}
	return true
}

func Average(salary []int) float32 {
	max := -1
	min := math.MaxInt64
	cntOfMax := 0
	cntOfMin := 0
	sum := 0
	for i := 0; i < len(salary); i++ {
		sum += salary[i]
		if salary[i] < min {
			min = salary[i]
			cntOfMin = 0
		}
		if salary[i] == min {
			cntOfMin++
		}
		if salary[i] > max {
			max = salary[i]
			cntOfMax = 0
		}
		if salary[i] == max {
			cntOfMax++
		}
	}
	return float32((sum - (max*cntOfMax + min*cntOfMin))) / float32((len(salary) - cntOfMax - cntOfMin))
}
func UniqueOccurrences(arr []int) bool {
	l := len(arr)
	m := make(map[int]int, 0)
	for i := 0; i < l; i++ {
		if _, ok := m[arr[i]]; ok {
			m[arr[i]]++
		} else {
			m[arr[i]] = 1
		}

	}
	r := make(map[int]int, 0)
	for _, v := range m {
		if _, ok := r[v]; ok {
			return false
		}
		r[v]++
	}
	return true
}

type S struct {
	index int
	val   int
}
type SliceS []S

func (ss SliceS) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
func (ss SliceS) Less(i, j int) bool {
	if ss[i].val < ss[j].val {
		return true
	}
	if ss[i].val == ss[j].val {
		return ss[i].index < ss[j].index
	}
	return false
}
func (ss SliceS) Len() int {
	return len(ss)
}
func KWeakestRows(mat [][]int, k int) []int {

	// l := len(mat)
	var Ss SliceS //每行战斗力值
	for i, vs := range mat {
		powerOfcur := 0
		for _, v := range vs {
			if v == 1 {
				powerOfcur++
			}
		}
		s := S{index: i, val: powerOfcur}
		Ss = append(Ss, s)
	}
	sort.Sort(Ss)
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, Ss[i].index)
	}
	return res
}
func TwoSum(nums []int, target int) []int {
	l := len(nums)
	// left, right := 0, l-1
	// // res := make([]int, 0)
	// for left < right && right >= 0 && left <= l-1 {
	// 	sum := nums[left] + nums[right]
	// 	if sum == target {
	// 		return []int{nums[left], nums[right]}
	// 	}
	// 	if sum > target {
	// 		right--
	// 	}
	// 	if sum < target {
	// 		left++
	// 	}
	// }
	// return []int{0, 0}
	m := make(map[int]int, l)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = target - v
		}
	}
	for _, v := range m {
		if _, ok := m[v]; ok {
			return []int{v, m[v]}
		}
	}
	return []int{0, 0}

}

func countConsistentStrings(allowed string, words []string) int {
	lOfAllowed := len(allowed)
	allMap := make(map[byte]bool)
	res := 0
	for i := 0; i < lOfAllowed; i++ {
		allMap[allowed[i]] = true
	}
	for _, str := range words {
		isThisOk := true
		for _, k := range str {
			if !allMap[byte(k)] {
				isThisOk = false
				break
			}
		}
		if isThisOk {
			res++
		}
	}
	return res
}
func InvalidateIP(ip string) string {
	var invalidIP strings.Builder
	len := len(ip)
	for i := 0; i < len; i++ {
		if ip[i] != '.' {
			invalidIP.WriteByte(ip[i])
		} else {
			invalidIP.WriteString("[.]")
		}
	}

	return invalidIP.String()
}

func SubtractProductAndSum(n int) int {
	//各数之积-各数之和
	sum := 0
	mul := 1
	for n != 0 {
		cur := n % 10
		n = n / 10
		sum = sum + cur
		mul = mul * cur
	}
	return mul - sum
}

func FirstUniqChar(s string) int {
	l := len(s)
	m := make(map[rune]int, l) //字符：出现次数
	for _, c := range s {
		if _, exist := m[c]; exist {
			//存在的话
			m[c] = m[c] + 1
		} else {
			m[c] = 1
		}
	}
	for i, c := range s {
		if m[c] == 1 {
			return i
		}
	}

	return -1
}
func numberOfSteps(n int) int {
	cnt := 0
	for n != 0 {
		if n%2 == 1 {
			n = n - 1
			cnt++

		}
		if n%2 == 0 {
			n = n / 2
			cnt++
		}
	}
	return cnt
}

func MinPartitions(n string) int {
	var max int
	max = 0
	for _, r := range n {
		if max < int(r-'0') {
			max = int(r - '0')
		}
	}
	return max
}

// 求对角线的和
func DiagonalSum(mat [][]int) int {
	n := len(mat)
	//左上-》右下
	i, j := 0, 0
	sum := 0
	for i < n && j < n {
		sum += mat[i][j]
		sum += mat[i][n-1-i]
		i++
		j++
	}
	//右上-》左下
	// i, j = 0, n-1
	// for i < n && j >= 0 {
	// 	sum += mat[i][j]
	// 	i++
	// 	j--
	// }
	if n%2 == 1 {
		sum -= mat[n/2][n/2]
	}
	return sum
}
func calculate(s string) int {
	x, y := 1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			x = 2*x + y
		}
		if s[i] == 'B' {
			y = 2*y + x
		}
	}
	return x + y
}

func shuffle(nums []int, n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, nums[i])
		res = append(res, nums[i+n])
	}
	return res
}
func minCount(coins []int) int {
	cnt := 0
	for i := 0; i < len(coins); i++ {
		if coins[i] <= 2 {
			cnt++
		} else {
			cnt += (coins[i] + 1) / 2
		}
	}
	return cnt
}

// 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。
func numIdenticalPairs(nums []int) int {
	cnt := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				cnt++
			}
		}
	}
	return cnt
}

// 根据评分分糖果，每个孩子至少一个，相邻孩子较高的必须糖果更多
func Candy(ratings []int) int {
	l := len(ratings)
	candy := make([]int, l)
	for i := 0; i < l; i++ {
		candy[i] = 1
	}
	// 从左往右比较
	for i := 0; i < l-1; i++ {
		if ratings[i+1] > ratings[i] {
			candy[i+1] = candy[i] + 1
		}
	}
	// 从右往左比较
	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candy[i] = (int)(math.Max(float64(candy[i+1]+1), float64(candy[i])))
		}
	}
	sum := 0
	for _, v := range candy {
		sum += v
	}
	return sum
}

func SmallerNumbersThanCurrent(nums []int) []int {
	l := len(nums)
	cntOfFre := make([]int, 101) //nums值域
	for i := 0; i < l; i++ {
		cntOfFre[nums[i]]++
	}
	for i := 1; i < 101; i++ {
		cntOfFre[i] = cntOfFre[i] + cntOfFre[i-1]
	}
	res := make([]int, l)
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			res[i] = 0
			continue
		}
		res[i] = cntOfFre[nums[i]-1]
	}
	return res
	// cntOfLessI := make([]int, l)
	// for i := 0; i < l; i++ {
	// 	tmp := 0
	// 	for j := 0; j < l; j++ {
	// 		if nums[j] < nums[i] {
	// 			tmp++
	// 		}
	// 	}
	// 	cntOfLessI[i] = tmp
	// }
	// return cntOfLessI
}

func FindContentChildren(g []int, s []int) int {
	//遍历饼干
	gn := len(g)
	sn := len(s)
	sort.Ints(g)
	sort.Ints(s)
	p, q := 0, 0
	cnt := 0
	for p < gn && q < sn {
		if s[q] >= g[p] {
			cnt++
			q++
			p++
		} else {
			q++
		}
	}
	return cnt
}

func Maximum69Number(num int) int {
	res := make([]int, 0)
	cur := 0
	for num != 0 {
		cur = num % 10
		num = num / 10
		res = append([]int{cur}, res...)
	}
	for i := 0; i < len(res); i++ {
		if res[i] == 6 {
			res[i] = 9
			break
		}
	}
	// 转为int
	p := 0
	val := res[0]
	for p < len(res)-1 {
		val = val*10 + res[p+1]
		p++
	}
	return val
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReversePrint(head *ListNode) []int {
	newHead := reverse(head)
	res := make([]int, 0)
	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}
	return res
}

// 递归链表倒置
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newList := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newList
}

func maxDepth(s string) int {
	// l := len(s)
	maxdp := 0
	curdp := 0
	for _, c := range s {
		if c == '(' {
			curdp++
			maxdp = max(curdp, maxdp)
		}
		if c == ')' {
			curdp--
		}
	}
	return maxdp
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func createTargetArray(nums []int, index []int) []int {
	len := len(nums)
	for i := 0; i < len; i++ {

	}
}
