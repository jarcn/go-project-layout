package test

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func TestWebp(t *testing.T) {
	file, err := os.Open("test_data/images/m4_q75.webp")
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create("example/output_decode.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	img, err := webp.Decode(file, &decoder.Options{})
	if err != nil {
		log.Fatalln(err)
	}

	if err = jpeg.Encode(output, img, &jpeg.Options{Quality: 75}); err != nil {
		log.Fatalln(err)
	}
}

func TestJpgToWebp(t *testing.T) {
	file, err := os.Open("/Users/chenjia/Desktop/test/cyz.png")
	if err != nil {
		log.Fatalln(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create("/Users/chenjia/Desktop/test/cyz.webp")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	if err := webp.Encode(output, img, options); err != nil {
		log.Fatalln(err)
	}
}
