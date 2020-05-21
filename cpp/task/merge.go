package main

import (
    "fmt"
    "sync"
)

func getInputChan() <-chan int {
	// Создание в динамической памяти канала на 100 элементов int и получаем на него указатель в input.
	input := make(chan int, 100)

	// Массив данных для канала.
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// В отдельной горутине выполняем отправку данных массива в канал.
	go func() {
		for num := range numbers {
			input <- num
		}
		close(input)
	}()

	// Функция возвращает указатель на канал доступный только для считывания данных.
	return input
}

func getSquareChan(input <-chan int) <-chan int {

	output := make(chan int, 100)

	// В горутине считываем из канала число, возводим его в квадрат и отправляем в канал output.
	go func() {
		for num := range input {
				output <- num * num
		}

		close(output)
	}()

	return output
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
		
	// step 1 Создание канала.
  chanInputNums := getInputChan()

  // step 2 Разделяем каналы для возведения в квадрат.
  chanOptSqr1 := getSquareChan(chanInputNums)
  chanOptSqr2 := getSquareChan(chanInputNums)

  // step 3 Объединение значений двух каналов в один.
  chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)

  // step 4 Сумма значений канала.
  sqrSum := 0
  for num := range chanMergedSqr {
  	sqrSum += num
  }

  // step 5 Результат.
  fmt.Println("Сумма квадратов чисел в от 0 до 9 равна: ", sqrSum)
}
