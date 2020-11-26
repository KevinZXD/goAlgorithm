package main

type LinkNode struct {
	val  int
	next *LinkNode
}

func swapKLink(head *LinkNode, k int) *LinkNode {
	if head == nil {
		return head
	}
	phead := head
	num := 0
	for head != nil {
		tail := head
		for i := 0; i < k; i++ {
			head = head.next
			if i < k && head == nil {
				break
			}
		}
		tail, head = reverse(tail, head)
		if num == 0 {
			phead = head
		}
		num++

	}
	return phead
}
func reverse(head *LinkNode, tail *LinkNode) (*LinkNode, *LinkNode) {
	prev := tail.next
    p:=head
	for prev != tail {
		nex:=p.next
		p.next = prev
		prev = p
		p=nex
	}
	return tail, head

}
