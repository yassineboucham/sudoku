package main

import (
	"github.com/01-edu/z01"
	"os"
)

func CheckSudokuArr(arg string) bool {
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
			z01.PrintRune(rune(args[i][j]))
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func main(){
	args := os.Args
	if CheckSudokuArr(args[1]) {
		z01.PrintRune('o')
	} else {
		z01.PrintRune('n')
	}
}