package main

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) InsertHead(data interface{})  {

	node := &LinkNode{
		data:data,
		next:nil,
	}

	if p.head == nil && p.tail == nil {
		p.head = node
		p.tail = node
		return
	}

	node.next = p.head
	p.head = node
}

func (p *Link) InsertTail(data interface{}) {

	node := &LinkNode{
		data:data,
		next:nil,
	}

	if p.head == nil && p.tail == nil {
		p.head = node
		p.tail = node
		return
	}

	p.tail.next = node
	p.tail = node
}

func (p *Link) Traverse() {
	node := p.head

	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}



