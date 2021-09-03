package main

func main() {
	m := NewCommonImplementor(NewImplementorA())
	m.Operation()

	m = NewCommonImplementor(NewImplementorB())
	m.Operation()

	m2 := NewCommonImplementor2(NewImplementorA())
	m2.Operation()

	m2 = NewCommonImplementor2(NewImplementorB())
	m2.Operation()
}
