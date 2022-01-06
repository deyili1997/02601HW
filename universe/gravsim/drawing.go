package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func CreateAnimationFrames(ulist []Universe, w, frameRate int) []image.Image {
	images := make([]image.Image, 0)
	for i, u := range ulist {
		if i%frameRate == 0 {
			images = append(images, u.DrawToImage(w))
		}
	}
	return images
}

func (u *Universe) DrawToImage(w int) image.Image {
	c := CreateNewPalettedCanvas(w, w, nil)

	c.SetFillColor(MakeColor(0, 0, 0))
	c.ClearRect(0, 0, w, w)
	c.Fill()

	for _, b := range u.bodies {
		b.DrawToCanvas(c, u.width, w)
	}
	return GetImage(c)
}

func (b *Body) DrawToCanvas(c Canvas, univSize float64, canvasSize int) {
	bx := (b.pos.x / univSize) * float64(canvasSize)
	by := (b.pos.y / univSize) * float64(canvasSize)
	r := b.scale * (b.radius / univSize) * float64(canvasSize)
	c.SetFillColor(MakeColor(uint8(b.red), uint8(b.green), uint8(b.blue)))
	c.Circle(bx, by, r)
	c.Fill()
}

func WriteImageAsPNG(i image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	b := bufio.NewWriter(f)
	err = png.Encode(f, i)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filename)
}
