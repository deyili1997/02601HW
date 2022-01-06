package main

import (
  "strconv"
  "strings"
  "fmt"
  "os"
  "bufio"
  "math/rand"
)

// take the filename as the input and read in the frequencyTable
func ReadInTable(filename string) Chain {
  var ReadinChain Chain
  file := ReadFile(filename)
  scanner := bufio.NewScanner(file)
  rowCounter := 1
  for scanner.Scan() {
    if rowCounter == 1 {
      length, err := strconv.Atoi(scanner.Text())
      if err != nil {
        fmt.Println("Conversion error!")
      }
      ReadinChain.prefixLen = length
      ReadinChain.chain = make(map[string]map[string]int)
    } else {
      // fmt.Println(scanner.Text())
      ReadinChain.AppendMap(scanner.Text())
    }
    rowCounter++
  }
  // fmt.Println(ReadinChain.chain)
  // fmt.Println(ReadinChain.prefixLen)
  return ReadinChain
}

func ReadFile (fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: Something went wrong when opening the file!")
		fmt.Println("Error: Probably you gave the wrong filename!")
	}
	return file
}
//read in each row and append to the map
func (c *Chain) AppendMap(row string) {
  itemList := strings.Split(row," ")
  itemList = itemList[:len(itemList)-1]
  // fmt.Println(itemList)
  prefixList := make([]string,0)
  for i := 0; i < c.prefixLen; i++ {
    prefixList = append(prefixList, itemList[i])
  }
  prefix := strings.Join(prefixList, " ")
  c.chain[prefix] = make(map[string]int)
  for j := c.prefixLen; j < len(itemList); j = j+2 {
    // fmt.Println(c.prefixLen,len(itemList))
    // fmt.Println(j)
    count, err := strconv.Atoi(itemList[j+1])
    if err != nil {
      fmt.Println("Conversion error!")
    }
    // fmt.Println(itemList[j],count)
    c.chain[prefix][itemList[j]] = count
  }
  // fmt.Println(c.chain[prefix])
}



// "Generate" returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
  for n := 0; n <c.prefixLen; n++ {
    p[n] = "\"\""
  }
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
    // if it has nothing to choose from, stop the program
    if len(choices) == 0 {
			break
		}
    // simulate the random choice according to frequency
    choiceList := make(Prefix, 0)
    for m,n := range choices {
      for j:=1; j <= n; j++ {
        choiceList = append(choiceList, m)
      }
    }
    fmt.Println("*****START*****",choiceList,"******END******")
		next := choiceList[rand.Intn(len(choiceList))]
    fmt.Println("&&&&&The next chosen word is&&&&&",next,"&&&&&&")
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}
