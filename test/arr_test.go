package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestArr(t *testing.T) {

	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针

	fmt.Printf("%v", *b)

}

func TestLen(t *testing.T) {
	m := [...]int{'a': 1, 'b': 2, 'c': 3}
	m['a'] = 3
	fmt.Println("m len =", len(m))
}

func TestF5(t *testing.T) {
	fmt.Println(f5())
}

func f5() (r int) {
	defer func(r *int) {
		*r = *r + 5
	}(&r)
	return r
}

func TestI(t *testing.T) {
	// 定义好接口
	var v Worker

	v = findSomething()
	if v != nil {
		// 走的是这个分支
		fmt.Printf("v(%v) != nil\n", v)
	} else {
		fmt.Printf("v(%v) == nil\n", v)
	}
}

type Worker interface {
	Work() error
}

type Qstruct struct{}

func (q *Qstruct) Work() error {
	return nil
}

// 返回一个 nil
func findSomething() *Qstruct {
	return nil
}

func TestArr1(t *testing.T) {
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	//数组是值类型,传递过程会造成内存复制,性能不如使用slice和指针数组
	test(a)
	fmt.Println(a)
}

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

func TestP2(t *testing.T) {
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	testp(&a)
	fmt.Println(a)
}

func testp(x *[2]int) {
	fmt.Printf("x: %p\n", x)
	x[1] = 1000
}

func TestSum(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000) // 产生一个0到1000随机数
	}
	fmt.Println(b)
	sum := sum(b)
	fmt.Printf("sum=%d\n", sum)

}

func sum(a [10]int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func TestArr2(t *testing.T) {
	b := [5]int{1, 3, 5, 8, 7}
	myTest(b, 8)
}

// 求元素和，是给定的值
func myTest(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

// slice 取值
func TestSlice(t *testing.T) {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[0]
	fmt.Println("s1", s1)
	s2 := s[:]
	fmt.Println("s2", s2)
	s3 := s[2:4]
	fmt.Println("s3", s3)
	s4 := s[:8]
	fmt.Println("s4", s4)
	s5 := s[3:]
	fmt.Println("s5", s5)
}

// slice 扩容规律 以原来2倍容量进行扩容
func TestSliceAddCap(t *testing.T) {
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}

func TestSC(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array data : ", data)
	s1 := data[8:]
	s2 := data[:5]
	fmt.Printf("slice s1 : %v\n", s1)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	fmt.Println("last array data : ", data)
}

func TestChangeStr(t *testing.T) {
	str := "Hello world"
	s := []byte(str) //中文字符需要用[]rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)
}
