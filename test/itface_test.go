package test

import (
	"fmt"
	"testing"
)

// 接口强验证(go 圈中比较常用的做法)
// 这是用来限制接口中的方法必须都要被实现
var _ Shape = (*Square)(nil)

type Shape interface {
	Sides() int
	Area() int
}
type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}

func (s *Square) Area() int {
	return 4
}

func TestObj(t *testing.T) {
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
}
