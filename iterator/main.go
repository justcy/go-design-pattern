package main

import "fmt"

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%#v\n", c)
	}
}

func main() {
	var aggregate Aggregate
	aggregate = NewConcreateAggregate(1, 8)

	IteratorPrint(aggregate.CreateIterator())
}

