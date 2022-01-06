package main
import (
	"fmt"
	"math"
)

//BarnesHut is our highest level function.
//Input: initial Universe object, a number of generations, and a time interval.
//Output: collection of Universe objects corresponding to updating the system
//over indicated number of generations every given time interval.
func BarnesHut(initialUniverse *Universe, numGens int, time, theta float64) []*Universe {
	timePoints := make([]*Universe, numGens+1)
	// Your code goes here. Use subroutines! :)
	timePoints[0] = initialUniverse
	for gen := 0; gen < numGens; gen++ {
		timePoints[gen+1] = UpdateUniverse(timePoints[gen], time, theta)
		fmt.Println("Finish appending Universe", gen + 1)
	}
	return timePoints
}
//
func UpdateUniverse(currUniv *Universe, time, theta float64) *Universe {
	newUniv := CopyUniverse(currUniv)
	//rule out the stars that has flied out of space
	for i, s := range newUniv.stars {
		if HasOutOfSpace(s, newUniv) {
			newUniv.stars = append(newUniv.stars[:i],newUniv.stars[i+1:]...)
		}
	}//start building new tree
	var tree QuadTree
	var root Node
	root.sector.x = 0
	root.sector.y = 0
	root.sector.width = newUniv.width
	tree.root = &root
	for i := 0; i <len(newUniv.stars);i++ {
		tree.root.InsertToNode(newUniv.stars[i])
	}

	for b := range newUniv.stars {
		// update pos, vel and accel
		newUniv.stars[b].Update(&tree, time, theta)
	}
	return newUniv
}

func HasOutOfSpace(s *Star, u *Universe) bool {
	if s.position.x < 0 || s.position.x > u.width || s.position.y <0 || s.position.y > u.width {
		return true
	}
	return false
}


func (s *Star) Update(tree *QuadTree, time, theta float64) {
	acc := s.NewAccel(tree, theta)
	vel := s.NewVelocity(time)
	pos := s.NewPosition(time)
	s.acceleration, s.velocity, s.position = acc, vel, pos
}

//Update position, velocity, acceleration according to the QuadTree
func (s *Star) NewAccel(tree *QuadTree, theta float64) OrderedPair {
	netForce := tree.root.ComputeNetForce(theta, s)
	return OrderedPair {
		x: netForce.x / s.mass,
		y: netForce.y / s.mass,
	}
}


func (s *Star) NewVelocity(t float64) OrderedPair {
	return OrderedPair{
		x: s.velocity.x + s.acceleration.x*t,
		y: s.velocity.y + s.acceleration.y*t,
	}
}

func (s *Star) NewPosition(t float64) OrderedPair {
	return OrderedPair{
		x: s.position.x + s.velocity.x*t + 0.5*s.acceleration.x*t*t,
		y: s.position.y + s.velocity.y*t + 0.5*s.acceleration.y*t*t,
	}
}
//Build tree function
func (this *Node) InsertToNode(newS *Star) {
	numLeaf := NodeType(this)
	if numLeaf == "EN" {	//this is a empty node

		this.star = newS
	}
		if numLeaf == "TS" { //this is a node with true star

			this.children = make([]*Node, 4)
			for i := range this.children {
				this.children[i] = this.CreatSubNode(i)
			}
			secId1 := this.GetSector(this.star)
			this.children[secId1].InsertToNode(this.star)
			secId2 := this.GetSector(newS)
			this.children[secId2].InsertToNode(newS)
			var dummy Star
			dummy.mass = this.star.mass + newS.mass
			dummy.position = OrderedPair{
				x: (this.star.mass * this.star.position.x + newS.position.x * newS.mass)/(this.star.mass + newS.mass),
				y: (this.star.mass * this.star.position.y + newS.position.y * newS.mass)/(this.star.mass + newS.mass),
			}
			this.star = &dummy
		}
		if numLeaf == "DS" { //this is a node with dummy star
			this.UpdateDummyPos(newS)
			this.UpdateDummyMass(newS)
			secId := this.GetSector(newS)
			this.children[secId].InsertToNode(newS)
		}
	}

//get which sector the star belong to
func (n *Node) GetSector(s *Star) int {
	if s.position.x < (n.sector.x + n.sector.width/2) && s.position.y < (n.sector.y + n.sector.width/2) {
		return 2
	}
	if s.position.x < (n.sector.x + n.sector.width/2) && s.position.y >= (n.sector.y + n.sector.width/2) {
		return 0
	}
	if s.position.x >= (n.sector.x + n.sector.width/2) && s.position.y < (n.sector.y + n.sector.width/2) {
		return 3
	}
	if s.position.x >= (n.sector.x + n.sector.width/2) && s.position.y >= (n.sector.y + n.sector.width/2) {
		return 1
	}
	panic("You should not go here")
	return -1
}
//creat a subnode and give sector info
func (n *Node) CreatSubNode(id int) *Node {
	var newN Node
	if id == 0 {
		newN.sector = Quadrant {
			x: n.sector.x,
			y: n.sector.y + n.sector.width/2,
			width: n.sector.width/2,
		}
	}
	if id == 1 {
		newN.sector = Quadrant {
			x: n.sector.x + n.sector.width/2,
			y: n.sector.y + n.sector.width/2,
			width: n.sector.width/2,
		}
	}
	if id == 2 {
		newN.sector = Quadrant {
			x: n.sector.x,
			y: n.sector.y,
			width: n.sector.width/2,

		}
	}
	if id == 3 {
		newN.sector = Quadrant {
			x: n.sector.x + n.sector.width/2,
			y: n.sector.y,
			width: n.sector.width/2,
		}
	}
	return &newN
}

//Update dummy star position and mass
func (n *Node) UpdateDummyPos (s *Star) {
	n.star.position.x = (n.star.position.x * n.star.mass + s.position.x * s.mass)/(n.star.mass + s.mass)
	n.star.position.y = (n.star.position.y * n.star.mass + s.position.y * s.mass)/(n.star.mass + s.mass)
}


func (n *Node) UpdateDummyMass (s *Star) {
	n.star.mass += s.mass
}

func NodeType(n *Node) string {
	var nodeType string
	if n.star == nil{
		nodeType = "EN"//Empty Node
	} else {
		if len(n.children) == 0 {
			nodeType = "TS"//True Star
		}
		if len(n.children) > 0 {
			nodeType = "DS"// Dummy Star
		}
		}
	return nodeType
	}



func (v *OrderedPair)Add(v2 OrderedPair) {
	v.x += v2.x
	v.y += v2.y
}
//Compute the force for certain star according to the tree
func (this *Node) ComputeNetForce (theta float64, b *Star) OrderedPair {
  nodeType := NodeType(this)
	var force OrderedPair
  //if the node contains no star or it is the star needed to compute itself
  if this.star == b || nodeType == "EN" {
  }
  //if the node contains a true star
  if this.star != b && nodeType == "TS" {
		return ComputeGravityForce(this.star, b)
  }
  //if the node contains a dummy star
  if nodeType == "DS" {
    d := Dist(this.star, b)
    s := this.sector.width
    if s/d <= theta {
			return ComputeGravityForce(this.star, b)
    } else {
      for _,c := range this.children {
        force.Add(c.ComputeNetForce(theta, b))
      }
    }
  }
	return force
}



// ComputeGravityForce computes the gravity force between body 1 and body 2.
func ComputeGravityForce(b1, b2 *Star) OrderedPair {
	d := Dist(b1, b2)
	deltaX := b2.position.x - b1.position.x
	deltaY := b2.position.y - b1.position.y
	F := G * b1.mass * b2.mass / (d * d)

	return OrderedPair{
		x: - (F * deltaX / d),
		y: - (F * deltaY / d),
	}
}

// Compute the Euclidian Distance between two bodies
func Dist(b1, b2 *Star) float64 {
	dx := b1.position.x - b2.position.x
	dy := b1.position.y - b2.position.y
	return math.Sqrt(dx*dx + dy*dy)
}
