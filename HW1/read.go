package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)


// Prefix is a Markov chain prefix of one or more words.
type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	//Modification: the structure of chain should be: keys: the strings of prefixes;
	//values: frequency table of the sufixes
	chain map[string]map[string]int
	prefixLen int
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	//Modification: the structure of chain should be: keys: the strings of prefixes;
	//values: frequency table of the sufixes
	return &Chain{make(map[string]map[string]int), prefixLen}
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
var readRound int = 1
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	//loop througth the article
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		//set the start prefixes to ""
		for i := range p {
			if p[i] == "" {
				p[i] = "\"\""
			}
		}
		key := p.String()
		if _,ok := c.chain[key];!ok {
			c.chain[key] = make(map[string]int)
			//lable the indicies of each prefix for sorting later.
			c.chain[key]["readRound"] = readRound
			c.chain[key][s]++
		} else {
			c.chain[key][s]++
		}
		p.Shift(s)
		readRound++
	}
}

type prefixOrder struct {
	prefix string
	number int
}
// Generate a txt file and output the c.chain to it in the given format
func (c *Chain) WriteTable(tableName string) {
	tableOut, err := os.Create(tableName)
	if err != nil{
		fmt.Println("Sorry, we cannot creat the file!")
	}
	defer tableOut.Close()
	fmt.Fprintf(tableOut,"%v\n",c.prefixLen)
	//sort map
	orderList := make([]prefixOrder,0)
	for prefix, table := range c.chain {
		var pref prefixOrder
		pref.prefix = prefix
		pref.number = table["readRound"]
		orderList = append(orderList, pref)
	}
	orderList = SortStructList(orderList)

	for _,p := range orderList {
		fmt.Fprintf(tableOut,"%v ", p.prefix)
		for suffix, occurence := range c.chain[p.prefix] {
			if suffix == "readRound" {
				continue
			} else {
				fmt.Fprintf(tableOut,"%v %v ",suffix,occurence)
			}
		}
		fmt.Fprintf(tableOut,"\n")
	}
}

// Sort the order of prefixes as the order of reading into the map 
func SortStructList (orderList []prefixOrder) []prefixOrder {
	for n := 0; n <= len(orderList); n++ {
      for i := 1; i < len(orderList)-n; i++ {
          if orderList[i].number < orderList[i-1].number {
              orderList[i], orderList[i-1] = orderList[i-1], orderList[i]
			}
		}
	}
	return orderList
}
