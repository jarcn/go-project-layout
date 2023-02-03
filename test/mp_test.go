package test

import (
	"fmt"
	"strings"
	"testing"
)

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, v := range arr {
		newArray = append(newArray, fn(v))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, v := range arr {
		newArray = append(newArray, fn(v))
	}
	return newArray
}

// 函数也可以作为参数进行传递
func TestMap(t *testing.T) {
	var list = []string{"jia", "chen", "MegaEase"}
	x := MapStrToStr(list, func(s string) string { return strings.ToUpper(s) })
	t.Logf("%v\n", x)

	y := MapStrToInt(list, func(s string) int { return len(s) })
	t.Logf("%v\n", y)
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, v := range arr {
		sum += fn(v)
	}
	return sum
}

func TestReduce(t *testing.T) {
	var list = []string{"jia", "chen", "MegaEase"}
	x := Reduce(list, func(s string) int { return len(s) })
	t.Logf("%v\n", x)
}

func Filter(arr []int, fn func(n int) bool) []int {
	var newArray = []int{}
	for _, v := range arr {
		if fn(v) {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func TestFilter(t *testing.T) {
	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(intset, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("%v\n", out)

	out = Filter(intset, func(n int) bool {
		return n > 5
	})
	fmt.Printf("%v\n", out)
}

func TestMapPut(t *testing.T) {
	var mp = map[string]string{}
	mp["key"] = "value"
	fmt.Printf("%v", mp)
}
