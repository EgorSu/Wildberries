package main

import (
	"fmt"
	"math/rand"
)

// ограничение comparable необходимо т.к. ключи map должны быть сравнимы
// значения заполним struct{} т.к. они имеют нулевой размер
type Set[T comparable] map[T]struct{}

// переопределим метод string
// хочется чтобы распечатывались только ключи map
func (s Set[T]) String() string {
	res := "["
	for k := range s {
		res += fmt.Sprintf(" %v ", k)
	}
	res += "]"
	return res
}

// дженерики для универсальности
func intersect[T comparable](a Set[T], b Set[T]) Set[T] {
	res := make(Set[T])
	//добавим в пересечение элементы, которые наши в обеих set
	for k := range a {
		if _, ok := b[k]; ok {
			res[k] = struct{}{}
		}
	}
	return res
}

// функция для заполнения множеств int, т.к. множества универсальные можно заполняться и другими сравнимыми значениями
func fillRandInt(set Set[int], size int) {
	for i := 0; i < size; {
		val := rand.Intn(2 * size)
		if _, ok := set[val]; !ok { //заполним без повторений
			set[val] = struct{}{}
			i++
		}
	}
}

func main() {

	size1 := rand.Intn(20)
	size2 := rand.Intn(20)
	setA := make(Set[int])
	setB := make(Set[int])
	fmt.Println(size1, size2)
	fillRandInt(setA, size1)
	fillRandInt(setB, size2)

	fmt.Printf("set a: %s\n", setA)
	fmt.Printf("set a: %s\n", setB)

	result := intersect(setA, setB)
	fmt.Printf("intersection: %s\n", result)
}
