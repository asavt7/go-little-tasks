package main

import "fmt"

type A interface {
	Do()
}

type S1 struct {
	a string
}

func (s S1) Do() {
	panic("implement me")
}

type S2 struct {
	a int
}

func (s S2) Do() {
	panic("implement me")
}

//Interface values are comparable. Two interface values are equal if they have identical dynamic types and equal dynamic values or if both have value nil.
func do(a1, a2 A) bool {
	return a1 == a2
}

func main() {

	fmt.Println(do(S1{}, S2{})) //false
	fmt.Println(do(S1{}, S1{})) //true

}
