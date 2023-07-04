package main

import (
	"fmt"
)

type adapter struct {
	slowAdapter *slow
}

func (adap *adapter) todo() {
	adap.slowAdapter.slowTodo()
}

type action interface {
	todo()
}

type Man struct{}

func (m *Man) startAction(a action) {
	a.todo()
}

// реализует интерфейс action
type fast struct{}

func (r *fast) todo() {
	fmt.Println("fast")
}

// будет адаптирован под интерфейс action
type slow struct{}

func (w *slow) slowTodo() {
	fmt.Println("slow")
}

func main() {

	man := &Man{}
	fastAction := &fast{}
	slowAction := &slow{}

	man.startAction(fastAction)
	changeSpeedAction := &adapter{
		slowAdapter: slowAction,
	}
	man.startAction(changeSpeedAction)
}
