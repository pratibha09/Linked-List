package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func push(head_ref **node, new_data int) {
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head_ref
	*head_ref = newnode
}

func Append(head_ref **node, new_data int) {
	newnode := new(node)
	var last *node = *head_ref
	newnode.data = new_data
	newnode.next = nil
	
	if *head_ref == nil {
		last = last.next
	}
	last.next = newnode
	return
}


func printlist(n *node) {
	for n != nil {
		fmt.Printf("%d\n", n.data)
		n = n.next
	}
}

func main() {
	var head *node
	head = &node{data : 1}
	push(&head, 7)
	push(&head, 10)
	//Append(&head, 11)
	fmt.Println("created linked list is:")
	printlist(head)
}


