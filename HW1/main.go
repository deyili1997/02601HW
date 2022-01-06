package main

import (
	"flag"
	"fmt"
	"os"
  "math/rand"
  "time"
  "strconv"
  "io"
)

func main() {
	// Register command-line flags.
  rand.Seed(time.Now().UnixNano()) // Seed the random number generator.
  command := os.Args[1]
  if command == "read" {
    prefixLength,err1 := strconv.Atoi(os.Args[2])
    if err1 != nil  {
  		panic("There is something wrong string convertion!")
  	}
    tableOutName := os.Args[3]
    prefixLen := flag.Int("prefix", prefixLength, "prefix length in words")
    flag.Parse()                     // Parse command-line flags.
    c := NewChain(*prefixLen)     // Initialize a new Chain.
    for i:=4; i < len(os.Args); i++ {
      var r io.Reader
      var err error
      r, err = os.Open(os.Args[i])
      if err != nil {
        panic("Something wrong with io.Reader")
      }
      c.Build(r)             // Build chains from standard input.
    }
    c.WriteTable(tableOutName)
  }

  if command == "generate" {
    modelFileName := os.Args[2]
    wordNum,err2 := strconv.Atoi(os.Args[3])
    if err2 != nil  {
  		panic("There is something wrong string convertion!")
  	}
    numWords := flag.Int("words", wordNum, "maximum number of words to print")
    flag.Parse()                     // Parse command-line flags.
    newChain := ReadInTable(modelFileName)
  	text := newChain.Generate(*numWords) // Generate text.
  	fmt.Println(text)             // Write text to standard output.
  }
}
