package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func push(head **node, new_data int) {
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head
	*head = newnode
}

func Nthfromlast(head *node, n int){
	var length int = 0
	var temp *node = head
	for temp != nil {        //count the number of nodes in linked list
		temp = temp.next
		length++
	}
	if length < n {
		return
	}
	
	temp = head
	for i := 1; i < length-n+1; i++ {
		temp = temp.next
	}
	fmt.Printf("%d\n", temp.data)
	return
}

func printlist (n *node) {
	for n != nil {
		fmt.Printf("%d\n", n.data)
		n = n.next
	}
}

func main() {
	var head *node
	head = &node{data : 7}
	push(&head, 8)
	push(&head, 6)
	push(&head, 4)
	push(&head, 3)
	fmt.Printf("the linkedlist is:---\n")
	printlist(head)
	
	//Nthfromlast(head, 5)
	fmt.Printf("nth from last\n")
	Nthfromlast(head, 2)
	
}


