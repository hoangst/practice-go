package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main()  {
	//funcValue()
	//closuresFunc()
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	one := 0
	second := 0
	return func() int {
		if count <= 2 {
			if count == 0 {
				count = 1
				return 0
			}
			if count == 1 {
				count = 2
				return 1
			}
			count = 3
			one = 1
			second = 1
			return 1
		}
		value := second + one
		one = second
		second = value
		return value
	}
}

func closuresFunc()  {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-i * 2))
	}
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func funcValue()  {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(6, 8))

	myFunc := func(x, y float64) float64 {
		return x*x + y*y*2
	}

	fmt.Println(compute(myFunc))

	fmt.Println(compute(hypot))

	fmt.Println(compute(math.Atan2))
}