package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	next *Node
	prev *Node
	val  int
	key  int
}

type LRU struct {
	cache map[int]*Node
	head  *Node
	tail  *Node
	cap   int
	len   int
}

func newNode(key int, val int) *Node {
	return &Node{val: val, key: key}
}

func newLRU(cap int) *LRU {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &LRU{
		cache: make(map[int]*Node, cap),
		head:  head,
		tail:  tail,
		cap:   cap,
		len:   0,
	}
}

func (c *LRU) put(key int, val int) {
	if _, ok := c.cache[key]; !ok {
		// 判断是否需要删除尾节点
		if c.cap == c.len {
			// 删除尾节点
			c.removeNode(c.tail.prev)
			c.len--
		}
		// 添加到头部
		node := newNode(key, val)
		c.addToHead(node)
		c.cache[key] = node
		c.len++
	} else {
		c.moveToHead(c.cache[key])
		c.cache[key].val = val
	}
}

func (c *LRU) get(key int) int {
	if _, ok := c.cache[key]; !ok {
		return -1
	}
	node := c.cache[key]
	c.moveToHead(node)
	return node.val
}

func (c *LRU) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRU) addToHead(node *Node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRU) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}
func (c *LRU) print() {
	builder := strings.Builder{}
	node := c.head
	for node != nil {
		builder.WriteString(strconv.Itoa(node.val) + "-> ")
		node = node.next
	}
	fmt.Print(builder.String() + "\n")
}

func main() {
	cache := newLRU(2)
	cache.put(1, 1)
	cache.put(2, 2)
	cache.print()
	cache.get(1)
	cache.print()
	cache.put(3, 3)
	cache.print()

}
