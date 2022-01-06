package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

// The data stored in a single cell of a field
type Cell struct {
	strategy  string //represents "C" or "D" corresponding to the type of prisoner in the cell
	score float64 //represents the score of the cell based on the prisoner's relationship with neighboring cells
}

// The game board is a 2D slice of Cell objects
type GameBoard [][]Cell


// Main function: takes the filename, the reward value of being a defector and how many times to run the simulation. It returns an array which takes every round gameboard as the output.
func RunSimulation(fileName string, b float64, steps int) []GameBoard {
	simulationResults := make([]GameBoard,0)
	var board GameBoard
	var rowNum, colNum int
	board, rowNum, colNum = ParseFile(fileName)
	simulationResults = append(simulationResults,board)
	for i:=0; i < steps; i++ {
		board = CalculateScore(board,b,rowNum,colNum)
		board = NextGenBoard(board,rowNum,colNum)
		simulationResults = append(simulationResults,board)
	}
	for j := 0; j < len(simulationResults); j++ {
		simulationResults[j] = DropAssistance(simulationResults[j], rowNum, colNum)
	}
	return simulationResults
}


// Take a string of the name of the file you want to parse and put each "strategy" value into a gameboard and return the gameboard
func ParseFile (fileName string) (GameBoard,int,int) {
// Get row number and column number
	file := ReadFile(fileName)
	scanner := bufio.NewScanner(file)
	currentRow := 0
	var rowNum, colNum int
	var rowVsCol []string
	var initBoard GameBoard
 	for scanner.Scan(){
		currentRow ++
		if currentRow == 1 {
			var firstRow string = scanner.Text()
			rowVsCol = strings.Split(firstRow," ")
			int, err1 := strconv.Atoi(rowVsCol[0])
			int2, err2 := strconv.Atoi(rowVsCol[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Something wrong with string-int conversion!")
			}
			rowNum = int+2
			colNum = int2+2
			// Creat a 2-D slice for the gameboard according to row number and column number and surround the gameboard with "D" cells
			initBoard = InitializeGameboard (rowNum, colNum)
		}
		// Extract "strategy" values from the txt and put them into the initialized gameboard. 0 is a default value for "score".
		if currentRow >= 2 {
			for currentCol := 1; currentCol < colNum-1; currentCol++ {
				initBoard[currentRow-1][currentCol].strategy = string(scanner.Text()[currentCol-1])
			}
		}
	}
// report reading error
	if scanner.Err() != nil {
		fmt.Println("Error: there was a problem reading the file")
		os.Exit(1)
		}
	file.Close()
	return initBoard, rowNum, colNum
}


// Creat a 2-D slice for the gameboard according to row number and column number and surround the gameboard with "D" cells
func InitializeGameboard(rowNum, colNum int) GameBoard {
	var initBoard GameBoard
	initBoard = make([][]Cell, rowNum)
	for i:= 0; i < rowNum; i++ {
		initBoard[i] = make([]Cell,colNum)
	}
// creat boundary "D" cells
	for currentCol1 := 0; currentCol1 < colNum; currentCol1++ {
		initBoard[0][currentCol1].strategy = "D"
		initBoard[rowNum-1][currentCol1].strategy = "D"
	}
	for currentRow1 :=0; currentRow1 < rowNum; currentRow1++ {
		initBoard[currentRow1][0].strategy = "D"
		initBoard[currentRow1][colNum-1].strategy = "D"
	}
	return initBoard
}



// take a string of the name of the file you want to load in and return *os.File
func ReadFile (fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: Something went wrong when opening the file!")
		fmt.Println("Error: Probably you gave the wrong filename!")
	}
	return file
}

//CalculateScore fucntion takes board, the reward value of being a defector and the number of rows and columns as input and returns the board with "score" calulated.
func CalculateScore(board GameBoard,b float64, rowNum, colNum int) GameBoard {
//calulate score for inner cells
	for currentRow:= 1; currentRow < rowNum-1; currentRow++ {
		for currentCol:=1; currentCol < colNum-1; currentCol++{
			if board[currentRow][currentCol].strategy == "C"{
				board[currentRow][currentCol].score = float64(CountCNearby(board, currentRow, currentCol))
			} else {
				board[currentRow][currentCol].score = b * float64(CountCNearby(board, currentRow, currentCol))
			}
		}
	}
	return board
}

//Take gameboard,current row and column as input, return the number of "C" cells in eight neighbourhood
func CountCNearby (board GameBoard, currentRow,currentCol int) int {
	var countC int = 0
	for i:= currentCol-1; i <= currentCol+1; i++{
		if board[currentRow-1][i].strategy == "C" {
			countC ++
		}
		if board[currentRow+1][i].strategy == "C" {
			countC ++
		}
	}
	if board[currentRow][currentCol-1].strategy == "C" {
		countC++
	}
	if board[currentRow][currentCol+1].strategy == "C" {
		countC++
	}
	return countC
}

//take the number of columns and row as input
//generate the next generation gameboard
func NextGenBoard(board GameBoard,rowNum,colNum int) GameBoard {
	nextBoard := InitializeGameboard(rowNum,colNum)
	for currentRow := 1; currentRow < rowNum - 1; currentRow ++ {
		for currentCol := 1; currentCol < colNum -1; currentCol ++ {
			nextBoard[currentRow][currentCol].strategy = GetMaxStrategy(board, currentRow, currentCol)
			}
		}
	return nextBoard
}

//take the gameboard, the cell we are checking and return the neighbour's best strategy
func GetMaxStrategy(board GameBoard, currentRow, currentCol int) string {
	var maxNeighbourStrategy string
	max := board[currentRow-1][currentCol-1].score
	maxPosition := []int {currentRow-1,currentCol-1}
	for i:= currentRow-1; i <= currentRow+1; i++ {
		for j:= currentCol-1; j <= currentCol+1; j++ {
			if board[i][j].score > max {
				max = board[i][j].score
				maxPosition = []int {i,j}
			}
		}
	}
	maxNeighbourStrategy = board[maxPosition[0]][maxPosition[1]].strategy
	return maxNeighbourStrategy
}

// take the gameboard surrounded by "D" prisoners (as the assistance) and drop the surrounding
func DropAssistance (board GameBoard, rowNum, colNum int) GameBoard {
	var finalBoard GameBoard
	finalBoard = make([][]Cell, rowNum-2)
	for i := 0; i < rowNum -2 ; i++ {
		finalBoard[i] = make([]Cell, colNum-2)
	}
	for currentRow := 1; currentRow < rowNum - 1; currentRow ++ {
		for currentCol := 1; currentCol < colNum -1; currentCol ++ {
			finalBoard[currentRow-1][currentCol-1].strategy = board[currentRow][currentCol].strategy
			finalBoard[currentRow-1][currentCol-1].score = board[currentRow][currentCol].score
		}
	}
	return finalBoard
}
