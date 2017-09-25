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

func insertionsort(head **node) {
	var sorted *node = nil
	var current *node = *head
	for current != nil {
		var next *node = current.next
		sortedinsert(&sorted, current)
		current = next
	}
	*head = sorted
}

func sortedinsert(head **node, new_node *node) {
	var current *node
	if *head == nil || (*head).data >= new_node.data {
		new_node.next = *head
		*head = new_node
	}else {
		current = *head	
		for current.next != nil && current.next.data < new_node.data {
			current = current.next
		} 
		new_node.next = current.next
		current.next = new_node
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
	insertionsort(&head)
	fmt.Printf("Linked list after sorting\n")
	printlist(head)

}
