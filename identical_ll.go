package main 

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func identical(a *node, b *node) bool {
	for a != nil && b != nil {
		if a.data != b.data {
			return false
		}
		a = a.next
		b = b.next
	}
	return a == nil && b == nil
}

func push(head **node, new_data int) {
	newnode := new(node)
	newnode.data = new_data
	newnode.next = *head
	*head = newnode
}


func main(){
	var a *node = nil
	var b *node = nil
	push(&a, 1)
	push(&a, 2)
	push(&a, 3)
	push(&b, 1)
	push(&b, 2)
	push(&b, 3)
	if identical(a, b){
		fmt.Printf("identical\n")
	}else{
		fmt.Printf("not identical")
	}
}

