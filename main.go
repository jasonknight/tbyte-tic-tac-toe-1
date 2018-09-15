package main

import (
	//"github.com/jasonknight/toolbox"
	"fmt"
)

type Square struct {
	Display string
	Owner int
	X,Y int
}
const UNOWNED = -2
type Row []Square
type Board []Row
func createBoard(size int) Board {
	var grid Board
	for i:=0; i < size; i++ {
		cols := make([]Square,size)
		for j:=0; j<size; j++ {
			cols[j].Display = fmt.Sprintf("%2s","*")
			cols[j].Owner = UNOWNED
			cols[j].X = j
			cols[j].Y = i
		}
		grid = append(grid,cols)
	}
	return grid
}
func unplayedSquares(board Board) Row {
	size := len(board)
	var results []Square
	for i:=0; i < size; i++ {
		for j:=0; j<size; j++ {
			col := board[i][j]
			if col.Owner == UNOWNED {
				results = append(results,col)
			}
		}
	}
	return results
}
func displayBoard(board Board) {
	size := len(board)
	for k:=0;k<size; k++ {
		fmt.Print("--")
	}
	fmt.Print("-\n")
	for i:=0; i < size; i++ {
		for j:=0; j<size; j++ {
			fmt.Print(board[i][j].Display)
		}
		fmt.Print("\n")
	}
	for k:=0;k<size; k++ {
		fmt.Print("--")
	}
	fmt.Print("-\n")
}
func play(board Board, owner, y, x int) Board {
	sq := board[y][x]
	sq.Owner = owner
	if ( owner == 0 ) {
		sq.Display = fmt.Sprintf("%2s","0")
	} 
	if owner == 1 {
		sq.Display = fmt.Sprintf("%2s","X")
	}
	board[y][x] = sq
	return board
}
func sumRow(board Board, y int) int {
	sum := 0
	row := board[y]
	for i:=0;i<len(row);i++ {
		sum += row[i].Owner
	}
	return sum
}
func sumCol(board Board,y,x int) int {
	sum := 0
	for i:=y; i<len(board);i++ {
		sum += board[i][x].Owner
	}
	return sum
}
type DiagStepFunction func (b Board, sx, sy int) bool
func sumLeftDiag(board Board) int {
	sum := 0
	x := 0
	for y := 0; y < len(board); y++ {
		sum += board[y][x].Owner	
		x++
	}
	return sum
}
func sumRightDiag(board Board) int {
	sum := 0
	x := len(board) - 1
	for y := 0; y < len(board); y++ {
		sum += board[y][x].Owner
		x--
	}
	return sum
}
func boardIsWon(board Board) (bool,int) {
	//fmt.Printf("Checking by row\n")
	for i:=0; i<len(board); i++ {
		s := sumRow(board,i)
		if s == len(board[0]) {
			return true,1
		}
		if s == 0 {
			return true,0
		}
	}
	//fmt.Printf("Checking by col\n")
	for i:= 0; i<len(board[0]); i++ {
		s := sumCol(board,0,i)
		if s == len(board[0]) {
			return true,1
		}
		if s == 0 {
			return true,0
		}
	}
	//fmt.Printf("Checking by left diag\n")
	s := sumLeftDiag(board)	
	//fmt.Printf("%d %d\n",s,len(board))
	if s == len(board[0]) {
		return true,1
	}
	if s == 0 {
		return true,0
	}
	//fmt.Printf("Checking by right diag\n")
	s = sumRightDiag(board) 
	if s == len(board[0]) {
		return true,1
	}
	if s == 0 {
		return true,0
	}
	//fmt.Print("returning false\n")
	return false, UNOWNED
}
func getSquaresByOwner(board Board, owner, y int) Row {
	var r Row
	for x := 0; x < len(board[y]); x++ {
		if board[y][x].Owner == owner {
			r = append(r,board[y][x])
		}
	}
	return r
}
func getLeftDiagSquaresByOwner(board Board, owner int) Row {
	var r Row
	x := 0
	for y := 0; y < len(board); y++ {
		if board[y][x].Owner == owner {
			r = append(r,board[y][x])
		}
		x++
	}
	return r
}
func getRightDiagSquaresByOwner(board Board, owner int) Row {
	var r Row
	x := len(board) - 1
	for y := 0; y < len(board); y++ {
		if board[y][x].Owner == owner {
			r = append(r,board[y][x])
		}
		x--
	}
	return r
}

func aiPlay(board Board) Board {
	// First we need to see if we have
	// an opportunity to win
	for y := 0; y < len(board); y++ {
		owned := getSquaresByOwner(board,1,y)
		if len(owned) > 0 && len(owned) >= (len(board)/2) {
			for x := 0; x < len(board); x++ {
				if board[y][x].Owner == UNOWNED {
					return play(board,1,y,x)
				}
			}
		}
	}
	// Try the same with the diagonals
	owned := getLeftDiagSquaresByOwner(board,1)
	if len(owned) > 0 && len(owned) >= (len(board)/2) {
		x := 0
		for y := 0; y < len(board); y++ {
			if board[y][x].Owner == UNOWNED {
				return play(board,1,y,x)
			}
			x++
		}
	}
	owned = getRightDiagSquaresByOwner(board,1)
	if len(owned) > 0 && len(owned) >= (len(board)/2) {
		x := len(board) - 1
		for y := 0; y < len(board); y++ {
			if board[y][x].Owner == UNOWNED {
				return play(board,1,y,x)
			}
			x--
		}
	}
	return board
}
func main() {

}
