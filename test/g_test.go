package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

//写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

var Mem = make(chan int, 5)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- rand.Intn(100)
	}
	defer close(ch)
}

func print(ch chan int) {
	for {
		n, ok := <-ch
		if ok {
			fmt.Printf("%d\n", n)
		}
	}

}

func TestMain(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go write(Mem)
	wg.Done()
	go print(Mem)
	wg.Done()
	wg.Wait()
}

func TestSyncMap(t *testing.T) {
	var m sync.Map
	m.Store("address", map[string]string{"province": "江苏", "city": "南京"})
	v, _ := m.Load("address")
	fmt.Println(v.(map[string]string)["province"])
}

// string 底层是字节数组,string 类型可以通过[]byte("string")进行强转
// rune 是 byte 的别名
func TestStringChange(t *testing.T) {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}
