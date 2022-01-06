package main

import (
	"image"
)

func DrawGameBoard(gb GameBoard, cellWidth int) image.Image {
	height := len(gb.board) * cellWidth
	width := len(gb.board[0]) * cellWidth
	c := CreateNewPalettedCanvas(width, height, nil)

	// declare colors
	// darkGray := MakeColor(50, 50, 50)
	black := MakeColor(0, 0, 0)
	grey1 := MakeColor(85, 85, 85)
	grey2 := MakeColor(170, 170, 170)
	// blue := MakeColor(0, 0, 255)
	// red := MakeColor(255, 0, 0)
	// green := MakeColor(0, 255, 0)
	// yellow := MakeColor(255, 255, 0)
	// magenta := MakeColor(255, 0, 255)
	white := MakeColor(255, 255, 255)
	// cyan := MakeColor(0, 255, 255)

	/*
		//set the entire board as black
		c.SetFillColor(gray)
		c.ClearRect(0, 0, height, width)
		c.Clear()
	*/

	/*
		// draw the grid lines in white
		c.SetStrokeColor(white)
		DrawGridLines(c, cellWidth)
	*/

	// fill in colored squares
	for i := range gb.board {
		for j := range gb.board[i] {
			if gb.board[i][j].num == 0 {
				c.SetFillColor(black)
			} else if gb.board[i][j].num == 1 {
				c.SetFillColor(grey1)
			} else if gb.board[i][j].num == 2 {
				c.SetFillColor(grey2)
			} else if gb.board[i][j].num == 3 {
				c.SetFillColor(white)
			} else {
				panic("Error: Out of range value " + string(gb.board[i][j].num) + " in board when drawing board.")
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}

	return GetImage(c)
}

func DrawGridLines(pic Canvas, cellWidth int) {
	w, h := pic.Width(), pic.Height()
	// first, draw vertical lines
	for i := 1; i < w/cellWidth; i++ {
		y := i * cellWidth
		pic.MoveTo(0.0, float64(y))
		pic.LineTo(float64(w), float64(y))
	}
	// next, draw horizontal lines
	for j := 1; j < h/cellWidth; j++ {
		x := j * cellWidth
		pic.MoveTo(float64(x), 0.0)
		pic.LineTo(float64(x), float64(h))
	}
	pic.Stroke()
}
