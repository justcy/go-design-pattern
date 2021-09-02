package main

import "fmt"

func main() {
	adaptee := NewAdaptee()
	targe := NewAdapter(adaptee)

	str := targe.Request()
	fmt.Println(str)
}
