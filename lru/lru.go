package main

import "fmt"

type Node struct {
	previous *Node
	next *Node
	value interface{} // this lets us use any type we want
}

type List struct {
	head *Node
	tail *Node
}

// The preceding argument (list *List) denotes a receiver;
// it means that any List can use this as a method, i.e.
// something := List()
// something.insert(5)
func (list *List) insert(insertedValue interface{}) {
	newElement := &Node{
		previous: list.tail,
		value: insertedValue,
		next: nil,
	}
	if list.head == nil {
		list.head = newElement
		list.tail = newElement
	} else {
		list.tail.next = newElement
		list.tail = newElement
	}
}

func (list *List) show() {
	element := list.head
	for element != nil {
		fmt.Printf("%+v ~ ", element.value)
		element = element.next
	}
	fmt.Printf("\n")
}

func main () {
	this := List{}
	for i := 0; i <= 5; i ++ {
		this.insert(i)
	}
	this.show()

}