package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					fmt.Println("a done")
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					fmt.Println("b done")
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func ff(a int) int {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return a
}

func Merge_Channels(f func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int) {
	
	var wg sync.WaitGroup
	res := make(chan int, 2)
	rerr := make(chan int, 2)
	
	Worker := func(inp <-chan int) {
		defer wg.Done()
		x, ok := <-inp
		if !ok {
			rerr <- 1
			res <- 0
		} else {
			rerr <- 0			
			res <- f(x)
		}
	}
	
	go func() {
		for i := 0; i < n; i++ {
			wg.Add(2)
			go Worker(in1)
			go Worker(in2)
			wg.Wait()
			k1 := <-res
			k2 := <-res
			kerr1 := <-rerr
			if kerr1 == 1 {
				break
			}
			kerr2 := <-rerr
			if kerr2 == 1 {
				break
			}
			out <- k1 + k2
		}		
		close(res)
		close(out)
	}()
}

func main() {
	a := asChan(1, 3, 5)
	b := asChan(2, 4, 6)
	c := make(chan int)
	Merge_Channels(ff, a, b, c, 5)
	// merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
