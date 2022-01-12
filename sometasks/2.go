package main

import "fmt"

func change(a int, b *int) {
	a++
	*b += 5
}

type Person struct {
	Name string
}

func changePersonP(p *Person) {
	p = &Person{Name: "BBB"}
}

func changePersonPP(p **Person) {
	*p = &Person{Name: "BBB"}
}

func changePersonRet(p *Person) *Person {
	return &Person{Name: "RET"}
}

func main() {
	var x = 1
	var p *int = &x
	change(x, p)
	fmt.Printf("x= %d ; *p= %d", x, *p)

	println("---------------------")

	ps := &Person{Name: "AAA"}
	changePersonP(ps) // no changes
	fmt.Println(ps)

	changePersonPP(&ps)
	fmt.Println(ps) // changed!

	ps = changePersonRet(ps)
	fmt.Println(ps) // changed!

}
