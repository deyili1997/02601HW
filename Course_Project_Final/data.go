package main

type vec2 struct {
	x int
	y int
}
type plasm struct {
	width int
}

type organism struct {
	position         vec2
	red, green, blue int
	radius, scale    int
}

type phage struct {
	organism
	shellGene  string
	otherGene  string
	lysisStart bool
}

type bacteria struct {
	organism
	bacteriaType      string
	recombinationSite string // if recombinationSite == "mutated" then phage cannot insert its genome into the bacteria genome
	insidePhage       *phage
	replicateTimer    int
	lysisTimer        int
	lifeSpan          int
}

type immuneCell struct {
	organism
	huntingRange int
}

type bacteriaGroup struct {
	group []*bacteria
}

type phageGroup struct {
	group []*phage
}

type immuneGroup struct {
	group []*immuneCell
}

type frame struct {
	pG            *phageGroup
	bG, infectedG *bacteriaGroup
	iG            *immuneGroup
}
