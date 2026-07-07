package gobyexample

import "fmt"

type genericsPerson struct {
	Name string
	Age  int
}

func (p genericsPerson) String() string {
	return fmt.Sprintf("name: %s age: %d", p.Name, p.Age)
}

// Generic function
// '~' means S can be any slice type whose underlying type is []E
func genericsSlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Generic Type
type genericsList[T any] struct {
	head, tail *genericsNode[T]
}

type genericsNode[T any] struct {
	next  *genericsNode[T]
	value T
}

// Generic method
func (l *genericsList[T]) push(v T) {
	if l.tail == nil { // list is empty
		l.head = &genericsNode[T]{value: v} // create a new node and assign it to head
		l.tail = l.head                     // assign the same node to tail
	} else { // list is not empty
		l.tail.next = &genericsNode[T]{value: v} // create a new node and assign it to the next of tail
		l.tail = l.tail.next                     // move the tail to the new node
	}
}

func (l *genericsList[T]) allElements() []T {
	var elements []T
	for node := l.head; node != nil; node = node.next {
		elements = append(elements, node.value)
	}
	return elements
}

func ShowGenericFunctionsAndTypes() {

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []string{"a", "b", "c", "d"}
	s3 := []bool{true, false, true}
	s4 := []genericsPerson{
		{Name: "Reda", Age: 30},
		{Name: "Ahmed", Age: 25},
		{Name: "Mohammed", Age: 35},
	}

	fmt.Println(genericsSlicesIndex(s1, 3))    // 2
	fmt.Println(genericsSlicesIndex(s2, "c"))  // 2
	fmt.Println(genericsSlicesIndex(s3, true)) // 0

	fmt.Println(genericsSlicesIndex(s4, genericsPerson{Name: "Ahmed", Age: 25})) // 1

	l1 := genericsList[int]{}
	l1.push(44)
	l1.push(55)
	l1.push(66)
	fmt.Println(l1.allElements()) // [44 55 66]
}
