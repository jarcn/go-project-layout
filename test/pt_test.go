package test

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
)

type JPG struct {
	Images []string `json:"images"`
}

func TestParse(t *testing.T) {

	url := "http://192.168.50.117:8866/predict/ocr_system"
	method := "POST"
	client := &http.Client{}
	base64 := jpg2Base64("t.jpg")
	jpg := JPG{Images: []string{base64}}
	v, _ := json.Marshal(jpg)
	payload := strings.NewReader(string(v))
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func jpg2Base64(jpgFile string) string {
	imgFile, err := os.Open(jpgFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	//fmt.Println("Base64 string is:", imgBase64Str)
	return imgBase64Str
}

// 指针地址、指针类型、指针取值
func TestP3(t *testing.T) {
	a := 10
	p := &a
	fmt.Printf("p的指针:%p\np的值:%v\n", p, *p)
}

func TestP4(t *testing.T) {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	fmt.Printf("p的类型%T\n", p)
	fmt.Println(reflect.TypeOf(p))
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

func TestP5(t *testing.T) {
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
}

func TestMake(t *testing.T) {
	m := make(map[string]int, 5)
	fmt.Printf("%T,%v\n", m, m)

	c := make(chan string, 1)
	c <- "测试"
	fmt.Printf("%T,%v\n", c, <-c)

	s := make([]int, 1)
	fmt.Printf("%T,%v\n", s, s)

}
