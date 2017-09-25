package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func push (head **node, new_data int) {
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head
	*head = newnode
}

func search(head *node, x int) bool {
	var current *node = head
	if current != nil {
		if current.data == x {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	var head *node
	head = &node{data : 3}
	push(&head, 5)
	push(&head, 7)
	push(&head, 8)
	push(&head, 9)
	search(head, 7)
	fmt.Println("present")

}


