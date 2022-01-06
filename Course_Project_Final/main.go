package main

import (
	"flag"
	"fmt"
	"gifhelper"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var plasmaSize, bacNum, phageNum, immuneCellNum, bacLifespan, bacReplicateRate, lysisRate, phageOffspring, numGens int
	flag.IntVar(&plasmaSize, "plasmaSize", 600, "The size of the plasma area.")
	flag.IntVar(&bacNum, "bacNum", 500, "Number of uninfected bacteria in the plasma.")
	flag.IntVar(&phageNum, "phageNum", 300, "Number of phages in the plasma.")
	flag.IntVar(&immuneCellNum, "immuneCellNum", 30, "Number of immune cells in the plasma.")
	flag.IntVar(&bacLifespan, "bacLifespan", 600, "How long a bacteria will live.")
	flag.IntVar(&bacReplicateRate, "bacReplicateRate", 200, "The rate of a bacteria replicating itself.")
	flag.IntVar(&lysisRate, "lysisRate", 80, "The rate of an infected bacteria being lysed by the inside phage.")
	flag.IntVar(&phageOffspring, "phageOffspring", 4, "The number of offsprings a phage can produce.")
	flag.IntVar(&numGens, "numGens", 2000, "The number of generations.")

	flag.Parse()

	c := InitializePlasm(plasmaSize)

	b := GenerateBacteriaGroup(c, bacNum, bacLifespan, bacReplicateRate)
	infected := bacteriaGroup{
		group: []*bacteria{},
	}
	i := GenerateImmuneCellaGroup(c, immuneCellNum)
	p := GeneratePhageGroup(c, phageNum)
	ini := GenerateInitialFrame(p, b, &infected, i)
	fmt.Println("********************Start simulating, please wait***************************")
	Result := Engine(ini, c, bacLifespan, bacReplicateRate, lysisRate, phageOffspring, numGens)
	fmt.Println("***************Simulation has finished, start drawing the gif, please wait**************************")

	imgWidth := 1000
	frameRate := 2
	pics := CreateAnimationFrames(Result, imgWidth, frameRate, plasmaSize, bacLifespan)
	gifhelper.ImagesToGIF(pics, "cartoon")
	fmt.Println("Finish drawing the gif!!")

}
