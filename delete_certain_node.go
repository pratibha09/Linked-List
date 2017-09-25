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

func deletenode(head **node, position int){
	if head == nil {
		return
	}
	var temp *node = *head
	if position == 0 {
		*head = temp.next
		return
	}
	for i := 0; temp != nil && i < position; i++ {
		temp = temp.next
	}
	if temp == nil || temp.next == nil {
		return
	}
	var next *node = temp.next.next
	temp.next = next 
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
	deletenode(&head, 4)
	fmt.Printf("after deletion\n")
	printlist(head)
}
