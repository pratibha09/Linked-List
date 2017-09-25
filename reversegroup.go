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

func reversegroup(head *node, n int) node {
	var current *node = head
	var next *node = nil
	var prev *node = nil
	var count int = 0
	
	//reverse first n nodes of ll
	for current != nil && count < n {
		next = current.next
		current.next = prev
		prev = current
		current = next
		count++ 
	}

	//next is point to n+1 th node
	//calling recursively
	if next != nil {
		head.next = reversegroup(next, n)	
	}
	return prev
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
	head = reversegroup(head, 3)
	fmt.Printf("reverses linked list\n")
	printlist(head)

}


