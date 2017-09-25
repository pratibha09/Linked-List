package main

import(
	"fmt"
)

type node struct {
	data int
	next *node
}

func push(head **node, new_data int) {
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head
	*head = newnode
}

func deletenode(head **node, data int) **node{
	var prev *node
	var temp *node = *head

	if temp != nil && temp.data == data {
		*head = temp.next
		return head
	}

	for temp != nil && temp.data != data {
		prev = temp
		temp = temp.next
	}

	if temp == nil {
		return head
	}
	prev.next = temp.next
	return head
}

func printlist(n *node) {
	for n != nil {
		fmt.Printf("%d\n", n.data)
		n = n.next
	}

}

func main() {
	var head *node
	head = &node{data : 0}
	push(&head, 8)
	push(&head, 9)
	push(&head, 5)
	fmt.Println("linked list  is")
	printlist(head)
	deletenode(&head, 8)
	fmt.Println("LInked list after deletion")
	printlist(head)

}


