package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

/* ------- Check_Sudoku_Index --------*/
func CheckSudokuIndex(args [][]rune) bool {
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if args[i][j] == '.' {
				continue
			}
			for k := i + 1; k < len(args); k++ {
				if args[k][j] == '.' {
					continue
				}
				if args[i][j] == args[k][j] {
					return false
				}
			}
		}
	}
	return true
}

/* ------- Check_Sudoku_Str --------*/
func CheckSudokuStr(args [][]rune) bool {
	for i, arg := range args {
		for j := i + 1; j < len(arg); j++ {
			if arg[i] == '.' {
				continue
			}
			if arg[i] == arg[j] {
				return false
			}
		}
	}
	return true
}

/* ------- Number_Args_Reapet --------*/
func NumberArgsReapet(args [][]rune, x int, y int, nb rune) int {
	x_index := x - x%3
	y_index := y - y%3
	var rep int
	for i := x_index; i < x_index+3; i++ {
		for j := y_index; j < y_index+3; j++ {
			if args[i][j] == nb {
				rep++
			}
		}
	}
	return rep
}

/* ------- Check_InBox --------*/
func CheckInBox(args [][]rune) bool {
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if NumberArgsReapet(args, i, j, args[i][j]) == 1 {
				return true
			}
		}
	}
	return false
}

/* ------- Is_Number --------*/
func IsNumber(args [][]rune) bool {
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if args[i][j] == '.' || (args[i][j] >= '1' && args[i][j] <= '9') {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

/* ------- Check_Sudoku --------*/
func CheckSudoku(args [][]rune) bool {
	if CheckSudokuStr(args) && CheckSudokuIndex(args) && CheckInBox(args) && IsNumber(args) {
		return true
	}
	return false
}

/* ------- print_sudoku --------*/
func printsudoku(args [][]rune) {
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if j != len(args[i])-1 {
				z01.PrintRune(args[i][j])
				z01.PrintRune(' ')
			} else {
				z01.PrintRune(args[i][j])
			}
		}
		z01.PrintRune('\n')
	}
}

/* ------- SolveSudoku --------*/
func SolveSudoku(grid [][]rune) bool {
	emptyCell := findEmptyCell(grid)
	if emptyCell == nil {
		return true
	}

	row, col := emptyCell[0], emptyCell[1]

	for num := '1'; num <= '9'; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num

			if SolveSudoku(grid) {
				return true
			}

			grid[row][col] = '.'
		}
	}
	return false
}

func findEmptyCell(grid [][]rune) []int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' {
				return []int{i, j}
			}
		}
	}
	return nil
}

func isSafe(grid [][]rune, row, col int, num rune) bool {
	return !usedInRow(grid, row, num) &&
		!usedInCol(grid, col, num) &&
		!usedInBox(grid, row-row%3, col-col%3, num)
}

func usedInRow(grid [][]rune, row int, num rune) bool {
	for i := 0; i < len(grid[row]); i++ {
		if grid[row][i] == num {
			return true
		}
	}
	return false
}

func usedInCol(grid [][]rune, col int, num rune) bool {
	for i := 0; i < len(grid); i++ {
		if grid[i][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(grid [][]rune, boxStartRow, boxStartCol int, num rune) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+boxStartRow][j+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}

/* ------- main --------*/
func main() {
	args := os.Args[1:]

	grid := make([][]rune, 9)
	for i := range grid {
		grid[i] = make([]rune, 9)
		for j, char := range args[i] {
			grid[i][j] = char
		}
	}
	if CheckSudoku(grid) {
		if SolveSudoku(grid) {
			printsudoku(grid)
		} else {
			fmt.Println("Error")
		}
	}
}
