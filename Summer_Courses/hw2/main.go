package main
import (
  "fmt"
  "math/rand"
  "time"
  "math"
  "strconv"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < 10; i++ {
    fmt.Println(WeightedDie())
    }
  fmt.Println(ShowTable())
  fmt.Println(RelativelyPrimeProbability(10,20,1000))
  list := []int{1,2,4,6,1}
  fmt.Println(HasRepeat(list))
  fmt.Println(BirthdayParadox(20,100))
  fmt.Println(ComputePeriodLength(list))
  fmt.Println(CountNumDigits(-12345))
  fmt.Println(SquareMiddle(123,2))
  fmt.Println(SquareMiddle(3972,4))
  fmt.Println(GenerateMiddleSquareSequence(1600,4))
  fmt.Println(GenerateMiddleSquareSequence(3792,4))
  fmt.Println(FourDigitCount())
}
// Exercise 1: Say that we would like to simulate a weighted die that has probability 0.5 in a given roll of resulting in a 3 and probability 0.1 of resulting in any of the other five numbers. Write and implement a function WeightedDie that models this die.
func WeightedDie() int {
  randFloat := rand.Float64()
  var diceResult int
  if randFloat <= 0.5 {
    diceResult = 3
  } else if randFloat > 0.5 && randFloat <= 0.6 {
    diceResult = 1
  } else if randFloat > 0.6 && randFloat <= 0.7 {
    diceResult = 2
  } else if randFloat > 0.7 && randFloat <= 0.8 {
    diceResult = 4
  } else if randFloat > 0.8 && randFloat <= 0.9 {
    diceResult = 5
  } else if randFloat > 0.9 && randFloat < 1.0 {
    diceResult = 6
  }
  return diceResult
}

// Exercise 2: Apply this approach to both TrivialGCD and EuclidGCD for ten randomly chosen pairs of numbers from each of the following ranges.
//
// 1,000 to 2,000
// 10,000 to 20,000
// 100,000 to 200,000
// 1,000,000 to 2,000,000
// Then, form a table comparing the average running time of each algorithm for values chosen from these ranges. What is the average speedup obtained by using EuclidGCD and how does this speedup vary between the different ranges?

func ShowTable () map[string] float64 {
  Trivial1 := Timing(1000)[0]
  Trivial2 := Timing(10000)[0]
  Trivial3 := Timing(100000)[0]
  Trivial4 := Timing(1000000)[0]
  Euclid1 := Timing(1000)[1]
  Euclid2 := Timing(10000)[1]
  Euclid3 := Timing(100000)[1]
  Euclid4 := Timing(1000000)[1]
  resultMap := make(map[string] float64)
  resultMap["TrivialResult of (1000,2000)"] = Trivial1
  resultMap["TrivialResult of (10000,20000)"] = Trivial2
  resultMap["TrivialResult of (100000,200000)"] = Trivial3
  resultMap["TrivialResult of (1000000,2000000)"] = Trivial4
  resultMap["EuclidResult of (1000,2000)"] = Euclid1
  resultMap["EuclidResult of (10000,20000)"] = Euclid2
  resultMap["EuclidResult of (100000,200000)"] = Euclid3
  resultMap["EuclidResult of (1000000,2000000)"] = Euclid4
  return resultMap
}

func Timing (rangebuttom int) []float64 {
  list := PairGenerator(rangebuttom)
  trivialTimeRecord := make([]float64,0)
  EuclidTimeRecord := make([]float64,0)
  finalResult := make([]float64,0)
  for i:= 0; i < 20; i += 2 {
    pair1 := list[i]
    pair2 := list[i+1]
    starttime1 := time.Now().UnixNano()
    TrivialGCD(pair1,pair2)
    endtime1 := time.Now().UnixNano()
    trivialTimeConsuming := float64(endtime1-starttime1)
    trivialTimeRecord = append(trivialTimeRecord, trivialTimeConsuming)
    starttime2 := time.Now().UnixNano()
    EuclidGCD(pair1,pair2)
    endtime2 := time.Now().UnixNano()
    EuclidTimeConsuming := float64(endtime2-starttime2)
    EuclidTimeRecord = append(EuclidTimeRecord, EuclidTimeConsuming)
  }
  var sum1, sum2 float64
  for _,j := range trivialTimeRecord {
    sum1 += j
  }
  for _,n := range EuclidTimeRecord {
    sum2 += n
  }
  finalResult = append(finalResult,sum1/10)
  finalResult = append(finalResult,sum2/10)
  return finalResult
}

func PairGenerator (rangebuttom int) []int {
  pairArray := make([]int, 0)
  for i := 0; i < 20; i++ {
    rand.Seed(time.Now().UnixNano())
    pairArray = append(pairArray, rand.Intn(rangebuttom) + rangebuttom)
  }
  return pairArray
}

func TrivialGCD (pair1 int, pair2 int) int {
  var smallNum int
  if pair1 > pair2 {
    smallNum = pair2
  } else {
    smallNum = pair1
  }
  for i:= smallNum; i >= 2; i-- {
    if pair1%i==0 && pair2%i == 0 {
      return i
    }
  }
  return 1
}

func EuclidGCD(pair1, pair2 int) int {
    var max,min int
    if pair1 > pair2 {
      max = pair1
      min = pair2
    } else {
      max = pair2
      min = pair1
    }
    var maxDivisor int
    complement := max % min
    if complement != 0 {
        maxDivisor = EuclidGCD(complement, min)
    } else {
        maxDivisor = min
    }
    return maxDivisor
}

// Exercise 3: Write and implement a function RelativelyPrimeProbability that takes three integers x, y, and numPairs as input, and returns an estimate of the probability that two randomly chosen numbers between x and y are relatively prime by selecting numPairs pairs of numbers between x and y, inclusively. Then, call your function on a very wide range of values with numPairs equal to 1 million. What is your conjectured value for the probability that two integers are relatively prime?

func RelativelyPrimeProbability(x, y, numPairs int) float64 {
  var posibility float64
  primeCount := 0
  numberList := make([]int, 0)
  for i := 0; i < 2*numPairs; i++ {
    rand.Seed(time.Now().UnixNano())
    numberList = append(numberList, int(math.Floor((rand.Float64()*float64((y-x))+float64(x)))))
  }
  for i := 0; i<len(numberList); i += 2 {
    if EuclidGCD(numberList[i],numberList[i+1]) == 1 {
      primeCount ++
    }
  }
  posibility = float64(primeCount)/float64(numPairs)
  return posibility
}

// Exercise 4: Write and implement a function HasRepeat that takes an array of integers as input; this function should return true if there is a repeated value and false otherwise.

func HasRepeat(list []int) bool {
  for i:= 0; i<len(list);i++ {
    for j:=i+1; j<len(list);j++{
      if list[i] == list[j] {
        return true
      }
    }
  }
  return false
}

// Exercise 5: Imagine a room of numPeople people, each with a birthday on one of the 365 days of the year (February 29th babies are forbidden).
//
// Write a function BirthdayParadox that takes two integers, numPeople and numTrials. It runs numTrials simulations and returns the average number of trials for which in a room of numPeople randomly generated people, at least one pair of people have the same birthday. (Hint: the month is irrelevant when running these simulations.)
//
// What is the smallest value of numPeople for which there is a greater than 50% chance of two people sharing the same birthday? Are you surprised?

func BirthdayParadox(numPeople, numTrials int) int {
  repeatCount := 0
  for n:=0; n<numTrials;n++ {
    birthdayList := make([]int, 0)
    for i:= 0; i<numPeople; i++ {
      rand.Seed(time.Now().UnixNano())
      birthdayList = append(birthdayList, rand.Intn(365))
    }
    if HasRepeat(birthdayList) == true {
      repeatCount ++
    }
  }
  return repeatCount
}

// Exercise 6: Write a function ComputePeriodLength that takes an array of integers as input. If there are no values that repeat, this function should return 0; if there is a repeated value in the input sequence, then the function should return the period of this sequence. Hint: if the sequence of numbers was (1473, 2856, 9830, 1789, 4468, 9830), then the length of the period would be 3 because the numbers that will repeat are 9830, 1789, and 4468.

func ComputePeriodLength (list []int) int {
  var length int
  if HasRepeat(list)!=true {
    length =0
  } else {
    for i:= 0; i<len(list);i++ {
      for j:=i+1; j<len(list);j++{
        if list[i] == list[j] {
          length = j-i
        }
      }
  }
}
  return length
}

// Exercise 7: We will need to ensure that numDigits is even. Write and implement a function CountNumDigits that takes an integer x as input and returns the number of digits in x. Your function should work for both positive and negative numbers. (Hint: first write your function for positive values of x, and then make a slight adjustment to it to accommodate negative values of x.)

func CountNumDigits (x int) int {
  var digitNum int
  stringx := strconv.Itoa(x)
  if x < 0 {
    digitNum = len(stringx)-1
  } else {
    digitNum = len(stringx)
  }
  return digitNum
}

// Exercise 8: Write and implement SquareMiddle. You should make sure to provide panic statements ensuring that numDigits is even, that both input parameters are positive, and that the number of digits in x is not greater than numDigits. You will also find it helpful to write a subroutine Pow10 that takes an integer n and returns 10n .

func SquareMiddle (x, numDigits int) int {
  y := x*x
  yLength := CountNumDigits(y)
  if numDigits % 2 != 0 {
    panic("numDigits should be even!")
  } else if x < 0 || numDigits < 0 {
    panic("Input parameters should be positive!")
  } else if numDigits > CountNumDigits(x) {
    panic("numDigits should not be bigger than the length of x!")
  } else {
    var middleNum int
    intermeidate := y % Pow10((yLength + numDigits)/2)
    middleNum = int(math.Floor(float64(intermeidate) / float64(Pow10((yLength - numDigits)/2))))
    return middleNum
  }
  return 0
}

func Pow10(n int) int{
  power:=1
  for i:=0; i<n;i++{
    power *= 10
  }
  return power
}

// GenerateMiddleSquareSequence(seed, numDigits)
//     seq ← sequence consisting of seed
//     while HasRepeat(seq) is false
//         seq ← append(seq, SquareMiddle(seed, numDigits))
//     return seq
// Exercise 9: Implement GenerateMiddleSquareSequence; test that it works on the example seeds 1600 and 3792 when numDigits is equal to 4. Then call GenerateMiddleSquareSequence on every four-digit seed between 1 and 9999. How many seeds produce a sequence of period 10 or smaller? Is the Middle-Square approach a good PRNG?

func GenerateMiddleSquareSequence(seed, numDigits int) []int {
  seq := make([]int,0)
  for HasRepeat(seq) == false {
    seq = append(seq, SquareMiddle(seed,numDigits))
  }
  return seq
}

func FourDigitCount() int {
  count := 0
  for i := 1000; i<=9999; i++ {
    list := GenerateMiddleSquareSequence(i,4)
    if ComputePeriodLength(list) <= 10 {
      count ++
    }
  }
  return count
}

// GenerateLinearCongruenceSequence(seed, a, c, m)
//     seq ← array of length 1
//     seq[1] ← seed
//     while HasRepeat(seq) is false
//         seed ← Remainder(a · x + c, m)
//         seq ← append(seq, seed)
//     return seq
// Exercise 10: Implement GenerateLinearCongruenceSequence. Then use these functions to answer the following two questions.
//
// Let m be the (Mersenne) prime number 8191, and consider c = 0 and any seed between 1 and m – 1 that you like. How many different values of a (between 1 and m – 1) produce a period having length equal to m – 1?
//
// Let a = 5, c = 1, and m = 213 = 8192. What is the period when the seed is equal to 1? What can we conclude about the period of every seed between 1 and m – 1?

func GenerateLinearCongruenceSequence(seed, a, c, m int) []int {
  seq := make([]int,1)
  seq[0] = seed
  for HasRepeat(seq) == false {
    seed = (a*seed+c)%m
    seq = append(seq,seed)
  }
  return seq
}
