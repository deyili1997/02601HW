package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	size, _ := strconv.Atoi(os.Args[1])
	pile, _ := strconv.Atoi(os.Args[2])
	placement := os.Args[3]
	// Set up the board
	rand.Seed(time.Now().UnixNano())
	Board := PutCoins(GenerateBoard(size), pile, placement)
	//***************************Serial Version*************************************
	start := time.Now()
	SerialResult := SerialComputing(Board)
	elapsed := time.Now().Sub(start)
	fmt.Println("Serial took", elapsed.Seconds(), "seconds.")

	// //****************************Parrall Version***********************************
	start1 := time.Now()
	ParalResult := ParallelComputing(Board, placement)
	elapsed1 := time.Now().Sub(start1)
	fmt.Println("Parallel took", elapsed1.Seconds(), "seconds.")
	//****************************Drawing 2 Pictures******************************
	fmt.Println("Now start drawing 2 pictures, please wait...")
	image1 := DrawGameBoard(SerialResult, 1)
	imageToPNG(image1, "serial.png")

	image2 := DrawGameBoard(ParalResult, 1)
	imageToPNG(image2, "parallel.png")
	fmt.Println("Finish drawing 2 pictures!")

}

//Draw Picture Function
func imageToPNG(img image.Image, name string) {
	emptyFile, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	finalImage := img
	defer emptyFile.Close()
	err = png.Encode(emptyFile, finalImage)
	if err != nil {
		fmt.Println(err)
	}
}
