package main

import(
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

func reverse(head **node) {
	var prev *node = nil
	var current *node = *head
	var next *node
	for current != nil {
		next = current.next 
		current.next = prev
		prev = current
		current= next
	}
	*head = prev
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
	push(&head, 20)
	fmt.Printf("The Linked List is : \n")
	printlist(head)
	reverse(&head)
	fmt.Printf("The reversed linked list is\n")
	printlist(head)
}

