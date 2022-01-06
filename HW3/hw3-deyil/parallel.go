package main

import "runtime"

//The highest func of parallel computing
func ParallelComputing(gb GameBoard, placement string) GameBoard {
	BoardForParal := CopyGameBoard(gb)
	maxRound := FindMaxRound(BoardForParal, placement)
	numProcs := runtime.NumCPU()
	ChannelList := make([]chan []Square, (numProcs-1)*2)
	for i := range ChannelList {
		ChannelList[i] = make(chan []Square, 1)
	}
	resultChannelList := make([]chan GameBoard, numProcs)
	for i := range resultChannelList {
		resultChannelList[i] = make(chan GameBoard)
	}
	subBoardList := make([]GameBoard, numProcs)
	for i := range subBoardList {
		startRow := i * (len(BoardForParal.board) / numProcs)
		endRow := (i + 1) * (len(BoardForParal.board) / numProcs)
		var subZone GameBoard
		if i < numProcs-1 {
			subZone.board = BoardForParal.board[startRow:endRow]
		} else {
			subZone.board = BoardForParal.board[startRow:]
		}
		copy := CopyGameBoard(subZone)
		finish := CircleBoard(copy)
		subBoardList[i] = finish
	}
	for i := 0; i < numProcs; i++ {
		if i == 0 {
			go DoUpSuqare(subBoardList[i], ChannelList[2*i], ChannelList[2*i+1], resultChannelList[i], maxRound)
		} else if i == numProcs-1 {
			go DoDownSuqare(subBoardList[i], ChannelList[2*i-1], ChannelList[2*i-2], resultChannelList[i], maxRound)
		} else {
			go DoMidSquare(subBoardList[i], ChannelList[2*i-1], ChannelList[2*i-2], ChannelList[2*i], ChannelList[2*i+1], resultChannelList[i], maxRound)
		}
	}
	var ParralFinal GameBoard
	for _, j := range resultChannelList {
		ParralFinal.board = append(ParralFinal.board, DeCircle(<-j).board...)
	}
	return ParralFinal
}

func FindMaxRound(gb GameBoard, placement string) int {
	biggest := 0
	for r := range gb.board {
		for c := range gb.board {
			if gb.board[r][c].num > biggest {
				biggest = gb.board[r][c].num
			}
		}
	}
	if placement == "central" {
		return biggest/4 + 1
	}
	return biggest
}

//Discuss in 3 situations up, mid and bottom
func DoUpSuqare(gb GameBoard, downSend, downReceive chan []Square, result1 chan GameBoard, maxRound int) {
	var border []Square
	round := 0
	currBoard := gb
	for round <= maxRound {
		currBoard = Topple(currBoard)
		downSendRow := CopyRow(currBoard.board[len(currBoard.board)-1])
		downSend <- downSendRow
		for i := range currBoard.board[len(currBoard.board)-1] {
			currBoard.board[len(currBoard.board)-1][i] = CreatSquare()
		}

		border = <-downReceive

		for i := 1; i < len(border)-1; i++ {
			currBoard.board[len(currBoard.board)-2][i].num += border[i].num
		}
		round++
	}
	result1 <- currBoard
}

func DoMidSquare(gb GameBoard, upSend, upReceive, downSend, downReceive chan []Square, result2 chan GameBoard, maxRound int) {
	var upBorder []Square
	var downBorder []Square
	round := 0
	currBoard := gb
	for round <= maxRound {
		currBoard = Topple(currBoard)
		upSendRow := CopyRow(currBoard.board[0])
		upSend <- upSendRow
		for i := range currBoard.board[0] {
			currBoard.board[0][i] = CreatSquare()
		}

		upBorder = <-upReceive

		for i := 1; i < len(upBorder)-1; i++ {
			currBoard.board[1][i].num += upBorder[i].num
		}

		downSendRow := CopyRow(currBoard.board[len(gb.board)-1])
		downSend <- downSendRow
		for i := range currBoard.board[len(gb.board)-1] {
			currBoard.board[len(gb.board)-1][i] = CreatSquare()
		}

		downBorder = <-downReceive
		for i := 1; i < len(downBorder)-1; i++ {
			currBoard.board[len(gb.board)-2][i].num += downBorder[i].num
		}
		round++
	}
	result2 <- currBoard
}

func DoDownSuqare(gb GameBoard, upSend, upReceive chan []Square, result3 chan GameBoard, maxRound int) {
	var border []Square
	currBoard := gb
	round := 0
	for round <= maxRound {
		currBoard = Topple(currBoard)
		upSendRow := CopyRow(currBoard.board[0])
		upSend <- upSendRow
		for i := range currBoard.board[0] {
			currBoard.board[0][i] = CreatSquare()
		}

		border = <-upReceive

		for i := 1; i < len(border)-1; i++ {
			currBoard.board[1][i].num += border[i].num
		}
		round++
	}
	result3 <- currBoard
}
