package main

import (
	"fmt"
	"gifhelper"
	"math"
	"os"
)

func main() {
	// the following sample parameters may be helpful for the "collide" command
	// all units are in SI (meters, kg, etc.)
	// but feel free to change the positions of the galaxies.
	command := os.Args[1]
	if command == "jupiter"{
		Jupyter := CreateJupiterSystem()
		numGens := 500000
		time := 10.0
		theta := 0.5
		timePoints := BarnesHut(Jupyter, numGens, time, theta)
		canvasWidth := 1000
		frequency := 1000
		scalingFactor := 30.0 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
		imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "Jupiter")
		fmt.Println("GIF drawn.")
	}

	if command == "galaxy" {
		g0 := InitializeGalaxy(500, 4e21, 5.0e22, 5.0e22)
		width := 1.0e23
		galaxies := []Galaxy{g0}
		initialUniverse := InitializeUniverse(galaxies, width)
		numGens := 50000
		time := 2e15
		theta := 0.5
		timePoints := BarnesHut(initialUniverse, numGens, time, theta)
		canvasWidth := 1000
		frequency := 100
		scalingFactor := 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
		imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "Galaxy")
		fmt.Println("GIF drawn.")
	}
	if command == "collision" {
		g0 := InitializeGalaxy(500, 5e21, 5.8e22, 5.8e22)
		g1 := InitializeGalaxy(500, 5e21, 6.2e22, 6.2e22)
		width := 1.0e23
		galaxies := []Galaxy{g0,g1}
		initialUniverse := InitializeUniverse(galaxies, width)
		numGens := 40000
		time := 2e15
		theta := 0.5
		timePoints := BarnesHut(initialUniverse, numGens, time, theta)
		canvasWidth := 1000
		frequency := 100
		scalingFactor := 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
		imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "Collision")
		fmt.Println("GIF drawn.")
	}
}

func CreateJupiterSystem() *Universe {
	// declaring objects
	var jupiter, io, europa, ganymede, callisto Star

	jupiter.red, jupiter.green, jupiter.blue = 223, 227, 202
	io.red, io.green, io.blue = 249, 249, 165
	europa.red, europa.green, europa.blue = 132, 83, 52
	ganymede.red, ganymede.green, ganymede.blue = 76, 0, 153
	callisto.red, callisto.green, callisto.blue = 0, 153, 76

	jupiter.mass = 1.898 * math.Pow(10, 27)
	io.mass = 8.9319 * math.Pow(10, 22)
	europa.mass = 4.7998 * math.Pow(10, 22)
	ganymede.mass = 1.4819 * math.Pow(10, 23)
	callisto.mass = 1.0759 * math.Pow(10, 23)

	jupiter.radius = 4000000
	io.radius = 1821000
	europa.radius = 1569000
	ganymede.radius = 2631000
	callisto.radius = 2410000

	jupiter.position.x, jupiter.position.y = 2000000000, 2000000000
	io.position.x, io.position.y = 2000000000-421600000, 2000000000
	europa.position.x, europa.position.y = 2000000000, 2000000000+670900000
	ganymede.position.x, ganymede.position.y = 2000000000+1070400000, 2000000000
	callisto.position.x, callisto.position.y = 2000000000, 2000000000-1882700000

	jupiter.velocity.x, jupiter.velocity.y = 0, 0
	io.velocity.x, io.velocity.y = 0, -17320
	europa.velocity.x, europa.velocity.y = -13740, 0
	ganymede.velocity.x, ganymede.velocity.y = 0, 10870
	callisto.velocity.x, callisto.velocity.y = 8200, 0

	// declaring universe and setting its fields.
	var jupiterSystem Universe
	jupiterSystem.width = 4000000000
	jupiterSystem.AddBody(&jupiter)
	jupiterSystem.AddBody(&io)
	jupiterSystem.AddBody(&europa)
	jupiterSystem.AddBody(&ganymede)
	jupiterSystem.AddBody(&callisto)
	return &jupiterSystem
}


func (u *Universe) AddBody(b *Star) {
	u.stars = append(u.stars, b)
}
