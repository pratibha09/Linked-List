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

func printmiddle(head *node){
	var slow_ptr *node = head
	var fast_ptr *node = head
	if head != nil {
		for fast_ptr != nil &&  fast_ptr.next != nil {
			fast_ptr = fast_ptr.next.next
			slow_ptr = slow_ptr.next
		}
		fmt.Printf("the middle element is %d \n", slow_ptr.data)
	}
}

func printlist(n *node) {
	for n != nil {
		fmt.Printf("%d \n", n.data)
		n = n.next
	}
	fmt.Printf("Null\n")
}

func main() {
	var head *node
	head =&node{data : 7}

		push(&head, 6)
		push(&head, 5)	
		push(&head, 8)
		printlist(head)
		printmiddle(head)
}

