package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//demoGo()
	//demoChan()
	doExercise()
}

func crawlData(note int, c chan int)  {
	for  {
		var timeDelay = myRand.Intn(MAX_DELAY - MIN_DELAY) + MIN_DELAY

		time.Sleep(time.Duration(timeDelay) * time.Millisecond)

		//fmt.Println("crawlData ", note)
		c <- note
	}

}

func showData(c chan int, quit chan int)  {
	for i := 0; i < MAX_CHAN; i++ {
		fmt.Println("Get value ", i, "data =", <-c)
	}
	quit <- 0
}

var myRand *rand.Rand

const MIN_DELAY = 1000
const MAX_DELAY = 2000

const MAX_CHAN = 100
const COUNT_CHAN_ONCE_TIME = 5

func doExercise()  {
	c := make(chan int, 100)
	quit := make(chan int)
	s1 := rand.NewSource(time.Now().UnixNano())
	myRand = rand.New(s1)

	for i := 0; i < COUNT_CHAN_ONCE_TIME; i++ {
		x := i
		go crawlData(x, c)
	}
	go showData(c, quit)

	fmt.Println("Quit - ", <- quit) //he thong doi xong moi quit
	//time.Sleep(50 * time.Second)
}

func sum(note string, s []int, c chan int)  {
	sum := 0
	for _,v := range s{
		sum += v
	}
	fmt.Println("Before = ", note)
	c <- sum //bo sum vao cho chan
	fmt.Println("Note = ", note)
}

func demoChan() {
	s := []int {1, 2, 3, 4, 5, 6, 7, 8}

	c := make(chan int, 2) //tao chan kieu int
	go sum("lan1", s[ : len(s) / 2], c)

	go sum("lan2", s[len(s)/2:], c)

	x := <-c //lay gia tri tu c ra
	fmt.Println("x = ", x)


	go sum("lan3", s, c)
	y := <-c //lay gia tri lan 2

	fmt.Println("y = ", y)
	fmt.Println( <-c)

	go sum("lan4", s, c)
	fmt.Println( <-c)
}

func demoGo()  {
	go say("world")
	say("hello")
}

func say(s string)  {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}