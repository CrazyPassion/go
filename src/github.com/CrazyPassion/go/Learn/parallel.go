package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched()
		fmt.Println(s)
		runtime.Gosched()
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func parallelf2() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func parallelf1() {
	runtime.GOMAXPROCS(1)
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行
	fmt.Println("end")
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func parallelf3() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciselect(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func parallelf4() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- 1
	}()
	fibonacciselect(c, quit)
}

var quit1 chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Printf("%d ", i)
	}
	quit1 <- 1
}

func paralleff5() {
	runtime.GOMAXPROCS(1)
	// 开两个goroutine跑函数loop, loop函数负责打印10个数
	go loop()
	go loop()

	//用来干什么？// 等待所有完成消息发送完毕？
	for i := 0; i < 2; i++ {
		<-quit1
	}
}

func overtime() {
	c := make(chan int)
	o := make(chan int)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-o:
				println("here")
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- 0
				c <- 0
				break
			}
		}
	}()
	// c <- 0
	o <- 1
	println("test")
	// <-o

	// <-c

}

func main() {
	parallelf1()
}
