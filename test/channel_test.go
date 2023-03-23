package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

//写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

var ic = make(chan int, 5)

func w(ic chan int, n int) {
	ic <- n
}

func TestWR(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			w(ic, rand.Intn(5))
		}
		close(ic) //这里要是不关,则chan会一直阻塞在读的那一端
	}()
	go func() {
		defer wg.Done()
		for v := range ic {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
