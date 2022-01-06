package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The highest level of frame generations.
func Engine(f *frame, c *plasm, bacLifespan, bacReplicateRate, lysisRate, phageOffspring, numGens int) []*frame {
	final := make([]*frame, 0)
	final = append(final, f)
	for i := 0; i < numGens; i++ {
		start := time.Now()
		newF := Update(final[i], c, bacLifespan, bacReplicateRate, lysisRate, phageOffspring)
		final = append(final, newF)
		elapsed := time.Now().Sub(start)
		fmt.Println("Finish No.", i, "frame, it took", elapsed.Seconds(), "seconds!")
		if elapsed.Seconds() > 0.5 {
			fmt.Println("Too many organisms, stop simulation!")
			break
		}
	}
	return final
}

//Compute each frame
func Update(f *frame, c *plasm, bacLifespan, bacReplicateRate, lysisRate, phageOffspring int) *frame {
	newFrame := f.CopyFrame()
	newFrame.EatenbyImmu()
	newFrame.InfectBac(lysisRate)
	newFrame.Lysis(phageOffspring)
	newFrame.UpdateLifespan()
	newFrame.BacReplicate(c, bacLifespan, bacReplicateRate)
	newFrame.OranismMove(c)
	return newFrame
}

//judge whether the phage got eaten by the immune cells
func (f *frame) EatenbyImmu() {
	deadZone := DeadZone(f.iG.group)

	for iP := 0; iP < len(f.pG.group); {
		if f.pG.group[iP].GetEatten(deadZone) {
			f.pG.group = append(f.pG.group[:iP], f.pG.group[(iP+1):]...)
		} else {
			iP++
		}
	}

	for iB := 0; iB < len(f.bG.group); {
		if f.bG.group[iB].GetEatten(deadZone) {
			f.bG.group = append(f.bG.group[:iB], f.bG.group[(iB+1):]...)
		} else {
			iB++
		}
	}

	for iIB := 0; iIB < len(f.infectedG.group); {
		if f.infectedG.group[iIB].GetEatten(deadZone) {
			f.infectedG.group = append(f.infectedG.group[:iIB], f.infectedG.group[(iIB+1):]...)
		} else {
			iIB++
		}
	}

}

func DeadZone(immunG []*immuneCell) []vec2 {
	deadZone := make([]vec2, 0)
	for _, each := range immunG {
		for x := each.position.x - each.huntingRange; x <= each.position.x+each.huntingRange; x++ {
			for y := each.position.y - each.huntingRange; y <= each.position.y+each.huntingRange; y++ {
				dot := vec2{x, y}
				deadZone = append(deadZone, dot)

			}
		}
	}
	return deadZone
}

//check whether any bacteria being infected by phages
func (f *frame) InfectBac(lysisRate int) {
	for iB := 0; iB < len(f.bG.group); {
		found := false
		for iP := 0; iP < len(f.pG.group); {
			if InfectPosition(f.pG.group[iP].position, f.bG.group[iB].position) && CanInject(f.bG.group[iB]) {
				f.bG.group[iB].insidePhage = f.pG.group[iP]
				f.bG.group[iB].SetTimer(lysisRate)
				f.infectedG.group = append(f.infectedG.group, f.bG.group[iB])
				f.pG.group = append(f.pG.group[:iP], f.pG.group[(iP+1):]...)
				f.bG.group = append(f.bG.group[:iB], f.bG.group[(iB+1):]...)

				found = true
				break
			} else {
				iP++
			}
		}
		if found == false {
			iB++
		}
	}

}

// if the bacteria was infected, then the counting down starts
func (b *bacteria) SetTimer(lysisRate int) {
	var newBacteria bacteria
	CopyBasicInfo(b, &newBacteria)
	newBacteria.insidePhage = b.insidePhage.CopyPhage()
	newBacteria.replicateTimer = b.replicateTimer
	newBacteria.lifeSpan = b.lifeSpan
	newBacteria.lysisTimer = rand.Intn(lysisRate)
	*b = newBacteria
}

func CopyBasicInfo(old, new *bacteria) {
	new.organism = old.organism
	new.bacteriaType = old.bacteriaType
	new.recombinationSite = old.recombinationSite
	new.radius = old.radius
	new.scale = old.scale
	new.red = old.red
	new.green = old.green
	new.blue = old.blue
}

//check whether the phage can infect bacteria
func CanInject(b *bacteria) bool {
	if (b.bacteriaType != "D") && b.recombinationSite == "normal" {
		return true
	}
	return false
}

func (o *organism) GetEatten(deadZone []vec2) bool {
	for _, dot := range deadZone {
		if SamePosition(o.position, dot) {
			return true
		}
	}
	return false
}

func SamePosition(p1, p2 vec2) bool {
	if p1.x == p2.x && p1.y == p2.y {
		return true
	}
	return false

}

func InfectPosition(phagePos, bacPos vec2) bool {
	if ((bacPos.x <= phagePos.x+2) && (bacPos.x >= phagePos.x-2)) &&
		((bacPos.y <= phagePos.y+2) && (bacPos.y >= phagePos.y-2)) {
		return true
	}
	return false

}

//check whether any infected bacteria should be lysed
func (f *frame) Lysis(phageOffspring int) {
	var roundOffspring phageGroup
	for q := 0; q < len(f.infectedG.group); q++ {
		f.infectedG.group[q].LysisCountDown()
		if f.infectedG.group[q].lysisTimer == 0 {
			eachOffspring := f.infectedG.group[q].insidePhage.Cycle(phageOffspring)
			//if this bacteria was lysed by the phage
			//exclude the bacteria from the infected bacteria group and append its phage offspring
			f.infectedG.group = append(f.infectedG.group[:q], f.infectedG.group[q+1:]...)
			roundOffspring.group = append(roundOffspring.group, eachOffspring.group...)
		}
	}
	//append all the offspring to the swimming phages, then copy, this is the intial phage group for the next round
	f.pG.group = append(f.pG.group, roundOffspring.group...)
}
func (b *bacteria) LysisCountDown() {
	var newBacteria bacteria
	CopyBasicInfo(b, &newBacteria)
	newBacteria.insidePhage = b.insidePhage.CopyPhage()
	newBacteria.replicateTimer = b.replicateTimer
	newBacteria.lifeSpan = b.lifeSpan
	newBacteria.lysisTimer = b.lysisTimer - 1
	*b = newBacteria
}

func (p *phage) Cycle(phageOffspring int) *phageGroup {
	var offSpring phageGroup
	offSpring.group = make([]*phage, phageOffspring)
	for i := range offSpring.group {
		offSpring.group[i] = p.CopyPhage()
	}

	return &offSpring
}

// after several steps, bacteria replicate itself
func (f *frame) BacReplicate(c *plasm, bacLifespan, bacReplicateRate int) {
	newBorn := make([]*bacteria, 0)
	for i := range f.bG.group {
		f.bG.group[i].ReplicateCountDown()
		if f.bG.group[i].replicateTimer == 0 {
			f.bG.group[i].ResetreplicateTimer(bacReplicateRate)
			newBorn = append(newBorn, GiveBirth(f.bG.group[i], bacLifespan, bacReplicateRate))
		}
	}
	f.bG.group = append(f.bG.group, newBorn...)
}

//bacteria replciate itself
func GiveBirth(parent *bacteria, bacLifespan, bacReplicateRate int) *bacteria {
	var b bacteria
	b.position = parent.position
	probability1 := rand.Float64()
	if parent.recombinationSite == "normal" {
		if probability1 < 0.95 {
			b.recombinationSite = "normal"
		} else {
			b.recombinationSite = "mutated"
		}
	} else if parent.recombinationSite == "mutated" {
		if probability1 < 0.5 {
			b.recombinationSite = "mutated"
		} else {
			b.recombinationSite = "normal"
		}
	}
	probability2 := rand.Float64()
	if parent.bacteriaType != "D" {
		if probability2 <= 0.3 {
			b.bacteriaType = "A"
		} else if probability2 > 0.3 && probability2 <= 0.6 {
			b.bacteriaType = "B"
		} else if probability2 > 0.6 && probability2 <= 0.95 {
			b.bacteriaType = "C"
		} else {
			b.bacteriaType = "D"
		}
	} else if parent.bacteriaType == "D" {
		if probability2 <= 0.2 {
			b.bacteriaType = "A"
		} else if probability2 > 0.2 && probability2 <= 0.4 {
			b.bacteriaType = "B"
		} else if probability2 > 0.4 && probability2 <= 0.6 {
			b.bacteriaType = "C"
		} else {
			b.bacteriaType = "D"
		}
	}
	b.radius = 1
	b.scale = 2
	b.red = 0
	b.green = 255
	b.blue = 0
	b.replicateTimer = rand.Intn(bacReplicateRate)
	b.lifeSpan = bacLifespan
	b.lysisTimer = 0

	return &b
}

func (b *bacteria) ReplicateCountDown() {
	var newBacteria bacteria
	CopyBasicInfo(b, &newBacteria)
	newBacteria.lifeSpan = b.lifeSpan
	newBacteria.replicateTimer = b.replicateTimer - 1
	newBacteria.lysisTimer = b.lysisTimer
	*b = newBacteria
}
func (b *bacteria) ResetreplicateTimer(bacReplicateRate int) {
	var newBacteria bacteria
	CopyBasicInfo(b, &newBacteria)
	newBacteria.lifeSpan = b.lifeSpan
	newBacteria.replicateTimer = rand.Intn(bacReplicateRate)
	newBacteria.lysisTimer = b.lysisTimer
	*b = newBacteria
}

//all organisms move one step
func (f *frame) OranismMove(c *plasm) {
	for _, eachP := range f.pG.group {
		eachP.UpdatePosition(c)
	}
	for _, each := range f.bG.group {
		each.UpdatePosition(c)
	}

	for _, each := range f.infectedG.group {
		each.UpdatePosition(c)
	}
	BacImmuCells := make([]*immuneCell, 0)
	PhageImmuCells := make([]*immuneCell, 0)
	InfectedBacImmuCells := make([]*immuneCell, 0)
	for i, each := range f.iG.group {
		if ((i % 7) == 0) || ((i % 7) == 1) || ((i % 7) == 2) || ((i % 7) == 3) {
			BacImmuCells = append(BacImmuCells, each)
		} else if ((i % 7) == 4) || ((i % 7) == 5) {

			PhageImmuCells = append(PhageImmuCells, each)
		} else if (i % 7) == 6 {

			InfectedBacImmuCells = append(InfectedBacImmuCells, each)
		}

	}

	overlapBacIm := make([]*immuneCell, 0)
	for i := 0; i < len(BacImmuCells); i++ {
		for j := i + 1; j < len(BacImmuCells); {
			if SamePosition(BacImmuCells[i].position, BacImmuCells[j].position) {
				overlapBacIm = append(overlapBacIm, BacImmuCells[j])
				BacImmuCells = append(BacImmuCells[:j], BacImmuCells[j+1:]...)
			} else {
				j++
			}
		}
	}
	overlapPhageIm := make([]*immuneCell, 0)
	for i := 0; i < len(PhageImmuCells); i++ {
		for j := i + 1; j < len(PhageImmuCells); {
			if SamePosition(PhageImmuCells[i].position, PhageImmuCells[j].position) {
				overlapPhageIm = append(overlapPhageIm, PhageImmuCells[j])
				PhageImmuCells = append(PhageImmuCells[:j], PhageImmuCells[j+1:]...)
			} else {
				j++
			}
		}
	}
	overlapInfectedIm := make([]*immuneCell, 0)
	for i := 0; i < len(InfectedBacImmuCells); i++ {
		for j := i + 1; j < len(InfectedBacImmuCells); {
			if SamePosition(InfectedBacImmuCells[i].position, InfectedBacImmuCells[j].position) {
				overlapInfectedIm = append(overlapInfectedIm, InfectedBacImmuCells[j])
				InfectedBacImmuCells = append(InfectedBacImmuCells[:j], InfectedBacImmuCells[j+1:]...)
			} else {
				j++
			}
		}
	}

	for _, each := range BacImmuCells {
		if len(f.bG.group) > 0 {
			each.ApproachBac(f)
		} else {
			each.UpdatePosition(c)
		}
	}
	for _, each := range overlapBacIm {
		each.UpdatePosition(c)
	}
	for _, each := range PhageImmuCells {
		if len(f.pG.group) > 0 {
			each.ApproachPhage(f)
		} else {
			each.UpdatePosition(c)
		}
	}
	for _, each := range overlapPhageIm {
		each.UpdatePosition(c)
	}
	for _, each := range InfectedBacImmuCells {
		if len(f.infectedG.group) > 0 {
			each.ApproachInfected(f)
		} else {
			each.UpdatePosition(c)
		}
	}
	for _, each := range overlapInfectedIm {
		each.UpdatePosition(c)
	}
}

func (im *immuneCell) ApproachBac(f *frame) {
	cloest := f.bG.group[0]
	D := (f.bG.group[0].position.x-im.position.x)*(f.bG.group[0].position.x-im.position.x) +
		(f.bG.group[0].position.y-im.position.y)*(f.bG.group[0].position.y-im.position.y)
	for i, each := range f.bG.group {
		dis := (f.bG.group[i].position.x-im.position.x)*(f.bG.group[i].position.x-im.position.x) +
			(f.bG.group[i].position.y-im.position.y)*(f.bG.group[i].position.y-im.position.y)
		if dis < D {
			D = dis
			cloest = each
		}
	}
	im.Approach(cloest.position)
}
func (im *immuneCell) ApproachPhage(f *frame) {
	cloest := f.pG.group[0]
	D := (f.pG.group[0].position.x-im.position.x)*(f.pG.group[0].position.x-im.position.x) +
		(f.pG.group[0].position.y-im.position.y)*(f.pG.group[0].position.y-im.position.y)
	for i, each := range f.pG.group {
		dis := (f.pG.group[i].position.x-im.position.x)*(f.pG.group[i].position.x-im.position.x) +
			(f.pG.group[i].position.y-im.position.y)*(f.pG.group[i].position.y-im.position.y)
		if dis < D {
			D = dis
			cloest = each
		}
	}
	im.Approach(cloest.position)
}

func (im *immuneCell) ApproachInfected(f *frame) {
	cloest := f.infectedG.group[0]
	D := (f.infectedG.group[0].position.x-im.position.x)*(f.infectedG.group[0].position.x-im.position.x) +
		(f.infectedG.group[0].position.y-im.position.y)*(f.infectedG.group[0].position.y-im.position.y)
	for i, each := range f.infectedG.group {
		dis := (f.infectedG.group[i].position.x-im.position.x)*(f.infectedG.group[i].position.x-im.position.x) +
			(f.infectedG.group[i].position.y-im.position.y)*(f.infectedG.group[i].position.y-im.position.y)
		if dis < D {
			D = dis
			cloest = each
		}
	}
	im.Approach(cloest.position)
}

func (im *immuneCell) Approach(p vec2) {
	if im.position.x < p.x {
		im.position.x++
	} else if im.position.x > p.x {
		im.position.x--
	}

	if im.position.y < p.y {
		im.position.y++
	} else if im.position.y > p.y {
		im.position.y--
	}

}

//all organisms take 1 step each time
func (o *organism) UpdatePosition(c *plasm) {
	moveDecision := TossCoin()
	switch moveDecision[0] {
	case "Forward":
		if o.position.x < c.width {
			o.position.x++
		} else {
			o.position.x--
		}
	case "Backward":
		if o.position.x > 0 {
			o.position.x--
		} else {
			o.position.x++
		}
	}
	switch moveDecision[1] {
	case "Forward":
		if o.position.y < c.width {
			o.position.y++
		} else {
			o.position.y--
		}
	case "Backward":
		if o.position.y > 0 {
			o.position.y--
		} else {
			o.position.y++
		}
	}
}

func TossCoin() []string {
	moveDecisions := make([]string, 0)
	for i := 0; i < 3; i++ {
		// rand.Seed(time.Now().UnixNano())
		toss := rand.Float64()
		if toss < 0.5 {
			moveDecisions = append(moveDecisions, "Forward")
		} else {
			moveDecisions = append(moveDecisions, "Backward")
		}
	}
	return moveDecisions
}

//Update the lifespan of bacteria. if lifespan == 0 then the bacteria is dead.
func (f *frame) UpdateLifespan() {
	for iB := 0; iB < len(f.bG.group); {
		f.bG.group[iB].LifeSpanCountDown()
		if f.bG.group[iB].lifeSpan == 0 {
			f.bG.group = append(f.bG.group[:iB], f.bG.group[(iB+1):]...)
		} else {
			iB++
		}
	}
}

func (b *bacteria) LifeSpanCountDown() {
	var newBacteria bacteria
	CopyBasicInfo(b, &newBacteria)
	newBacteria.lifeSpan = b.lifeSpan - 1
	newBacteria.replicateTimer = b.replicateTimer
	newBacteria.lysisTimer = b.lysisTimer
	*b = newBacteria
}
