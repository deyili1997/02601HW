package main

import (
	"fmt"
	"gifhelper"
	"os"
	"strconv"
	"image"
	"image/png"
	"log"
)

// command-line parameters are stored in an array of strings called os.Args
//its length is equal to #parameters +1
//os.Args[0] is name of program
//os.Args[1] is first parameter given
//os.Args[2] is second parameter given ...
//...
//os.Args[len(os.Args)-1] is final parameter given

func main() {
	fmt.Println("Prison strategy game!")

	fileName := os.Args[1]     //read what template

	b, err1 := strconv.ParseFloat(os.Args[2], 64)//the reward for defector
	if err1 != nil  {
		panic("There is something wrong with 'b' string convertion!")
	}

	steps, err2 := strconv.Atoi(os.Args[3])// how many rounds to run the simulation
	if err2 != nil  {
		panic("There is something wrong with 'steps' string convertion!")
	}

	fmt.Println("Parameters read in successfully!")
	boardGens := RunSimulation(fileName, b, steps)

	fmt.Println("Automaton played. Now, drawing images.")

	// we need a slice of image objects
	var cellWidth int = 5
	imglist := DrawGameBoards(boardGens, cellWidth)
	fmt.Println("Boards drawn to images! Now, convert to animated GIF.")
	imageToPNG(imglist)
	// convert images to a GIF
	outputFile := "Homework0"   // where to draw the final animated GIF of boards
	gifhelper.ImagesToGIF(imglist, outputFile)

	fmt.Println("Success! GIF produced.")


}

func imageToPNG(imglist []image.Image){
	emptyFile, err := os.Create("Prisoners.png")
	if err != nil {
	  log.Fatal(err)
	}
	finalImage:=imglist[len(imglist)-1]
	defer emptyFile.Close()
	err = png.Encode(emptyFile,finalImage)
	if err!=nil{
		fmt.Println(err)
	}
}
