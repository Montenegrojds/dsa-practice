/*
Valid Sudoku
Medium Topics Company Tags
Hints

You are given a 9 x 9 Sudoku board board. A Sudoku board is valid if the following rules are followed:

    Each row must contain the digits 1-9 without duplicates.
    Each column must contain the digits 1-9 without duplicates.
    Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.

Return true if the Sudoku board is valid, otherwise return false

Note: A board does not need to be full or be solvable to be valid.

Example 1:

Input: board =
[["1","2",".",".","3",".",".",".","."],
 ["4",".",".","5",".",".",".",".","."],
 [".","9","8",".",".",".",".",".","3"],
 ["5",".",".",".","6",".",".",".","4"],
 [".",".",".","8",".","3",".",".","5"],
 ["7",".",".",".","2",".",".",".","6"],
 [".",".",".",".",".",".","2",".","."],
 [".",".",".","4","1","9",".",".","8"],
 [".",".",".",".","8",".",".","7","9"]]

Output: true

Example 2:

Input: board =
[["1","2",".",".","3",".",".",".","."],
 ["4",".",".","5",".",".",".",".","."],
 [".","9","1",".",".",".",".",".","3"],
 ["5",".",".",".","6",".",".",".","4"],
 [".",".",".","8",".","3",".",".","5"],
 ["7",".",".",".","2",".",".",".","6"],
 [".",".",".",".",".",".","2",".","."],
 [".",".",".","4","1","9",".",".","8"],
 [".",".",".",".","8",".",".","7","9"]]

Output: false

Explanation: There are two 1's in the top-left 3x3 sub-box.

Constraints:

    board.length == 9
    board[i].length == 9
    board[i][j] is a digit 1-9 or '.'.

*/

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    squares := make(map[string]map[byte]bool)
	for i:=0;i<9;i++{
		rows[i]=make(map[byte]bool)
		cols[i]=make(map[byte]bool)
	}
	
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[i]);j++{
			if board[i][j]=='.'{
				continue
			}
			val:=board[i][j]
			key:=fmt.Sprintf("%d,%d",i/3,j/3)
			if rows[i][val]|| cols[j][val]|| squares[key][val]{
				return false
			}
			rows[i][val]=true
			cols[j][val]=true
			
			if squares[key]==nil{
				squares[key]= make(map[byte]bool)
			}
			squares[key][val]=true
		}
	}
	return true
}
