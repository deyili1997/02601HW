package main

type GameBoard struct {
	board [][]Square
}
type Square struct {
	num int
}

//The highest func of serial computing
func SerialComputing(gb GameBoard) GameBoard {
	boardForSerial := CopyGameBoard(gb)

	init := CircleBoard(boardForSerial)

	circledFinished := Topple(init)

	Finished := DeCircle(circledFinished)
	return Finished
}

func Topple(gb GameBoard) GameBoard {
	curBoard := gb
	Update(curBoard)
	return curBoard
}

func Update(gb GameBoard) {
	stable := true
	for r := 1; r <= len(gb.board)-2; r++ {
		for c := 1; c <= len(gb.board[0])-2; c++ {
			if gb.board[r][c].num >= 4 {
				stable = false
				gb.board[r][c].num -= 4
				gb.board[r-1][c].num++
				gb.board[r+1][c].num++
				gb.board[r][c-1].num++
				gb.board[r][c+1].num++
			}
		}
	}
	if !stable {
		Update(gb)
	}
}

func CheckStable(gb GameBoard) bool {
	stable := true
	for r := 1; r <= len(gb.board)-2; r++ {
		for c := 1; c <= len(gb.board[0])-2; c++ {
			if gb.board[r][c].num >= 4 {
				stable = false
			}
		}
	}
	return stable
}
