package pattern

type Strategy interface {
	stat([]int) int
}

type mean struct{}

func (m mean) stat(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum / len(arr)
}

type median struct{}

func (m *median) stat(arr []int) int {
	res := arr[(len(arr)-1)/2]
	return res
}

type Context struct {
	strat Strategy
}

func (c *Context) setStrategy(s Strategy) {
	c.strat = s
}
func (c *Context) stat(arr []int) int {
	return c.strat.stat(arr)
}
