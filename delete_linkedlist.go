package main

import(
	"fmt"
)

type node struct{
	data int
	next *node
}

func push(head **node, new_data int){
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head
	*head = newnode
}


func printlist(n *node) {
	for n != nil {
		fmt.Printf("%d\n", n.data)
		n = n.next
	}
	fmt.Printf("null\n")
}

func deletelist(head **node) {
	var current *node = *head
	var next *node
	for current != nil {
		next = current.next 
		current = next
	}
	*head = nil

}

func main() {
	var head *node
	head = &node{data : 7}
	push(&head, 6)
	push(&head, 5)
	push(&head, 3)
	fmt.Printf("linked list is \n")
	printlist(head)
	fmt.Printf("deleting linked list\n")
	deletelist(&head)
	printlist(head)
	fmt.Printf("linkedlist deleted\n")

}


