package hard

/*
 https://leetcode.com/problems/sudoku-solver/description/

Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
Empty cells are indicated by the character '.'.

Note:

The given board contain only digits 1-9 and the character '.'.
You may assume that the given Sudoku puzzle will have a single unique solution.
The given board size is always 9x9.

*/

func SolveSudoku(board [][]byte) {
	dfs(board, 0)
}

func dfs(board [][]byte, d int) bool {
	if d == 81 {
		return true
	}

	row := d / 9
	col := d % 9
	if board[row][col] != '.' {
		return dfs(board, d + 1)
	}
	var valArr []bool = make([]bool, 10)
	validate(board, row, col, valArr)
	for i := 1; i <= 9; i++  {
		if !valArr[i] {
			board[row][col] = byte('0' + i)
			if dfs(board, d + 1) {
				return true
			}
		}
	}
	board[row][col] = '.'
	return false
}

/*
对于solveSudoku1的优化
同样是检查 行和列 以及 小的9宫格里是否有可填数字
 */
func validate(board [][]byte, row, col int, valArr []bool) {

	for i := 0; i < 9; i ++ {
		//当前元素所在竖 行的所有元素遍历
		if board[i][col] != '.' {
			valArr[board[i][col] - '0'] = true
		}

		if board[row][i] != byte('.') {
			valArr[board[row][i] - '0'] = true
		}

		//当前元素所在九宫格元素遍历
		/*
		9宫格 行计算说明：
		传入的是一个外围9 * 9的数组行号
		通过 /3 之后能够确定是从那个小9宫格开始，*3 得出的是当前小9宫格的开始行
		然后通过 循环次数确定在当前小9宫格中的行移动
		  */
		r99 := row / 3 * 3 + i / 3

		/*
		9宫格 行计算说明：
		传入的是一个外围9 * 9的数组行号
		通过 /3 之后能够确定是从那个小9宫格开始，*3 得出的是当前小9宫格的开始列
		然后通过 循环次数确定在当前小9宫格中的行移动
		  */
		c99 := col / 3 * 3 + i % 3

		if board[r99][c99] != byte('.') {
			valArr[board[r99][c99] - '0'] = true
		}
	}
}


func solveSudoku1(board [][]byte, row, col int) {

	var valMap map[byte]byte

	for i := 0; i < 9; i++ {
		if board[row][i] != byte('.') {
			valMap[board[row][i]] = board[row][i]
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][col] != byte('.') {
			valMap[board[row][i]] = board[row][i]
		}
	}

	rowcount := (row / 3 + 1) * 3
	colcount := (col / 3 + 1) * 3

	for i := rowcount - 3; i < rowcount; i++ {
		for j := colcount - 3; j < colcount; j++ {
			if board[i][j] != byte('.') {
				valMap[board[row][i]] = board[row][i]
			}
		}
	}

	if len(valMap) == 9 {
		for i := 1; i < 10; i++ {
			if _, ok := valMap[byte(i)]; !ok {
				board[row][col] = byte(i)
			}
		}
		return
	}
}
















