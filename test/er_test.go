package test

import (
	"fmt"
	"testing"
)

type User interface {
	say(language string) error
	eat(food string) error
}

type Man struct{}
type Woman struct{}

func (m *Man) say(lang string) error {
	fmt.Println("man say language:", lang)
	return nil
}

func (m *Man) eat(food string) error {
	fmt.Println("man eat :", food)
	return nil
}

func (m Woman) say(lang string) error {
	fmt.Println("woman say language:", lang)
	return nil
}

func (m Woman) eat(food string) error {
	fmt.Println("woman eat :", food)
	return nil
}
func TestUser(t *testing.T) {
	var u User = &Man{} //指针类型接收者
	u.eat("rice")
	u.say("ch")
	var u1 User = Woman{} //值类型接收者
	u1.eat("fish")
	u1.say("en")
}

type People interface {
	Speak(string) string
}

type Student1 struct{}

func (stu *Student1) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func TestStudent1(t *testing.T) {
	var peo People = &Student1{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

func TestOdd(t *testing.T) {
	num := 1319587513358278727
	if (num & 0x1) == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}
}
