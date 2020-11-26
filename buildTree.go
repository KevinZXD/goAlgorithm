package main

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func buildTree(pre []int, mid []int) Tree {
	var head Tree
	if len(pre) == 1 {

		head.value = pre[0]
		head.left = nil
		head.right = nil
		return head

	}
	head.value = pre[0]
	temp0 := mid.index(pre[0])
	length := len(mid[:temp0])
	if temp0 == 0 {
		head.left = nil
	} else {
		head = buildTree(pre[1:length+1], mid[1:temp0])
	}
	if length+1 > len(pre) {
		head.right = nil
	} else {
		head = buildTree(pre[length+1:], mid[temp0+1:])
	}
	return head
}
