package main

import (
	"fmt"
)

type node struct {
	data int
	right *node
	left *node
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

func merge(a *node, b *node) *node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	var result *node 
	if a.data < b.data {
		result = a
		result.next = merge(a.next , b)
	}else{
		result = b
		result.next = merge(a, b.next)
	}
	return result
}

func flatten(head *node) *node{
	if head == nil || head.right == nil {
		return head
	}
	return merge(head, flatten(head.right))

}

func main(){
	var head *node
	head = &node{data : 30}
	push(&head, 8)
	push(&head, 7)
	push(&head, 5)
	push(&(head.right), 20)
	push(&(head.right), 10)
	push(&(head.right.right), 50)
	push(&(head.right.right), 22)
	push(&(head.right.right), 19)
	push(&(head.right.right.right), 45)
	push(&(head.right.right.right), 40)
	push(&(head.right.right.right), 35)
	push(&(head.right.right.right), 20)
	fmt.Printf("The Linked List is : \n")
	printlist(head)
	fmt.Println()
	head = flatten(head)
	
	printlist(head)

}
