package main

func minNumberInRotateArray(array [] int) int {
	if len(array) == 0 {
		return -1
	}
	low := 0
	high := len(array) - 1
	mid := (low + high) / 2
	for {

		if array[mid] > array[high] {
			low = mid + 1
		} else {
			high = mid - 1
		}
		if low > high {
			break
		}
		mid = (low + high) / 2
	}
	return array[mid]

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	first := pHead
	var second *ListNode
	for {
		if pHead == nil {
			break
		}
		first = &ListNode{Val: pHead.Val}
		first.Next = second
		second = first
		pHead=pHead.Next
	}
	return first
}
func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	array1 := []int{4, 5, 6, 1, 2, 3}
	println(minNumberInRotateArray(array))
	println(minNumberInRotateArray(array1))
}
