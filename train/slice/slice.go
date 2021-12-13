package main

import (
	"fmt"
)

func main() {
	a := *new([]int)
	fmt.Printf("a %+v cap %d len %d\n", a, cap(a), len(a))

	b := make([]int, 0, 5)
	fmt.Printf("%+v cap %d len %d\n", b, cap(b), len(b))

	// increase slice capacity
	t := make([]int, len(b), 2*cap(b))
	copy(t, b)
	fmt.Printf("%+v cap %d len %d\n", t, cap(t), len(t))

	//append
	n1 := append(t, 1)
	n2 := append(t, 1)
	fmt.Printf("address t= %p, address n1=%p  n2=%p\n", &t, &n1, &n2)
	//	fmt.Printf("address t= %d, address n1=%d  n2=%d\n",uintptr(unsafe.Pointer(&t)),uintptr(unsafe.Pointer(&n1)),uintptr(unsafe.Pointer(&n2)))
	//	fmt.Println(uintptr(unsafe.Pointer(&n1)) - uintptr(unsafe.Pointer(&t)))
	//	fmt.Println(uintptr(unsafe.Pointer(&n2)) - uintptr(unsafe.Pointer(&n1)))

	//delete
	q := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	q = append(q[:1], q[5:]...)
	fmt.Printf("%+v cap %d len %d\n", q, cap(q), len(q))

	//
	q = q[0:cap(q):cap(q)]
	fmt.Printf("%+v cap %d len %d\n", q, cap(q), len(q))

}
