package main

import "math/rand"

func GenerateBoard(size int) GameBoard {
	var gb GameBoard
	gb.board = make([][]Square, size)
	for r := range gb.board {
		gb.board[r] = make([]Square, size)
		for c := range gb.board[r] {
			gb.board[r][c] = CreatSquare()
		}
	}
	return gb
}

func CreatSquare() Square {
	return Square{
		num: 0,
	}
}

func CircleBoard(gb GameBoard) GameBoard {
	var circled GameBoard
	circled.board = make([][]Square, len(gb.board)+2)
	for r := range circled.board {
		circled.board[r] = make([]Square, len(gb.board[0])+2)
		for c := range circled.board[r] {
			circled.board[r][c] = CreatSquare()
		}
	}
	for r := range gb.board {
		for c := range gb.board[r] {
			circled.board[r+1][c+1] = CopySuqare(gb.board[r][c])
		}
	}
	return circled
}

func DeCircle(gb GameBoard) GameBoard {
	var decircle GameBoard
	decircle.board = make([][]Square, len(gb.board)-2)
	for r := range decircle.board {
		decircle.board[r] = make([]Square, len(gb.board[0])-2)
		for c := range decircle.board[r] {
			decircle.board[r][c] = CopySuqare(gb.board[r+1][c+1])
		}
	}
	return decircle
}

func CopyGameBoard(gb GameBoard) GameBoard {
	var copyGB GameBoard
	copyGB.board = make([][]Square, len(gb.board))
	for r := range copyGB.board {
		copyGB.board[r] = make([]Square, len(gb.board[0]))
		for c := range copyGB.board[r] {
			copyGB.board[r][c] = CopySuqare(gb.board[r][c])
		}
	}
	return copyGB
}

func CopyRow(r []Square) []Square {
	copyR := make([]Square, len(r))
	for i := range copyR {
		copyR[i] = CopySuqare(r[i])
	}
	return copyR
}

func CopySuqare(s Square) Square {
	var copyS Square
	copyS.num = s.num
	return copyS
}

func PutCoins(gb GameBoard, n int, placement string) GameBoard {
	size := len(gb.board)
	if placement == "central" {
		cor := 0
		if size%2 == 0 {
			cor = size / 2
		} else if size%2 == 1 {
			cor = (size + 1) / 2
		}
		gb.board[cor][cor].num = n
	}

	if placement == "random" {
		positionList := make([][]int, 100)
		for i := range positionList {
			positionList[i] = make([]int, 2)
			positionList[i][0] = rand.Intn(len(gb.board))
			positionList[i][1] = rand.Intn(len(gb.board[0]))
		}
		for i := 0; i < n; i++ {
			pick := rand.Intn(100)
			gb.board[positionList[pick][0]][positionList[pick][1]].num++
		}
	}
	return gb
}
