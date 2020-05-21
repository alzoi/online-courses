// Объединение каналов в один.

package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func getInputChan(numbers []int) <-chan int {

	input := make(chan int, 100)

	// В отдельной горутине выполняем отправку данных массива в канал.
	go func() {
		for _, num := range numbers {
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
	chan1 := getInputChan([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25})	
	chan2 := getInputChan([]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24})	

  chan3 := merge(chan1, chan2)

  for num := range chan3 {
  	fmt.Println(num)
  }

}
