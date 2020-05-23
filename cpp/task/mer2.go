package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func getInputChan(numbers []int, test chan int) <-chan int {
	var wg sync.WaitGroup
	
	input := make(chan int, 100)

	wg.Add(len(numbers))

	for _, num := range numbers {
		go func(k int) {
			rand.Seed(time.Now().UTC().UnixNano())
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			input <-k
			test <-k
			wg.Done()
		}(num)
	}

	go func() {
		wg.Wait()
		close(input)
		close(test)
	}()	

	// Функция возвращает указатель на канал доступный только для считывания данных.
	return input
}

func ff(n int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return n
}

func Merge2(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
// Параллельное чтение по одному числу из двух каналов x1 и x2
// и отправка в выходной канал суммы f(x1)+f(x2) в функции алгоритм повторяется n раз.

	go func() {
		var wg sync.WaitGroup		
		// Создаём слайс из n каналов.
		var ch1 []chan int
		var ch2 []chan int

		for i:=0; i<n; i++ {
			// Добавляем канал в слайс.
			 ch1 = append(ch1, make(chan int))
			 ch2 = append(ch2, make(chan int))
		}
		wg.Add(n)

		go func(){
			wg.Wait()
			close(out)	
		}()

		for i:=0; i<n; i++ {
			go func(p_k int){
				out <- f(<-ch1[p_k]) + f(<-ch2[p_k])
				wg.Done()
			}(i)
		}

		for i:=0; i<n; i++ {
			go func(p_k int) {
				ch1[p_k]<- <-in1
				ch2[p_k]<- <-in2	
			}(i)
		}
	}()
}

func Merge2PS(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
// Последовательное чтение.
	go func() {
		var wg sync.WaitGroup
		
		wg.Add(n)

		go func(){
			wg.Wait();
			close(out)
		}()
		
		for i := 0; i < n; i++ {
			x1:=<-in1
			x2:=<-in2
			out <- f(x1) + f(x2)
			wg.Done()
		}
	}()

}

func main() {
	d1 := make(chan int, 100)
	d2 := make(chan int, 100)
	c3 := make(chan int)

	c1 := getInputChan([]int{11,12,13,14,15,16,17,18,19}, d1)
	c2 := getInputChan([]int{21,22,23,24,25,26,27,28,29}, d2)
	
	Merge2(ff, c1, c2, c3, 9)
	fmt.Println("=== Сумма чисел из двух каналов")	
	for c := range c3 {
		fmt.Print(c, ",")
	}	
	fmt.Println("")
	fmt.Println("=== Канал 1")
	for num := range d1 {
		fmt.Print(num, ",")
	}		
	fmt.Println("")
	fmt.Println("=== Канал 2")
	for num := range d2 {
		fmt.Print(num, ",")
	}				

}
