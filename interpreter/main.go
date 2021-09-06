package main

import "fmt"

func main() {
	p := &Parser{}
	str:= "1 + 2 + 3 - 4 + 5 - 6"
	p.Parse(str)
	res :=p.Result().Interpreter()
	fmt.Printf("Express:%s=%d\n",str,res)
	str2 := "1 + 2 + 3 - 4 + 5 - 6 + 8 + 12"
	p.Parse(str2)
	res2 :=p.Result().Interpreter()
	fmt.Printf("Express:%s=%d\n",str2,res2)
}
