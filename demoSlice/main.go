package main

import (
	"fmt"
	"math/rand"
	"time"
)

var myRand *rand.Rand

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	myRand = rand.New(s1)

	//lesson1Slice()
	//lesson2SliceOfSlice()
	//lesson3AppendSlice()
	//lesson4Range()
}

func lesson1Slice() {
	fmt.Println("Simple slice")

	a := make([]int, 5)
	printSlice("a", a)

	b := a[1:3]
	printSlice("b", b)

	c := b[1:4]
	printSlice("c", c)

	c[2] = 10
	b[1] = 10

	printSlice("a", a)
	printSlice("b", b)
	printSlice("c", c)
}

func printSlice(s string, x []int) {
	if x == nil {
		fmt.Println("Slice " + s + " is nil......")
	} else {
		fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
	}
}

func lesson2SliceOfSlice() {
	fmt.Println("Start lesson 2 => slice of slice")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		{"_", "_", "_"},
	}

	for i := 0; i < len(board); i++ {
		str := board[i]
		for j := 0; j < len(str); j++ {
			if myRand.Intn(100) > 50 {
				str[j] = "X"
			} else {
				str[j] = "O"
			}
		}
	}
	printSliceOfSlice("board", board)
}

func printSliceOfSlice(s string, x [][]string) {
	fmt.Println("Slice s = " + s)
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

func lesson3AppendSlice() {
	fmt.Println("Lesson 3 - Append slice")

	var s []int
	printSlice("s", s)

	s = append(s, 0)
	printSlice("s", s)

	s = append(s, 1)
	printSlice("s", s)

	s = append(s, 2, 3, 4, 5)
	printSlice("s", s)

	s1 := s[1:5]
	printSlice("s1", s1)

	fmt.Println("Change s1 ========> change s")
	s1[1] = 10
	printSlice("s", s)
	printSlice("s1", s1)

	fmt.Println("Append s1 ====> dont change s")
	s1 = append(s1, 6, 7, 8)
	printSlice("s", s)
	printSlice("s1", s1)

	fmt.Println("Then change s1 ====> dont change s")
	s1[2] = 20
	printSlice("s", s)
	printSlice("s1", s1)
}

func lesson4Range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 120}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow1 := make([]int, 10)
	for i := range pow1 {
		pow1[i] = 1 << i // == 2 ** i
	}

	for i, value := range pow1 {
		fmt.Printf("Value = %d, value2 = %d\n", value, uint(value)>>i)
	}
}

//demoSlice on tour of go
func Pic(dx, dy int) [][]uint8 {
	var myPic = make([][]uint8, dy)
	var l = len(myPic)
	for i := 0; i < l; i++ {
		myPic[i] = make([]uint8, dx)
		for j := range myPic[i] {
			myPic[i][j] = uint8(myRand.Intn(255))
		}
	}

	return myPic
}
