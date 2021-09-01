package main

import "fmt"

func main() {
	instanceA := GetInstance()
	instanceB := GetInstance()
	if instanceA != instanceB {
		fmt.Println("instance is not equal")
	}
	fmt.Println("instance is equal")
}
