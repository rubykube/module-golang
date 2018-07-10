package stack

//package main

//import "fmt"

type Stack struct {
	i  int
	st []int
}

func New() *Stack {
	return &Stack{}
}

func (c *Stack) Push(numb int) {
	c.i++
	c.st = append(c.st, numb)
}

func (c *Stack) Pop() int {
	c.i--
	res := c.st[c.i]
	return res
}

/*func main() {
	s := make(stack, 0)
	s = c.Push(1)
	s = c.Push(2)
	s = c.Push(3)

	p := c.Pop()
	fmt.Println(p)
}*/
