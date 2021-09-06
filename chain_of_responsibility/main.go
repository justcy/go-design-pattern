package main

func main() {
	h1 := NewConcreteHandler1()
	h2 := NewConcreteHandler2()
	h3 := NewConcreteHandler3()

	h1.SetSuccessor(h2)
	h2.SetSuccessor(h3)

	test := []int{2,3,4,6,16,23,24,25,39}
	for _, value := range test {
		h1.HandleRequest(value)
	}
}
