package main

type ReduceFunc func(prev, cur int) int

func Reduce(source []int, accum int, reduceFunc ReduceFunc) int {
	if len(source) == 0 {
		return accum
	}
	for i := 0; i < len(source); i++ {
		accum = reduceFunc(accum, source[i])
	}
	return accum
}

func main() {

	sumFunc := func(prev, cur int) int {
		return cur + prev
	}

	s := []int{1, 1, 1, 2, 3, 4, 5}
	res := Reduce(s, 0, sumFunc)
	println(res)
}
