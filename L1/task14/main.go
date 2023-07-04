package main

import (
	"fmt"
)

func typeDetected(i interface{}) string {
	return fmt.Sprintf("%T", i)
}
func main() {
	var myI []interface{} = []interface{}{
		456, "abc", true, make(chan int)}
	for _, e := range myI {
		fmt.Println(typeDetected(e))
	}
}
