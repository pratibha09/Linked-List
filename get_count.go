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

func getCount(head *node)int {
	count := 0
	var current *node = head
	for current != nil {
		count++
		current = current.next
	}
	return count
}


func main(){
	var head *node
	head = &node{data : 3}
	push(&head, 10)
	push(&head, 9)
	push(&head, 8)
	fmt.Println("the linked list is", getCount(head))
	
}

