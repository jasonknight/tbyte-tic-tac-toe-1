package main

import "testing"
import "fmt"
func TestCreateBoard(t *testing.T) {
	size := 3
	board := createBoard(size)
	for i:=0; i < size; i++ {
		for j:=0; j<size; j++ {
			if board[i][j].Owner != UNOWNED {
				t.Errorf("owner should be -2")
			}
		}
	}
	displayBoard(board)
	// Output:
	// * * *
  // * * *
  // * * *
}
func TestUnplayed(t *testing.T) {
	size := 3
	board := createBoard(size)
	up := unplayedSquares(board)
	if ( len(up) != (size*size) ) {
		t.Errorf("expected %d got %d",(size*size),len(up))
	}
	board = play(board,0,1,2)
	up = unplayedSquares(board)
	if ( len(up) != (size*size) - 1 ) {
		t.Errorf("expected %d got %d",(size*size) - 1,len(up))
	}
	play(board,1,0,0)
	displayBoard(board)
}
func TestWonHorizontal(t *testing.T) {
	size := 3
	board := createBoard(size)
	for i:= 0; i < size; i++ {
		board = play(board,1,0,i)
	}
	displayBoard(board)
	isWon,owner := boardIsWon(board)
	if isWon != true {
		t.Errorf("Expected board to be won!")
	}
	if owner != 1 {
		t.Errorf("expected owner to be 1")
	}
	board = createBoard(size)
	board = play(board,1,0,0)
	board = play(board,1,1,1)
	board = play(board,1,2,2)
	displayBoard(board)
	isWon,owner = boardIsWon(board)
	if isWon != true {
		t.Errorf("Expected board to be won on diag")
	}
	if owner != 1 {
		t.Errorf("expected owner to be 1 on diag")
	}
	board = createBoard(size)
	board =play(board,0,0,0)
	board =play(board,0,1,1)
	board =play(board,0,2,2)
	displayBoard(board)
	isWon,owner = boardIsWon(board)
	if isWon != true {
		t.Errorf("Expected board to be won on diag")
	}
	if owner != 0 {
		t.Errorf("expected owner to be 0 on diag")
	}
	board = createBoard(size)
	board = play(board,0,0,2)
	board = play(board,0,1,1)
	board = play(board,0,2,0)
	displayBoard(board)
	isWon,owner = boardIsWon(board)
	if isWon != true {
		t.Errorf("Expected board to be won on diag")
	}
	if owner != 0 {
		t.Errorf("expected owner to be 0 on diag")
	}
}
func TestAIPlay(t *testing.T) {
	fmt.Printf("TestingAIPlay\n")
	size := 4
	board := createBoard(size)
	board = play(board,1,0,0)
	board = play(board,1,0,1)
	board = play(board,1,0,2)
	board = aiPlay(board)
	displayBoard(board)
	isWon,owner := boardIsWon(board)
	if isWon != true && owner != 1 {
		t.Error("expected board to be won")
	}
	board = createBoard(size)
	board = play(board,1,2,0)
	board = play(board,1,2,2)
	board = play(board,1,2,1)
	board = aiPlay(board)
	displayBoard(board)
	isWon,owner = boardIsWon(board)
	if isWon != true && owner != 1 {
		t.Error("expected board to be won")
	}
	// Left diagonal
	board = createBoard(size)
	board = play(board,1,0,0)
	board = play(board,1,2,2)
	board = play(board,1,3,3)
	board = aiPlay(board)
	displayBoard(board)
	isWon,owner = boardIsWon(board)
	if isWon != true && owner != 1 {
		t.Error("expected board to be won")
	}
	board = createBoard(size)
	board = play(board,1,1,1)
	board = play(board,1,2,2)
	board = play(board,1,3,3)
	aiPlay(board)
	displayBoard(board)
	isWon,owner = boardIsWon(board)
	if isWon != true && owner != 1 {
		t.Error("expected board to be won")
	}
}
