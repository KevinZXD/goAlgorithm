package main

type DoubleLinkedListNode struct {
	key  string
	val  string
	next *DoubleLinkedListNode
	prev *DoubleLinkedListNode
}
type LRUCache struct {
	lruMap   map[string]*DoubleLinkedListNode
	head     *DoubleLinkedListNode
	tail     *DoubleLinkedListNode
	capacity int
}

type LRUCacheIn interface {
	put(key string, value string)
	get(key string) string
}

func (lruCache LRUCache) get(key string) string {
	node, ok := lruCache.lruMap[key]
	if ! ok {
		return ""
	}
	moveToHead(node, lruCache.head)
	return node.val
}

func (lruCache LRUCache) init(){

}
func (lruCache LRUCache) put(key string, value string) {
	node, ok := lruCache.lruMap[key]
	if ok {
		node.val = value
		moveToHead(node, lruCache.head)
	} else {
		node := DoubleLinkedListNode{key, value, nil, nil}
		cap := len(lruCache.lruMap)
		if cap >= lruCache.capacity {
			rm := lruCache.tail.prev
			remove(rm)
			delete(lruCache.lruMap, rm.key)

		}
		insertHead(&node, lruCache.head)
		lruCache.lruMap[key] = &node
	}

}
func insertHead(node *DoubleLinkedListNode, head *DoubleLinkedListNode) {
	next := head.next
	head.next = node
	node.prev = head
	node.next = next
	next.prev = node
}
func moveToHead(node *DoubleLinkedListNode, head *DoubleLinkedListNode) {
	remove(node)
	insertHead(node, head)
}
func remove(node *DoubleLinkedListNode) {
	front := node.prev
	end := node.next
	front.next = end
	end.prev = front
}
func main() {
	lruMap := make(map[string]*DoubleLinkedListNode)
	capacity := 2
	head := DoubleLinkedListNode{"-1", "-1", nil, nil}
	tail := DoubleLinkedListNode{"-1", "-1", nil, nil}
	head.next = &tail
	head.prev = &tail
	tail.next = &head
	tail.prev = &head
	lru := LRUCache{lruMap, &head, &tail, capacity}
	lru.put("name", "zxd")
	lru.put("sex", "female")
	lru.put("age", "22")
	lru.get("sex")
	lru.put("mobile", "3333")
	print("11111")
}
