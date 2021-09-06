package main

func main() {
	o1 := NewOriginator()
	o1.Status()

	process := o1.CreateMemento()
	o1.Change(100)
	o1.Status()

	o1.RestoreMemento(process)
	o1.Status()
}
