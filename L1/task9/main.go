package main

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

// типы, которым соответствует этот интерфейс, будут ограничены численными типами
type Number interface {
	constraints.Integer | constraints.Float
}

// используем дженерики для создания универсальной функции для всех численный типов
func source[T Number](num []T, out chan<- T) {
	for _, val := range num {
		out <- val
	}
	close(out) //закрытие этого канала позволит завершить for range в func operation
}

// при возведении в квадрат мы можем выйти за пределы диапазона допустимых значений типа,
// тогда вернётся некорректное значение
// поэтому добавим возможность использовать разные типы для входного и выходного канала
// возможно, при возникновении такого события, лучше возваращать ошибку
func operation[T, K Number](out chan<- T, in <-chan K) {
	for val := range in {
		out <- T(val) * T(val)
		reflect.TypeOf(val)
	}
	close(out) //закрытие этого канала позволит завершить for range в main
}
func main() {
	//используем универсальную функцию для типов int и float32
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	number := []int{1, 2, 3, 4, 5}
	go source(number, ch1)
	go operation(ch2, ch1)
	for result := range ch2 {
		fmt.Println(result)
	}

	ch1f := make(chan float32, 1)
	ch2f := make(chan float32, 1)
	numberf := []float32{1.1, 2.5, 3.0, 4.4, 5.9}
	go source(numberf, ch1f)
	go operation(ch2f, ch1f)
	for result := range ch2f {
		fmt.Println(result)
	}
	//попробуем использовать uint8
	//чтобы получить корректный результат возведения в квадрат числа 200
	//для возврата значений будем использовать chan uint16
	ch1u := make(chan uint8, 1)
	ch2u := make(chan uint16, 1)
	numberu := []uint8{0, 1, 2, 3, 200}
	go source(numberu, ch1u)
	go operation(ch2u, ch1u)
	for result := range ch2u {
		fmt.Println(result)
	}
}
