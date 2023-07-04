package main

import "fmt"

type Human struct {
	firstName, lastName string
	todo                func() string
	age                 int
}

// создадим 2 метода для структуры Human
func (h *Human) describe() string {
	return fmt.Sprintf("Human %v %v age %d", h.firstName, h.lastName, h.age)
}
func (h *Human) changeAge(years int) {
	h.age += years
}

type Action struct {
	Human  //встраиваем структуру Human в структуру Action
	action string
}

func main() {
	// создаём экземпляр стуктуры Action и передём ей значения полей
	a := Action{Human: Human{
		firstName: "First",
		lastName:  "Last",
		todo:      func() string { return "todo" },
		age:       99},
		action: "go"}
	//вызываем методы вложенной стурктуры
	a.changeAge(3)
	fmt.Println(a.describe())
}
