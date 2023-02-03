package test

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-tika/tika"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func TestPdfv1(t *testing.T) {
	f, err := os.Open("pdf.txt")
	defer f.Close()
	if err != nil {
		panic("文件读取异常")
	}
	r := csv.NewReader(f)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		// fmt.Println(row)
		url := row[0]
		fn := url[strings.LastIndex(url, "/")+1:]
		downPdf(url, fn)
		html := fn[:strings.LastIndex(fn, ".")+1] + "txt"
		c, _ := readPdf(fn)
		txt := trimHtml(c)
		writeFile(html, txt)
	}
}

func writeFile(path, content string) {
	f, err := os.Create(path)
	if err != nil {
		panic("file create error")
	}
	f.WriteString(content)
	defer f.Close()
}

func readPdf(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	client := tika.NewClient(nil, "http://192.168.50.117:9998")
	return client.Parse(context.TODO(), f)
}

func downPdf(url, path string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Create output file
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	// copy stream
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func TestConvertPdfToJpg(t *testing.T) {
	ConvertPdfToJpg("t.pdf", "t.jpg")
}

func ConvertPdfToJpg(pdfName, imgName string) error {
	// Setup
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// Must be *before* ReadImageFile
	// Make sure our image is high quality
	if err := mw.SetResolution(300, 300); err != nil {
		return err
	}

	// Load the image file into imagick
	if err := mw.ReadImage(pdfName); err != nil {
		return err
	}

	// Must be *after* ReadImageFile
	// Flatten image and remove alpha channel, to prevent alpha turning black in jpg
	if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_ACTIVATE); err != nil {
		return err
	}

	// Set any compression (100 = max quality)
	if err := mw.SetCompressionQuality(95); err != nil {
		return err
	}

	// Select only first page of pdf
	mw.SetIteratorIndex(0)

	// Convert into JPG
	if err := mw.SetFormat("jpg"); err != nil {
		return err
	}

	// Save File
	return mw.WriteImage(imgName)
}
