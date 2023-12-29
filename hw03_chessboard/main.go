package main

import "fmt"

func main() {
	fmt.Println(`Генератор шахматного поля приветствует Вас!`)
	var rows, cols int
	fmt.Print(`Введите размер поля в формате 'axb',
где 'a' - это количество строк,
а 'b' - количество столбцов: `)
	fmt.Scanf("%dx%d", &rows, &cols)

	fmt.Print(" ")
	for a := 0; a < cols; a++ {
		fmt.Print("---")
	}
	fmt.Println()

	for a := 0; a < rows; a++ {
		fmt.Print("|")
		for b := 0; b < cols; b++ {
			if (a+b)%2 == 0 {
				fmt.Print("   ")
			} else {
				fmt.Print(" # ")
			}
		}
		fmt.Println("|")
	}

	fmt.Print(" ")
	for a := 0; a < cols; a++ {
		fmt.Print("---")
	}
}
