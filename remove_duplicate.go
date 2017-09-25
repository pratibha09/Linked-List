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

func removeduplicate(head *node) {
	var current *node = head
	var next_next *node
	if current == nil {
		return
	}
	for current.next != nil {
		if current.data == current.next.data { 
			next_next = current.next.next
			current.next = next_next	
		}else{
			current = current.next
		}
	}
}

func printlist(n *node){
	for n != nil {
		fmt.Printf("% d\n", n.data)
		n = n.next
	}
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
	removeduplicate(head)
	fmt.Printf("After removing duplicates\n")
	printlist(head)
}

//time complexity is O(n)
