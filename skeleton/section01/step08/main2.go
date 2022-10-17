package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main2() {
	var n int
	fmt.Println("1:単発ガチャ 2:11連ガチャ 3:やめる")

LOOP:
	for {
		fmt.Println(">")
		var kind int
		fmt.Scanln(&kind)

		switch kind {
		case 1:
			n = 1
			break LOOP

		case 2:
			n = 11

		case 3:
			n = 0
		default:
			fmt.Println("もう一度入力してください")
		}
	}

	rand.Seed(time.Now().Unix())
	for i := 1; i <= n; i++ {
		num := rand.Intn(100)

		fmt.Printf("%d回目 ", i)
		switch {
		case num < 80:
			fmt.Println("normal")
		case num < 95:
			fmt.Println("rare")
		case num < 99:
			fmt.Println("super rare")
		default:
			fmt.Println("ultra rare")
		}
	}
}
