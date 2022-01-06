package main

import (
	"flag"
	"fmt"
	"gifhelper"
	"math"
)

// gravitataional constant
const G = 2.0 * 6.67408e-11
const moonScale = 6

// Universe represents the entire universe
type Universe struct {
	bodies []Body
	width  float64
}

// AddBody adds a body to the universe
func (u *Universe) AddBody(b Body, scale float64) {
	b.scale = scale
	u.bodies = append(u.bodies, b)
}

func (u *Universe) PrettyPrint() {
	for _, b := range u.bodies {
		b.PrettyPrint()
	}
}

// Body represents a moon or planet
type Body struct {
	name             string
	mass, radius     float64
	pos, vel, acc    Vec2
	red, green, blue int
	scale            float64
}

func (b *Body) PrettyPrint() {
	fmt.Printf("   %s %v %v %v\n",
		b.name,
		b.pos, b.vel, b.acc,
	)
}

// methods:

// Vec2 is a 2d vector.
type Vec2 struct {
	x, y float64
}

func (v *Vec2) Add(v2 Vec2) {
	v.x += v2.x
	v.y += v2.y
}

// Compute the Euclidian Distance between two bodies
func Dist(b1, b2 Body) float64 {
	dx := b1.pos.x - b2.pos.x
	dy := b1.pos.y - b2.pos.y
	return math.Sqrt(dx*dx + dy*dy)
}

// SimulateUniverse creates numGen universes after the initUniv.
func SimulateUniverse(initUniv Universe, numGen int, t float64) []Universe {
	universes := make([]Universe, numGen+1)
	universes[0] = initUniv
	for gen := 0; gen < numGen; gen++ {
		universes[gen+1] = UpdateUniverse(universes[gen], t)
	}
	return universes
}

// UpdateUniverse returns a new universe after time t.
func UpdateUniverse(univ Universe, t float64) Universe {

	newUniverse := CopyUniverse(univ)
	for b := range univ.bodies {
		// update pos, vel and accel
		newUniverse.bodies[b].Update(univ, t)
	}

	return newUniverse
}

// need to change this to use current acc/vel/pos
func (b *Body) Update(univ Universe, t float64) {
	acc := b.NewAccel(univ)
	vel := b.NewVelocity(t)
	pos := b.NewPosition(t)
	b.acc, b.vel, b.pos = acc, vel, pos
}

// NewVelocity makes the velocity of this object consistent with the acceleration.
func (b *Body) NewVelocity(t float64) Vec2 {
	return Vec2{
		x: b.vel.x + b.acc.x*t,
		y: b.vel.y + b.acc.y*t,
	}
}

// NewPosition computes the new poosition given the updated acc and velocity.
//
// Assumputions: constant acceleration over a time step.
// => DeltaX = v_avg * t
//    DeltaX = (v_start + v_final)*t/ 2
// because v_final = v_start + acc*t:
//	  DeltaX = (v_start + v_start + acc*t)t/2
// Simplify:
//	DeltaX = v_start*t + 0.5acc*t*t
// =>
//  NewX = v_start*t + 0.5acc*t*t + OldX
//
func (b *Body) NewPosition(t float64) Vec2 {
	return Vec2{
		x: b.pos.x + b.vel.x*t + 0.5*b.acc.x*t*t,
		y: b.pos.y + b.vel.y*t + 0.5*b.acc.y*t*t,
	}
}

// UpdateAccel computes the new accerlation vector for b
func (b *Body) NewAccel(univ Universe) Vec2 {
	F := ComputeNetForce(univ, *b)
	return Vec2{
		x: F.x / b.mass,
		y: F.y / b.mass,
	}
}

// CopyUniverse creates a copy of a universe
func CopyUniverse(univ Universe) Universe {
	u := univ
	u.bodies = make([]Body, len(univ.bodies))
	for i, b := range univ.bodies {
		u.bodies[i] = b
	}
	return u
}

// ComputeNetForce sums the forces of all bodies in the universe
// acting on b.
func ComputeNetForce(univ Universe, b Body) Vec2 {
	var netForce Vec2
	for _, body := range univ.bodies {
		if body != b {
			f := ComputeGravityForce(b, body)
			netForce.Add(f)
		}
	}
	return netForce
}

// ComputeGravityForce computes the gravity force between body 1 and body 2.
func ComputeGravityForce(b1, b2 Body) Vec2 {
	d := Dist(b1, b2)
	deltaX := b2.pos.x - b1.pos.x
	deltaY := b2.pos.y - b1.pos.y
	F := G * b1.mass * b2.mass / (d * d)

	return Vec2{
		x: F * deltaX / d,
		y: F * deltaY / d,
	}
}

func CreateJupiterSystem() Universe {
	// declaring objects
	var jupiter, io, europa, ganymede, callisto Body

	jupiter.name = "Jupiter"
	io.name = "Io"
	europa.name = "Europa"
	ganymede.name = "Ganymede"
	callisto.name = "Callisto"

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

	jupiter.radius = 71000000
	io.radius = 1821000
	europa.radius = 1569000
	ganymede.radius = 2631000
	callisto.radius = 2410000

	jupiter.pos.x, jupiter.pos.y = 2000000000, 2000000000
	io.pos.x, io.pos.y = 2000000000-421600000, 2000000000
	europa.pos.x, europa.pos.y = 2000000000, 2000000000+670900000
	ganymede.pos.x, ganymede.pos.y = 2000000000+1070400000, 2000000000
	callisto.pos.x, callisto.pos.y = 2000000000, 2000000000-1882700000

	jupiter.vel.x, jupiter.vel.y = 0, 0
	io.vel.x, io.vel.y = 0, -17320
	europa.vel.x, europa.vel.y = -13740, 0
	ganymede.vel.x, ganymede.vel.y = 0, 10870
	callisto.vel.x, callisto.vel.y = 8200, 0

	// declaring universe and setting its fields.
	var jupiterSystem Universe
	jupiterSystem.width = 4000000000
	jupiterSystem.AddBody(jupiter, 1)
	jupiterSystem.AddBody(io, moonScale)
	jupiterSystem.AddBody(europa, moonScale)
	jupiterSystem.AddBody(ganymede, moonScale)
	jupiterSystem.AddBody(callisto, moonScale)
	return jupiterSystem
}

func main() {
	jupiter := CreateJupiterSystem()
	fmt.Println(jupiter)

	var numGen int
	var t float64
	var imgWidth int
	var outputFilename string
	var animOutputFile string
	var frameRate int

	flag.IntVar(&numGen, "numGen", 1000000, "Number of steps to run the universe.")
	flag.Float64Var(&t, "t", 0.1, "Interval of each step.")
	flag.IntVar(&imgWidth, "width", 500, "Width (and height) of the image to create.")
	flag.StringVar(&outputFilename, "o", "out.png", "Name of PNG to output.")
	flag.StringVar(&animOutputFile, "a", "anim.gif", "Animated GIF to write.")
	flag.IntVar(&frameRate, "frameRate", 10000, "Frame writing interval")
	flag.Parse()

	evolution := SimulateUniverse(jupiter, numGen, t)
	/*
		for i, u := range evolution {
			fmt.Println("\nTIMESTEP", i)
			u.PrettyPrint()
		}
	*/

	// write out the final state of the universee
	img := evolution[len(evolution)-1].DrawToImage(imgWidth)
	WriteImageAsPNG(img, outputFilename)

	// write out an animation of the universe
	frames := CreateAnimationFrames(evolution, imgWidth, frameRate)
	gifhelper.ImagesToGIF(frames, animOutputFile)
}
