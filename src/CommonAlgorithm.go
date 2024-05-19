package main

import "sort"

//二分查找
func search(num []int, target int) int {
	left, right := 0, len(num)-1
	sort.Ints(num)
	for left <= right {
		mid := (left + right) / 2
		if num[mid] == target {
			return mid
		} else if num[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

//快速排序
func quickSort(nums []int) {
	tar := nums[0]
	head, tail := 0, len(nums)-1
	for i := 1; i <= tail; {
		if nums[i] < tar {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		} else {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		}
	}
	quickSort(nums[:head])
	quickSort(nums[head+1 : tail])
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	cur := head
	//这里保存上一个结点信息
	node := &ListNode{}
	for cur != nil {
		//先保存下一个结点信息
		next := cur.Next
		//之后改变当前结点的next指针，指向上一个结点
		cur.Next = node
		//令node变更为当前结点，方便下一个结点指向它
		node = cur
		//cur变更为下一个结点
		cur = next
	}
	return node
}

//归并排序
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

//冒泡排序
func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-2-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

//选择排序
func selectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

//栈
