package main

func main() {
	receiver := NewRecirver()
	A := NewConcreateCommandA(receiver)
	invoker := NewInvoker(A)
	invoker.Call()

	B := NewConcreateCommandB(receiver)
	invoker2 := NewInvoker(B)
	invoker2.Call()
}
