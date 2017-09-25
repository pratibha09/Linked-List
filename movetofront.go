package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func push(head **node, new_data int){
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head
	*head = newnode
}

func printlist(n *node){
	for n != nil {
		fmt.Printf("% d\n", n.data)
		n = n.next
	}
}

func movetofront(head **node) {
	
	if *head == nil || (*head).next == nil {
		return
	}
	var seclast *node = nil
	var last *node = *head
	for last.next != nil {
		seclast = last
		last = last.next
	}
	seclast.next = nil
	last.next = *head
	*head = last
}

func main(){
	var head *node
	head = &node{data : 12}
	push(&head, 12)
	push(&head, 12)
	push(&head, 13)
	push(&head, 14)
	push(&head, 19)
	push(&head, 20)
	fmt.Printf("The Linked List is : \n")
	printlist(head)
	movetofront(&head)
	fmt.Printf("the linked list after removing last to first\n")
	printlist(head)
}
