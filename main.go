package main

import (
	"os"

	"github.com/01-edu/z01"
)

func CheckSudokuIndex(args []string) bool {
	for i := 1; i < len(args); i++ {
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

func CheckSudokuStr(arg string) bool {
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
	return true
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
	args := os.Args
	printsudoku(args)
	z01.PrintRune('\n')
	z01.PrintRune('\n')
}
