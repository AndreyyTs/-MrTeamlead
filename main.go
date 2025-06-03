//go:build amd64

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 100_000_000

type Foo struct { // size=16 (0x10)
	a uint64
	b uint16
	c uint16
}

func main() {
	// Инициализируем генератор случайных чисел
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	v := make([]Foo, N)
	for i := range v {
		v[i] = Foo{
			a: uint64(rng.Intn(65536)), // 0-255
			b: uint16(rng.Intn(65536)), // случайное uint64
			c: uint16(rng.Intn(65536)), // 0-65535
		}
	}

	{
		// // Выводим по 4 элемента в строку
		// fmt.Println("[")
		// for i := 0; i < len(v); i += 4 {
		// 	end := i + 4
		// 	if end > len(v) {
		// 		end = len(v)
		// 	}

		// 	// Формируем строку с элементами
		// 	var line string
		// 	for j := i; j < end; j++ {
		// 		if j > i {
		// 			line += ", "
		// 		}
		// 		line += fmt.Sprintf("{a:%d b:%d c:%d}", v[j].a, v[j].b, v[j].c)
		// 	}

		// 	// Добавляем запятую, если это не последняя строка
		// 	if end < len(v) {
		// 		line += ","
		// 	}

		// 	fmt.Println("   ", line)
		// }
		// fmt.Println("]")
	}

	var start time.Time
	var dur time.Duration
	var sum uint64

	start = time.Now()
	sum = Sum4(v)
	dur = time.Since(start)
	fmt.Println("sum4 =", sum)
	fmt.Println("dur", dur)
	fmt.Println("")

	start = time.Now()
	sum = SumSIMD(v)
	dur = time.Since(start)
	fmt.Println("SumSIMD =", sum)
	fmt.Println("dur", dur)
	fmt.Println("")

	start = time.Now()
	sum = Sum(v)
	dur = time.Since(start)
	fmt.Println("sum =", sum)
	fmt.Println("dur", dur)
	fmt.Println("")

	start = time.Now()
	sum = Sum4(v)
	dur = time.Since(start)
	fmt.Println("sum4 =", sum)
	fmt.Println("dur", dur)
	fmt.Println("")

	start = time.Now()
	sum = Sum(v)
	dur = time.Since(start)
	fmt.Println("sum =", sum)
	fmt.Println("dur", dur)
	fmt.Println("")

	start = time.Now()
	sum = SumSIMD(v)
	dur = time.Since(start)
	fmt.Println("SumSIMD =", sum)
	fmt.Println("dur", dur)
	fmt.Println("")

	start = time.Now()
	sum = Sum4(v)
	dur = time.Since(start)
	fmt.Println("sum4 =", sum)
	fmt.Println("dur", dur)
	fmt.Println("")

	start = time.Now()
	sum = Sum(v)
	dur = time.Since(start)
	fmt.Println("sum =", sum)
	fmt.Println("dur", dur)
	fmt.Println("")

}

func Sum(arr []Foo) (sum uint64) {
	for i := range arr {
		sum += arr[i].a
	}
	return sum
}

func Sum4(arr []Foo) (sum uint64) {
	for i := 0; i < len(arr); i = i + 4 {
		sum += arr[i+0].a
		sum += arr[i+1].a
		sum += arr[i+2].a
		sum += arr[i+3].a

	}
	return sum
}

// SIMD NEON версия - реализация будет в .s файлах или fallback
//
//go:noescape
func SumSIMD(arr []Foo) uint64
