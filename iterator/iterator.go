package main

type Iterator interface {
	First()
	Next() interface{}
	IsDone() bool
}
type Aggregate interface {
	CreateIterator() Iterator
}
type ConcreteIterator struct {
	numbers *ConcreateAggregate
	next    int
}

func (c *ConcreteIterator) First() {
	c.next = c.numbers.start
}

func (c *ConcreteIterator) Next() interface{}{
	if !c.IsDone() {
		next :=c.next
		c.next++
		return next
	}
	return nil
}

func (c *ConcreteIterator) IsDone() bool {
	return c.next > c.numbers.end
}

type ConcreateAggregate struct {
	start, end int
}

func (c *ConcreateAggregate) CreateIterator() Iterator {
	return &ConcreteIterator{
		numbers:c,
		next:c.start,
	}
}

func NewConcreateAggregate(start, end int) *ConcreateAggregate {
	return &ConcreateAggregate{
		start: start,
		end:   end,
	}
}
