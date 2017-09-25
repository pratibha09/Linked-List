package main

import (
	"fmt"
)

type node struct{
	data int
	next *node
}

func push(head **node, new_data int) {
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head
	*head = newnode
}

func printlist(head *node){
	var temp *node = head
	for temp != nil {
		fmt.Printf("%d\n", temp.data)
		temp = temp.next
	}
}

func MdeleteN(head *node, m int, n int) {
	var curr *node = head
	var t *node
	var temp *node
	for count := 1; count < m && curr != nil; count++ {
		curr = curr.next
	}
	if curr == nil {
		return
	}
	t = curr.next
	for count := 1; count <= n && t != nil; count++ {
		temp = t
		t = t.next
	}
	curr.next = t
	curr = t
}

func main(){
	var m int = 2
	var n int = 3
	var head *node
	head = &node{data : 6}
	push(&head, 7)
	push(&head, 9)
	push(&head, 5)
	push(&head, 5)
	push(&head, 1)
	push(&head, 7)
	push(&head, 2)
	fmt.Printf("the linkedlist is: \n")
	fmt.Printf("M is %d and N is %d \n", m, n)
	printlist(head)
	MdeleteN(head, m, n)
	fmt.Printf("linkedlist after delettion\n")
	printlist(head)
	
}

