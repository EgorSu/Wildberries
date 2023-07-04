package main

import (
	"errors"
	"fmt"
)

// изменяться будет исходный slice
// чтобы не изменять исходый слайс нужно использовать
// temp := make([]T, len(arr))
// copy(temp, arr)
// return temp[:id] or append(temp[:id], temp[id+1:]...)
func deleteId[T any](arr []T, id int) ([]T, error) {

	switch {
	case id == len(arr)-1:
		return arr[:id], nil
	case id >= 0 && id < len(arr)-1:
		return append(arr[:id], arr[id+1:]...), nil
	}
	return make([]T, 0), errors.New(fmt.Sprintf("id %d out of range slice with len %d", id, len(arr)))
}

func main() {

	for i := 0; i < 5; i++ {
		arrS := []string{"ad", "asd", "ljklhj"}
		result, err := deleteId(arrS, i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
