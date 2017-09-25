package main

import(
	"fmt"
)

type node struct {
	data int
	next *node
	prev *node
}

func push(head **node, new_data int){
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head
	newnode.prev = nil
	if *head != nil {
		(*head).prev = newnode
	}
	*head = newnode
}

func printlist(n *node) {
	var last *node
	fmt.Printf("\nTraversal in forward direction\n")
	for n != nil {
		fmt.Printf("%d\n", n.data)
		last = n
		n = n.next
	}
	
	fmt.Printf("\nTraversal in reverse direction\n")
	for last != nil {
		fmt.Printf("%d\n", last.data)
		last = last.prev
	}
}

func main(){
	var head *node
	head = &node{data : 7}
	push(&head, 5)
	push(&head, 4)
	push(&head, 2)
	fmt.Printf("\nCreated doubly linked list\n")
	printlist(head)
}
