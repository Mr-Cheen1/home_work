package main

import "fmt"

func generateChessboard(rows, cols int) string {
	if rows <= 0 || cols <= 0 {
		return ""
	}

	var chessboard string
	chessboard += " "
	for a := 0; a < cols; a++ {
		chessboard += "---"
	}
	chessboard += "\n"

	for a := 0; a < rows; a++ {
		chessboard += "|"
		for b := 0; b < cols; b++ {
			if (a+b)%2 == 0 {
				chessboard += "   "
			} else {
				chessboard += " # "
			}
		}
		chessboard += "|\n"
	}

	chessboard += " "
	for a := 0; a < cols; a++ {
		chessboard += "---"
	}
	return chessboard
}

func main() {
	fmt.Println(`Генератор шахматного поля приветствует Вас!`)
	var rows, cols int
	fmt.Print(`Введите размер поля в формате 'axb',
где 'a' - это количество строк,
а 'b' - количество столбцов: `)
	fmt.Scanf("%dx%d", &rows, &cols)

	fmt.Println(generateChessboard(rows, cols))
}
