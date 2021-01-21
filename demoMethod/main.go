package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) String() string  {
	return fmt.Sprintf("x = %v , y = %v", v.X, v.Y)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func Abs(v *Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	//testInterface()
	//testInterface2()
	//testInterfaceEmpty()
	testInterfaceAssertion()
	testMethod()
	doExerciseStringer()
}

func testMethod()  {
	v := Vertex{
		X: 3,
		Y: 4,
	}
	fmt.Println(v)
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(Abs(&v))

	Scale(&v, 10)
	fmt.Println(v.Abs())

	f := MyFloat(-1000)
	fmt.Println(f.Abs())
}

type Abser interface {
	Abs() float64
}

func testInterface()  {
	var a Abser
	f := MyFloat(-1000)

	a = f
	fmt.Println(a.Abs())

	v := Vertex{4, 5 }

	a = v
	fmt.Println(a.Abs())
	Scale(&v, 10)

	fmt.Println(a.Abs())
	fmt.Println(f.Abs())
	fmt.Println(v.Abs())
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M()  {
	fmt.Println(t.S)
}

func testInterface2()  {
	var i I = T{S : "Test interface"}
	i.M()
	describe(i)

	var i2 I
	describe(i2)
}

func describe(i I)  {
	fmt.Printf("%v, %T \n", i, i)
}

func testInterfaceEmpty()  {
	var i interface{}

	i = 100
	describeEmpty(i)

	i = T { "str"}
	describeEmpty(i)
}

func describeEmpty(i interface{})  {
	fmt.Printf("%v, %T \n", i, i)
}

func testInterfaceAssertion()  {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	var i2 interface{} = 10.0
	f, ok := i2.(float64)
	fmt.Println(f, ok)

	f = i2.(float64) // panic
	fmt.Println(f)
}

type IPAddr [4]byte

func (ip IPAddr) String() string  {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func doExerciseStringer()  {
	ip := IPAddr{127, 0, 0, 1}
	fmt.Println(ip)
}