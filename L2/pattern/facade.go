package pattern

type Facade struct {
	w1 *work1
	w2 *work2
}

func (f *Facade) Todo() string {
	strWork1 := f.w1.todoWork1()
	strWork2 := f.w2.todoWork2()
	return strWork1 + "\r" + strWork2
}

type work1 struct {
}

func (w *work1) todoWork1() string {
	return "do work 1"
}

type work2 struct {
}

func (w *work2) todoWork2() string {
	return "do work 2"
}
