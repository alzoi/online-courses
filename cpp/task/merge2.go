// Объединение двух каналов в один.

// Вывод:
// Два канала последовательно: 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 
// Два объединённых канала:    1, 2, 2, 1, 2, 1, 1, 2, 1, 2, 1, 2, 2, 1, 1, 2, 1, 2, 

package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

//func getInputChan(numbers []int) <-chan int {
func getInputChan(numbers ...int) <-chan int {

	input := make(chan int, 100)

	// В отдельной горутине выполняем отправку данных массива в канал.
	go func() {
		//for _, num := range numbers {
		for num := range numbers {
			input <- num
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		}
		close(input)
	}()

	// Функция возвращает указатель на канал доступный только для считывания данных.
	return input
}

func merge(outputsChan ...<-chan int) <-chan int {
	// Объединение каналов.

	// Создаём группу для синхронизации горутин.
	var wg sync.WaitGroup

	// Канал для объединения данных каналов.
	merged := make(chan int, 100)

	// Указываем количество элементов в группе синхронизации.
	wg.Add(len(outputsChan))

	// Анонимная функция для чтения данных из канала.
	p_output := func(sc <-chan int) {
		for sqr := range sc {
				merged <- sqr
		}
		// Отметка о завершении работы горутины из группы синхронизации.
		wg.Done()
	}

	// Итерация по списку входных каналов.
	for _, optChan := range outputsChan {
		// Запуск в горутине чтения из канала.
		go p_output(optChan)
	}

	// Ожидание завершения всех горутин из группы.
	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	n1 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	n2 := []int{2, 2, 2, 2, 2, 2, 2, 2, 2}

	chan1 := getInputChan(n1)	
	fmt.Print("Два канала последовательно: ")
	for num := range chan1 {
		fmt.Print(num, ", ")
	}
	chan2 := getInputChan(n2)	
	for num := range chan2 {
		fmt.Print(num, ", ")
	}

	fmt.Println()
	fmt.Print("Два объединённых канала:    ")
	chan1  = getInputChan(n1)	
	chan2  = getInputChan(n2)	
	chan3 := merge(chan1, chan2)

	for num := range chan3 {
		fmt.Print(num, ", ")
	}

}
