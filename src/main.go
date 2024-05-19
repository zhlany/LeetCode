package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
)

func main() {
	//arr := []int{15, 11, 7, 2}
	//ss := twoSum(arr, 9)
	//
	//s := "PAYPALISHIRING"
	//aa := lengthOfLongestSubstring(s)
	//aa := longestPalindrome(s)
	//a := []int{1, 2, 3, 4}
	//aa := convert2(s, 4)
	//aa := myAtoi("21474836460")
	//aa := isPalindrome(1001)
	//aa := isMatch("mississippi", "mis*is*ip*.")
	xx := []int{-1, 0, 1, 2, -1, -4}
	aa := threeSum(xx)
	fmt.Println("aa====", aa)
}

func twoSum(arr []int, tar int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > tar {
			continue
		}
		en := tar - arr[i]
		for j := 0; j < i; j++ {
			if arr[j] == en {
				return []int{j, i}
			}
		}
	}
	return []int{}
}

/*func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	target := node
	val := 0
	for l1 != nil || l2 != nil || val != 0 {
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		target.Next = &ListNode{Val: val % 10, Next: nil}
		target = target.Next
		val /= 10
	}
	return node.Next
}*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	lastMap := make(map[byte]int)
	// start 记录无序字符串的起点
	MaxLength, start := 0, 0
	for index, value := range []byte(s) {
		// 如果存在相同的字符，则更新起点
		if lastIndex, ok := lastMap[value]; ok && lastIndex >= start {
			start = lastIndex + 1
		}
		// 注意：这里必须判断 index-start 如果大于 MaxLength，因为相同字符可能造成 MaxLength 减少
		if index-start+1 > MaxLength {
			MaxLength = index - start + 1
		}
		lastMap[value] = index
	}

	return MaxLength
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var arr = append(nums1, nums2...)
	sort.Ints(arr)
	length := len(arr)
	if length%2 == 1 {
		return float64(arr[length/2])
	}
	return float64(arr[length/2]+arr[length/2-1]) / 2
}

func longestPalindrome(s string) string {
	start, end, length := 0, 0, len(s)
	i := 1
	str := ""
	for end < length {
		//找中心点
		if start+i < length && s[start+i] == s[start+i-1] {
			end = start + i
			i++
			continue
		}
		//重置中心点
		if end != start && s[start] != s[end] {
			start = end
			i = 1
		}
		//以s[start]和s[end]为中心，前后检测是否回文
		tmp := handPalindrome(start, end, s)
		if len(tmp) > len(str) {
			str = tmp
		}
		end++
	}
	if str == "" {
		return string(s[0])
	}
	return str
}

func handPalindrome(start, end int, s string) string {
	//记录移动的步长
	move := 0
	for j := 1; j <= start; j++ {
		if start-j >= 0 && end+j < len(s) {
			//如果移动后，前后字符不相等，则退出
			if s[start-j] != s[end+j] {
				break
			}
			move = j
		}
	}
	start -= move
	end += move
	return s[start : end+1]
}

func convert(s string, numRows int) string {
	//先计算一下可能需要的列数
	cols := numRows
	//这里+2是为了防止len(s)%numRows == numRows - 1 || numRows - 2
	n := len(s)/numRows + 2
	if numRows > 2 {
		//计算一个有多少小Z, +1 防止有余
		n = len(s)/(2*numRows-2) + 1
		//-2是减去头和尾的重复计算, +1 是加上最长一列
		cols = numRows - 2 + 1
		n = n*cols + 1 //总列数
	}
	fmt.Println("nn:====", n)

	arr := make([][]byte, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]byte, numRows)
	}
	keys := make(map[int][]byte)
	i, k := 0, 0
	index, pos := 0, 0
	up := false
	for k < len(s) {
		if index == 0 {
			up = false
		}
		if index < numRows && !up {
			arr[i][index] = s[k]
			if index == numRows-1 {
				pos = numRows
				up = true
			} else {
				index++
				pos = index
			}
		} else if up {
			index--
			i++
			if index == 0 {
				continue
			}
			arr[i][index] = s[k]
			pos = index
		}
		keys[pos] = append(keys[pos], s[k])
		k++
	}
	str := ""
	for key := 1; key <= numRows; key++ {
		str += string(keys[key])
		fmt.Printf("\n str::===%v", str)
	}
	for _, bytes := range arr {
		fmt.Printf("\n   %s", len(bytes))
		for _, b := range bytes {
			fmt.Printf("%v ", string(b))
		}
		//fmt.Println("ss:==", g, "arr::====", string(bytes))
	}
	fmt.Println("keys", keys)

	return str
}

func convert2(s string, numRows int) string {
	length := len(s)
	if numRows < 2 || length <= numRows {
		return s
	}

	ret := make([]byte, length)
	step := numRows<<1 - 2 // 2 * numRows - 2 步长
	index := 0             // ret中的下标
	i := 0                 // 原始字符串的下标

	// 遍历行
	for row := 0; row < numRows; row++ {
		// 遍历每一行的数据
		for i = row; i < length; i += step {
			ret[index] = s[i]
			index++
			// 如果不是第一行和最后一行，每一行会多加一个字符
			// middle := i + step - row<<1 // (2*(numRows-row) - 2)
			if row != 0 && row != numRows-1 && i+step-row<<1 < length {
				ret[index] = s[i+step-row<<1]
				index++
			}
		}
	}
	return string(ret)
}

func reverse(x int) int {
	rev := 0
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return rev
}

func myAtoi(s string) int {
	str := strings.TrimSpace(s)
	if str == "" {
		return 0
	}
	negative := 1
	if str[0] == '-' {
		negative = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	ret := 0
	for _, ch := range str {
		if ch < '0' || ch > '9' {
			break
		}
		ret = 10*ret + int(ch-'0')
		if negative*ret > math.MaxInt32 {
			return math.MaxInt32
			break
		} else if negative*ret < math.MinInt32 {
			return math.MinInt32
			break
		}
	}

	return negative * ret
}

//回文数整数转为字符串来解决
/*方法2: 从右往左算，原值取除以10＋商值，循环直至终，结果与原值比较
    tmp := 0
	x1 := x
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
}*/
func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	if x/10 == 0 {
		return true
	}
	//计算输入的int有几位数
	n := int(math.Log10(float64(x))) + 1

	//先把最高位和最地位获取出来
	xMax := x / int(math.Pow10(n-1))
	xMin := x % 10
	if xMax != xMin {
		return false
	}
	//从第二位算起
	i := 2
	xMin = x / 10
	for i <= n-1 {
		if i == n-1 {
			break
		}
		xMax = (x / int(math.Pow10(n-2))) % 10
		if xMax != xMin%10 {
			return false
		}
		xMin /= 10
		i++
		n--
	}
	return true
}

//未提交
func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}
	sl, pl := len(s), len(p)
	match := func(i int, str string) int {
		if s[i] != str[0] {
			return i
		}
		for k := i; k < sl; k++ {
			if s[k] != str[0] {
				return k
			}
		}
		return sl - 1
	}
	o, a := false, false
	index := 0
	check := func(tmpI int) bool {
		for i := tmpI; i < pl; {
			if i+2 < pl && p[i+1] == s[index] && p[i+2] == '*' {
				return true
			}
			i += 2
		}
		return false
	}
	pi := 0
	for i, v := range p {
		if v == '*' {
			a = true
			ch := string(p[i-1])
			newIndex := match(index, ch)
			if newIndex == index && !check(i) {
				return false
			}
			index = newIndex
			pi = i
			continue
		}
		if v == '.' {
			o = true
			index++
			pi = i
			continue
		}
		if v != '*' && v != '.' && i+1 < pl && p[i+1] != '*' {
			if s[index] != uint8(v) {
				return false
			}
			index++
		}
		pi = i
	}
	if pi != pl-1 || index != sl-1 {
		return false
	}
	if !o && !a && s != p {
		return false
	}
	return true
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ret := 0
	min := func(m, n int) int {
		if m > n {
			return n
		}
		return m
	}
	for i < j {
		s := (j - i) * min(height[i], height[j])
		if s > ret {
			ret = s
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ret
}

/*
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
*/
func intToRoman(num int) string {
	// 罗马数字字母表
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	// 罗马数字对应数值表
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var result strings.Builder
	for i := 0; num > 0; i++ {
		// 如果num比当前的value[i]大，就减去value[i]，并把对应的罗马数字添加到结果中
		for num >= value[i] {
			num -= value[i]
			result.WriteString(roman[i])
		}
	}
	return result.String()
}

func romanToInt(s string) int {
	// 罗马数字字母表
	M := map[string]int{
		"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40,
		"L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000,
	}
	R := map[string]bool{"CM": true, "CD": true, "XC": true, "XL": true, "IX": true, "IV": true}

	i, ret := 0, 0
	for i < len(s) {
		if (s[i] == 'C' || s[i] == 'X' || s[i] == 'I') && (i+1 < len(s)) {
			tmpStr := s[i : i+2]
			if _, ok := R[tmpStr]; ok {
				ret += M[tmpStr]
				i += 2
				continue
			}
		}
		ret += M[s[i:i+1]]
		i++
	}
	return ret
}

//no
func threeSum(nums []int) [][]int {
	l := len(nums)
	arr := make([][]int, 0)

	check := func(obj []int) bool {
		equal := false
		for _, e := range arr {
			equal = reflect.DeepEqual(obj, e)
		}
		return equal
	}

	k := l - 1
	for i := 0; i < l-2; i++ {
		j := i + 1
		k -= i
		for ; j < k; j++ {
			if nums[i]+nums[j]+nums[k] == 0 {
				if !check([]int{nums[i], nums[j], nums[k]}) {
					arr = append(arr, []int{nums[i], nums[j], nums[k]})
				}
			}
			k--
		}
	}

	return arr
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	return []string{digits}
}
