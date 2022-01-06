package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func CreateAnimationFrames(flist []*frame, canvasSize, frameRate, plasmaSize, bacLifespan int) []image.Image {
	images := make([]image.Image, 0)
	for i, u := range flist {
		if i%frameRate == 0 {
			images = append(images, u.DrawToImage(plasmaSize, bacLifespan, canvasSize))
		}
	}
	return images
}

func (f *frame) DrawToImage(plasmaSize, bacLifespan, canvasSize int) image.Image {
	c := CreateNewPalettedCanvas(canvasSize, canvasSize, nil)

	c.SetFillColor(MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasSize, canvasSize)
	c.Fill()

	for _, p := range f.pG.group {
		p.DrawToCanvas(c, plasmaSize, canvasSize)
	}
	for _, immun := range f.iG.group {
		immun.DrawToCanvas(c, plasmaSize, canvasSize)
	}
	for _, b := range f.bG.group {
		if ((bacLifespan / 10 * 9) < b.lifeSpan) && (b.lifeSpan <= bacLifespan) {
			b.green = 230
		} else if ((bacLifespan / 10 * 8) < b.lifeSpan) && (b.lifeSpan <= (bacLifespan / 10 * 9)) {
			b.green = 205
		} else if ((bacLifespan / 10 * 7) < b.lifeSpan) && (b.lifeSpan <= (bacLifespan / 10 * 8)) {
			b.green = 180
		} else if ((bacLifespan / 10 * 6) < b.lifeSpan) && (b.lifeSpan <= (bacLifespan / 10 * 7)) {
			b.green = 155
		} else if ((bacLifespan / 10 * 5) < b.lifeSpan) && (b.lifeSpan <= (bacLifespan / 10 * 6)) {
			b.green = 130
		} else if ((bacLifespan / 10 * 4) < b.lifeSpan) && (b.lifeSpan <= (bacLifespan / 10 * 5)) {
			b.green = 105
		} else if ((bacLifespan / 10 * 3) < b.lifeSpan) && (b.lifeSpan <= (bacLifespan / 10 * 4)) {
			b.green = 80
		} else if ((bacLifespan / 10 * 2) < b.lifeSpan) && (b.lifeSpan <= (bacLifespan / 10 * 3)) {
			b.green = 55
		} else if ((bacLifespan / 10 * 1) < b.lifeSpan) && (b.lifeSpan <= (bacLifespan / 10 * 2)) {
			b.green = 30
		} else if b.lifeSpan <= (bacLifespan / 10 * 1) {
			b.green = 5
		}
		b.DrawToCanvas(c, plasmaSize, canvasSize)
	}
	for _, ib := range f.infectedG.group {
		ib.red = 255
		ib.green = 153
		ib.blue = 51
		ib.DrawToCanvas(c, plasmaSize, canvasSize)
	}

	return GetImage(c)
}

func (p *phage) DrawToCanvas(c Canvas, plasmaSize, canvasSize int) {
	bx := (float64(p.position.x) / float64(plasmaSize)) * float64(canvasSize)
	by := (float64(p.position.y) / float64(plasmaSize)) * float64(canvasSize)
	r := (float64(p.scale) * (float64(p.radius) / float64(plasmaSize)) * float64(canvasSize))
	c.SetFillColor(MakeColor(uint8(p.red), uint8(p.green), uint8(p.blue)))
	c.Circle(bx, by, r)
	c.Fill()
}

func (immu *immuneCell) DrawToCanvas(c Canvas, plasmaSize, canvasSize int) {
	bx := (float64(immu.position.x) / float64(plasmaSize)) * float64(canvasSize)
	by := (float64(immu.position.y) / float64(plasmaSize)) * float64(canvasSize)
	r := (float64(immu.scale) * (float64(immu.radius) / float64(plasmaSize)) * float64(canvasSize))
	c.SetFillColor(MakeColor(uint8(immu.red), uint8(immu.green), uint8(immu.blue)))
	c.Circle(bx, by, r)
	c.Fill()
}

func (b *bacteria) DrawToCanvas(c Canvas, plasmaSize, canvasSize int) {
	bx := (float64(b.position.x) / float64(plasmaSize)) * float64(canvasSize)
	by := (float64(b.position.y) / float64(plasmaSize)) * float64(canvasSize)
	r := (float64(b.scale) * (float64(b.radius) / float64(plasmaSize)) * float64(canvasSize))
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
