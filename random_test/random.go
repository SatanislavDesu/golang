package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	reader := bufio.NewReader(os.Stdin)

	// Запрос минимального значения
	fmt.Print("Введите минимальное значение: ")
	minInput, _ := reader.ReadString('\n')
	min, err := strconv.Atoi(minInput[:len(minInput)-1]) // Удаляем символ новой строки
	if err != nil {
		fmt.Println("Ошибка ввода минимального значения:", err)
		return
	}

	// Запрос максимального значения
	fmt.Print("Введите максимальное значение: ")
	maxInput, _ := reader.ReadString('\n')
	max, err := strconv.Atoi(maxInput[:len(maxInput)-1]) // Удаляем символ новой строки
	if err != nil {
		fmt.Println("Ошибка ввода максимального значения:", err)
		return
	}

	// Проверка диапазона
	if min >= max {
		fmt.Println("Минимальное значение должно быть меньше максимального")
		return
	}

	// Запрос количества чисел
	fmt.Print("Сколько случайных чисел сгенерировать? ")
	countInput, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(countInput[:len(countInput)-1]) // Удаляем символ новой строки
	if err != nil {
		fmt.Println("Ошибка ввода количества чисел:", err)
		return
	}

	// Генерация и вывод случайных чисел
	fmt.Printf("\nСлучайные числа в диапазоне от %d до %d:\n", min, max)
	for i := 0; i < count; i++ {
		num := rand.Intn(max-min+1) + min
		fmt.Println(num)
	}
}
