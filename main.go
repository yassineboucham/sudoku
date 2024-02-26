package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func CheckSudokuIndex(args []string) bool {
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

func CheckSudokuStr(args []string) bool {
	for _, arg := range args {
		for i := 0; i < len(arg); i++ {
			if arg[i] == '.' {
				continue
			}
			for j := i + 1; j < len(arg); j++ {
				if arg[i] == '.' {
					continue
				}
				if arg[i] == arg[j] {
					return false
				}
			}
		}
	}
	return true
}

func SendArgs3(args [][]rune, x int, y int, nb rune) bool {
	x_index := x - x%3
	y_index := y - y%3
	for i := x_index; i < x_index+3; i++ {
		for j := y_index; j < y_index+3; j++ {
			if args[i][j] == nb {
				return false
			}
		}
	}
	return true
}

func check_nb_in_9number(nb []rune, n rune) bool {
	for i := 0; i < len(nb); i++ {
		if nb[i] == n {
			return false
		}
	}
	return true
}

// func CheckSudoku3(args []string) {
// 	var index int
// 	for xi := index; xi <= len(args)-1; xi += 3 {
// 		fmt.Println(xi)
// 		for xj := xi + 1; xj < xi+3; xj++ {
// 			fmt.Println(xj)
// 			for yi := 0; yi <= 3; yi++ {
// 				for yj := 0; yj < 3; yj++ {
// 					if args[xj][yj] == args[xi][xj] {
// 						fmt.Println(args[xj][yj], " = ", args[xi][xj])
// 					}
// 				}
// 			}
// 		}
// 		index = xi
// 	}
// }

func CheckSudoku(args []string) bool {
	if CheckSudokuStr(args) && CheckSudokuIndex(args) {
		return true
	}
	return false
}

func printsudoku(args []string) {
	for i := 1; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if j != len(args[i])-1 {
				z01.PrintRune(rune(args[i][j]))
				z01.PrintRune(' ')
			} else {
				z01.PrintRune(rune(args[i][j]))
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	// args := os.Args[1:]
	// fmt.Println("suduku")
	// printsudoku(args)
	// z01.PrintRune('\n')
	// fmt.Println("---------------------------------------------------------------")
	// z01.PrintRune('\n')
	// if CheckSudoku(args) {
	// 	fmt.Println("is good :)")
	// } else {
	// 	fmt.Println("is not good <_>")
	// }
	// fmt.Println("---------------------------------------------------------------")
	// // z01.PrintRune('\n')
	// // CheckSudoku3(args)

	sudoku := [][]rune{}
	sudoku = append(sudoku, {})
	sudoku[0] = append(sudoku[0], '3')
	sudoku[1] = append(sudoku[1], '4')
	// for i, v := range args {
	// 	for _, v2 := range v {

	// 	}
	// }

	fmt.Print(sudoku)
	fmt.Println()
}
