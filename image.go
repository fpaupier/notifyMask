package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
)

// bytesToJpeg decodes imgByte into a jpeg image and saves it to the imgPath passed.
func bytesToJpeg(imgByte []byte, imgPath string) {
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}
	out, _ := os.Create(imgPath)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)
	if err != nil {
		log.Println(err)
	}
}
