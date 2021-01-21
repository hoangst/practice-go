package main

import "fmt"

func main2() {
	demoSelect()
}

func fibonacci(c, quit chan int)  {
	x, y := 0, 1
	for {
		select {
			case c <- x: //cu push x vao chan
				x, y = y, x + y
			case <-quit://khi ma get duoc quit => da get du data => return
				fmt.Println("quit")
				return
		}
	}
}

func showData(c, quit chan int)  {
	for i := 0; i < 15; i++ {
		fmt.Println(<-c) //=> lay data trong channel
	}
	quit <- 0 //khi dang xong
}

func demoSelect()  {
	c := make(chan  int)
	quit := make(chan int)
	go showData(c, quit)
	fibonacci(c, quit)
}