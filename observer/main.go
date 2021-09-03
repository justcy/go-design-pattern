package main

func main() {
	subject := NewSubject()
	observer1 := NewConcreteObserver("A")
	observer2 := NewConcreteObserver("B")
	observer3 := NewConcreteObserver("C")

	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.Attach(observer3)

	subject.SetState(1)
	subject.Notify()

	subject.Detach(observer1)
	subject.SetState(2)
	subject.Notify()

}

