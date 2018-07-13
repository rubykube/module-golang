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
	c.st = c.st[:c.i]
	return res
}

/*func main() {
	var c *Stack = New()
	//	s := make(stack, 0)
	c.Push(1)
	c.Push(2)
	c.Push(3)

	p := c.Pop()
	fmt.Println(p)
	c.Push(4)
	d := c.Pop()
	fmt.Println(d)
}*/
