package main
//
//type DoubleLinkedListNode struct {
//	key  string
//	val  string
//	next *DoubleLinkedListNode
//	prev *DoubleLinkedListNode
//}
//
//func get(key string, lruMap map[string]*DoubleLinkedListNode, head *DoubleLinkedListNode) string {
//	node, ok := lruMap[key]
//	if ! ok {
//		return ""
//	}
//	moveToHead(node, head)
//	return node.val
//}
//
//func put(key string, value string, lruMap map[string]*DoubleLinkedListNode, capacity int, head *DoubleLinkedListNode, tail *DoubleLinkedListNode) {
//	node, ok := lruMap[key]
//	if ok {
//		node.val = value
//		moveToHead(node, head)
//	} else {
//		node := DoubleLinkedListNode{key, value, nil, nil}
//		cap := len(lruMap)
//		if cap >= capacity {
//			rm := tail.prev
//			remove(rm)
//			delete(lruMap, rm.key)
//
//		}
//		insertHead(&node, head)
//		lruMap[key] = &node
//	}
//
//}
//func insertHead(node *DoubleLinkedListNode, head *DoubleLinkedListNode) {
//	next := head.next
//	head.next = node
//	node.prev = head
//	node.next = next
//	next.prev = node
//}
//func moveToHead(node *DoubleLinkedListNode, head *DoubleLinkedListNode) {
//	remove(node)
//	insertHead(node, head)
//}
//func remove(node *DoubleLinkedListNode) {
//	front := node.prev
//	end := node.next
//	front.next = end
//	end.prev = front
//}
//func main() {
//	lruMap := make(map[string]*DoubleLinkedListNode)
//	capacity := 2
//	head := DoubleLinkedListNode{"-1", "-1", nil, nil}
//	tail := DoubleLinkedListNode{"-1", "-1", nil, nil}
//	head.next = &tail
//	head.prev = &tail
//	tail.next = &head
//	tail.prev = &head
//	put("name", "zxd", lruMap, capacity, &head, &tail)
//	put("sex", "female", lruMap, capacity, &head, &tail)
//	put("age", "22", lruMap, capacity, &head, &tail)
//	get("sex", lruMap,  &head)
//	put("mobile", "3333", lruMap, capacity, &head, &tail)
//	print("11111")
//}
