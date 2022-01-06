package main

import (
	"math/rand"
)

func InitializePlasm(w int) *plasm {
	var plasm plasm
	plasm.width = w
	return &plasm
}

func GeneratePhageGroup(c *plasm, n int) *phageGroup {
	var phageList phageGroup
	phageList.group = make([]*phage, n)
	for i := 0; i < n; i++ {
		phageList.group[i] = GeneratePhage(c)
	}
	return &phageList
}

func GeneratePhage(space *plasm) *phage {
	var p phage
	p.position = GenerateOrganism(space)
	probability1 := rand.Float64()
	if probability1 < 0.98 {
		p.shellGene = "normal"
	} else {
		p.shellGene = "mutated"
	}
	probability2 := rand.Float64()
	if probability2 < 0.8 {
		p.otherGene = "normal"
	} else {
		p.otherGene = "mutated"
	}
	p.lysisStart = false
	p.radius = 1
	p.scale = 1
	p.red = 255
	p.green = 255
	p.blue = 0
	return &p
}

func GenerateBacteriaGroup(c *plasm, n, bacLifespan, bacReplicateRate int) *bacteriaGroup {
	var bacteriaList bacteriaGroup
	bacteriaList.group = make([]*bacteria, n)
	for i := 0; i < n; i++ {
		bacteriaList.group[i] = GenerateBacteria(c, bacLifespan, bacReplicateRate)
	}
	return &bacteriaList
}

func GenerateBacteria(space *plasm, bacLifespan, bacReplicateRate int) *bacteria {
	var b bacteria
	b.position = GenerateOrganism(space)
	probability1 := rand.Float64()
	if probability1 < 0.95 {
		b.recombinationSite = "normal"
	} else {
		b.recombinationSite = "mutated"
	}
	probability2 := rand.Float64()
	if probability2 <= 0.3 {
		b.bacteriaType = "A"
	} else if probability2 > 0.3 && probability2 <= 0.6 {
		b.bacteriaType = "B"
	} else if probability2 > 0.6 && probability2 <= 0.95 {
		b.bacteriaType = "C"
	} else {
		b.bacteriaType = "D"
	}
	b.radius = 1
	b.scale = 2
	b.red = 0
	b.green = 255
	b.blue = 0
	b.replicateTimer = rand.Intn(bacReplicateRate)
	b.lysisTimer = 0
	b.lifeSpan = bacLifespan

	return &b
}

func GenerateImmuneCellaGroup(c *plasm, n int) *immuneGroup {
	var immuneCellList immuneGroup
	immuneCellList.group = make([]*immuneCell, n)
	for i := 0; i < n; i++ {
		immuneCellList.group[i] = GenerateImmuneCell(c)
	}
	return &immuneCellList
}

func GenerateImmuneCell(space *plasm) *immuneCell {
	var i immuneCell
	i.position = GenerateOrganism(space)
	i.huntingRange = 6
	i.radius = 3
	i.scale = 2
	i.red = 255
	i.green = 255
	i.blue = 255

	return &i
}

func GenerateOrganism(space *plasm) vec2 {
	position := vec2{rand.Intn(space.width + 1), rand.Intn(space.width + 1)}
	return position
}

func GenerateInitialFrame(p *phageGroup, b, infected *bacteriaGroup, i *immuneGroup) *frame {
	return &frame{
		pG:        p,
		bG:        b,
		infectedG: infected,
		iG:        i,
	}
}
