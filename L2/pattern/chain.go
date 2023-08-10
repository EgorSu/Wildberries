package pattern

import "fmt"

type chain interface {
	setNextAction(chain)
	Do(subject)
}
type action1 struct {
	next chain
}

func (a *action1) setNextAction(next chain) {
	a.next = next
}
func (a *action1) Do(sub subject) {
	if !sub.isAction1 {
		fmt.Println("do something in action 1")
	}
	sub.isAction1 = true
	a.next.Do(sub)
}

type action2 struct {
	next chain
}

func (a *action2) setNextAction(next chain) {
	a.next = next
}
func (a *action2) Do(sub subject) {
	if !sub.isAction2 {
		fmt.Println("do something in action 2")
	}
	sub.isAction2 = true
	a.next.Do(sub)
}

type action3 struct {
	next chain
}

func (a *action3) setNextAction(next chain) {
	a.next = next
}
func (a *action3) Do(sub subject) {
	if !sub.isAction3 {
		fmt.Println("do something in action 3")
	}
}

type subject struct {
	isAction1 bool
	isAction2 bool
	isAction3 bool
}
