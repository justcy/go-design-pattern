package main

func main() {
	factory := NewFlyweightFactory()
	fw:=factory.Get("one")
	fw.Operation()

	fw =factory.Get("two")
	fw.Operation()


	fw =factory.Get("three")
	fw.Operation()

	fw =factory.Get("one")
	fw.Operation()
}
