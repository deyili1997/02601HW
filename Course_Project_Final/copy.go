package main

//this is a support document that copies objects
func (p *phageGroup) CopyPhageGroup() *phageGroup {
	var copyPGroup phageGroup
	copyPGroup.group = make([]*phage, 0)
	for _, j := range p.group {
		copyPGroup.group = append(copyPGroup.group, j.CopyPhage())
	}
	return &copyPGroup
}

func (b *bacteriaGroup) CopyBacteriaGroup() *bacteriaGroup {
	var copyBGroup bacteriaGroup
	copyBGroup.group = make([]*bacteria, 0)
	for _, j := range b.group {
		copyBGroup.group = append(copyBGroup.group, j.CopyBacteria())
	}
	return &copyBGroup
}

func (i *immuneGroup) CopyImmunCellGroup() *immuneGroup {
	var copyImmunGroup immuneGroup
	copyImmunGroup.group = make([]*immuneCell, 0)
	for _, j := range i.group {
		copyImmunGroup.group = append(copyImmunGroup.group, j.CopyImmuCell())
	}
	return &copyImmunGroup
}

func (b *bacteria) CopyBacteria() *bacteria {
	var newBacteria bacteria
	newBacteria.organism = b.organism
	newBacteria.bacteriaType = b.bacteriaType
	newBacteria.recombinationSite = b.recombinationSite
	if b.insidePhage != nil {
		newBacteria.insidePhage = b.insidePhage.CopyPhage()
	} else {
		newBacteria.insidePhage = nil
	}
	newBacteria.radius = b.radius
	newBacteria.scale = b.scale
	newBacteria.red = b.red
	newBacteria.green = b.green
	newBacteria.blue = b.blue
	newBacteria.replicateTimer = b.replicateTimer
	newBacteria.lysisTimer = b.lysisTimer
	newBacteria.lifeSpan = b.lifeSpan

	return &newBacteria
}

func (p *phage) CopyPhage() *phage {
	var newPhage phage
	newPhage.organism = p.organism
	newPhage.shellGene = p.shellGene
	newPhage.otherGene = p.otherGene
	newPhage.lysisStart = p.lysisStart
	newPhage.radius = p.radius
	newPhage.scale = p.scale
	newPhage.red = p.red
	newPhage.green = p.green
	newPhage.blue = p.blue
	return &newPhage
}

func (i *immuneCell) CopyImmuCell() *immuneCell {
	var newImmunCell immuneCell
	newImmunCell.organism = i.organism
	newImmunCell.huntingRange = i.huntingRange
	newImmunCell.radius = i.radius
	newImmunCell.scale = i.scale
	newImmunCell.red = i.red
	newImmunCell.green = i.green
	newImmunCell.blue = i.blue

	return &newImmunCell
}

func (f *frame) CopyFrame() *frame {
	var nf frame
	nf.pG = f.pG.CopyPhageGroup()
	nf.bG = f.bG.CopyBacteriaGroup()
	nf.iG = f.iG.CopyImmunCellGroup()
	nf.infectedG = f.infectedG.CopyBacteriaGroup()
	return &nf
}
