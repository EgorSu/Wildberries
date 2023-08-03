package pattern

import "fmt"

type reciever interface {
	start()
	stop()
}
type Command interface {
	execute()
}

type processor struct {
	name string
}

func (p *processor) start() {
	fmt.Printf("start processor %s\n", p.name)
}
func (p *processor) stop() {
	fmt.Println("stop processor %s\n", p.name)
}

type collector struct {
	name *string
}

func (c *collector) start() {
	fmt.Println("start collecotor %s\n", c.name)
}
func (c *collector) stop() {
	fmt.Println("stop collector %s\n", c.name)
}

type commandStart struct {
	reciever
}

func (s *commandStart) execute() {
	s.reciever.start()
}

type commandStop struct {
	reciever
}

func (s *commandStop) execute() {
	s.reciever.stop()
}
