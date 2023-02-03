package test

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/h2non/bimg"
// )

// func pdf2img() {

// 	buffer, err := bimg.Read("1614761708234776576.pdf")
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}

// 	newImage, err := bimg.NewImage(buffer).Convert(bimg.JPEG)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}

// 	if bimg.NewImage(newImage).Type() == "jpeg" {
// 		fmt.Fprintln(os.Stderr, "The image was converted into jpeg")
// 	}

// 	bimg.Write("test.jpg", newImage)

// }

// func TestPdf1(t *testing.T) {
// 	pdf2img()
// }
