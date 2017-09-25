package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
	prev *node
}

func merge(first *node, second *node) *node {
	if first == nil {
		return second
	}
	if second == nil{
		return first
	}
	if first.data < second.data {
		first.next = merge(first.next, second)
		first.next.prev = first
		first.prev = nil
		return second
	}else{
		second.next = merge(first, second.next)
		second.next.prev = second
		second.prev = nil
		return second
	}
}

func mergesort(head *node) *node{
	if head == nil || head.next == nil {
		return head
	}
	var second *node = split(head)
	head = mergesort(head)
	second = mergesort(second)
	return merge(head, second)
}

func insert(head **node, data int){
	temp := new(node)
	temp.data = data
	temp.next = nil
	temp.prev = nil
	if *head == nil {
		*head =temp
	}else{
		temp.next = *head
		(*head).prev = temp
		(*head) = temp
	}
}

func printlist(head *node){
	if head == nil {
		return
	} 
	var temp *node = head
	fmt.Printf("\nForward traversal using next pointer\n")
	for head != nil {
		fmt.Printf("%d\n", head.data)
		temp = head
		head = head.next
	}
	fmt.Printf("\nBackward traversal using prev pointer\n")
	for temp != nil {
		fmt.Printf("%d\n", temp.data)
		temp = temp.prev
	}
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func split(head *node) *node{
	var fast *node = head
	var slow *node = head
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	var temp *node = slow.next
	slow.next = nil
	return temp
}

func main(){
	var head *node = &node{data : 6}
	insert(&head, 5)
	insert(&head, 8)
	insert(&head, 9)
	insert(&head, 1)
	insert(&head, 4)
	insert(&head, 11)
	insert(&head, 13)
	head = mergesort(head)
	fmt.Printf("Linked list after sorting\n")
	printlist(head)
}

//time complexity is O(nlogn)0

