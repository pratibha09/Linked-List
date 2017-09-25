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

func rotate(head **node, k int){
	if k == 0 {
		return	
	}

	var current *node = *head
	var count int = 1 
	for count < k && current != nil {
		current = current.next
		count++
	}
	if current == nil{
		return
	}
	
	var kth *node = current
	for current.next != nil {
		current = current.next
	}
	current.next = *head
	*head = kth.next
	kth.next = nil
}

func printlist(n *node){
	for n != nil {
		fmt.Printf("%d\n", n.data)
		n = n.next
	}
}

func main(){
	var head *node
	head = &node{data : 7}
	push(&head, 5)
	push(&head, 2)
	push(&head, 9)
	push(&head, 4)
	push(&head, 3)
	fmt.Printf("The Linked List is : \n")
	printlist(head)
	rotate(&head, 2)
	fmt.Printf("the rotated list is: \n")
	printlist(head)	

}

