package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Student struct {
	ID     int    `json:"id"`     //通过指定tag实现json序列化该字段时的key
	Gender string `json:"gender"` //json序列化是默认使用字段名作为key
	Name   string `json:"name"`   //私有不能被json包访问
}

func TestStrut(t *testing.T) {
	s1 := Student{
		ID:     1,
		Gender: "女",
		Name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}
}
