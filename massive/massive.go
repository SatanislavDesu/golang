package main

import "fmt"

// func main() {
// 	var x = [...]int{1, 2, 3} //это массив
//  var x =[]int{1, 2, 2} // а это срез
// 	var y = [3]int{1, 2, 3}
// 	fmt.Println(x == y)
// }

// func main() {
// 	var x = [10]int{}
// 	fmt.Println(len(x))
// }
// len возвращает количество элементов в массиве

//

// func main() {
// 	x := []int{1, 2, 3}
// 	y := []int{1, 2, 3}
// 	s := []int{1, 2, 3, 4, 5, 6}
// 	z := []string{"a", "b", "c"}

// 	fmt.Println(slices.Equal(x, y)) // true
// 	fmt.Println(slices.Equal(x, s)) // false
// 	fmt.Println(slices.Equal(x, z)) // нельзя сравнивать срезы разных типов

// }
// slices.Equal проверяет, равны ли два среза
// они должны быть одного типа и одинаковой длины

// func main() {
// 	var x = []int{1, 2, 8: 100, 44: 1} // срез с пропусками
// 	x = append(x, 9)                   // добавляем элемент в конец среза

// 	fmt.Println(len(x)) // длина среза
// 	fmt.Println(cap(x)) // ёмкость среза
// 	fmt.Println(x)      // вывод среза

// }

// append добавляет элементы в срез
// если срез не может вместить новые элементы, он будет расширен
// при этом может быть выделена новая память
// если срез был расширен, то старый срез будет недействителен

// func main() {
// 	x := make([]int, 10) // создаем срез длиной 10
// 	x = append(x, 10)    // добавляем элементы в конец среза
// 	fmt.Println(len(x), cap(x), x)

// 	y := make([]int, 20, 100) // создаем срез длиной 20 и ёмкостью 100
// 	y = append(y, 10)         // добавляем элементы в конец среза
// 	fmt.Println(len(y), cap(y), y)
// }

// make создает срез с заданной длиной и ёмкостью

func main() {
	fmt.Println("Hello, World")
}
