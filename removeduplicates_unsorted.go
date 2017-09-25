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

func removeduplicate_unsorted(start *node){
	var ptr1 *node = start
	var ptr2 *node
	//var dup *node 
	
	for ptr1 != nil && ptr1.next != nil {
		ptr2 = ptr1
		for ptr2.next != nil {
			if ptr1.data == ptr2.next.data {
				var dup *node = ptr2.next
				ptr2.next = ptr2.next.next
				
			}else{
				ptr2 = ptr2.next
			}
		}
		ptr1 = ptr1.next
	}
}

func main(){
	var start *node
	start = &node{data : 12}
	push(&start, 19)
	push(&start, 21)
	push(&start, 13)
	push(&start, 14)
	push(&start, 18)
	push(&start, 20)
	fmt.Printf("The Linked List is : \n")
	printlist(start)
	
	removeduplicate_unsorted(start)
	fmt.Printf("Linked list after removing duplicates\n")
	printlist(start)

}

//time complexity is O(n^2)
