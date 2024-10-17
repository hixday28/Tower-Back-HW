package main

import "fmt"

type node struct {
	val  int
	next *node
}

type queue struct {
	first *node
	last  *node
}

func (q *queue) Add(new_val int) {
	new_node := &node{val: new_val}
	if q.last == nil {
		q.first = new_node
		q.last = new_node
	} else {
		q.last.next = new_node
		q.last = new_node
	}
}

func (q *queue) Pop() int {
	if q.first == nil {
		return 0
	}
	value := q.first.val
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	return value
}

func (q *queue) IsExist() bool {
	return q.first != nil
}

func main() {
	q := &queue{}
	n := 5

	for i := 0; i < n; i++ {
		q.Add(i)
	}

	fmt.Println("В очереди есть значения?", q.IsExist())

	for i := 0; i < n-1; i++ {
		value := q.Pop()
		fmt.Println("Удаляем ", value)
	}

	fmt.Println("В очереди есть значения?", q.IsExist())
	value := q.Pop()
	fmt.Println("Удаляем ", value)
	fmt.Println("В очереди есть значения?", q.IsExist())

}
