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
	for n!= nil {
		fmt.Printf("%d\n", n.data)
		n = n.next
	}
}

func deletenode(n *node){
	var temp *node = n.next
	n.data = temp.data
	n.next = temp.next
}


func main(){
	var head *node
	head = &node{data : 7}
	push(&head, 6)
	push(&head, 8)
	push(&head, 7)
	fmt.Printf("The Linked list is: \n")
	printlist(head)
	fmt.Printf("deleting node\n")
	deletenode(head)
	printlist(head)

	fmt.Printf("deleting node\n")
	deletenode(head.next)
	printlist(head)

	
}

