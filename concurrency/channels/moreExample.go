package main

import "fmt"

func digit(number int, digitCh chan int) {
	for number != 0 {
		digit := number % 10
		digitCh <- digit
		number /= 10
	}
	close(digitCh)
}
func calcSquares(number int, squaresOp chan int) {
	sum := 0
	digitCh := make(chan int)
	go digit(number, digitCh)
	for v := range digitCh {
		sum += v * v
	}
	squaresOp <- sum
}

func calcCubes(number int, cubesOp chan int) {
	sum := 0
	digitCh := make(chan int)
	go digit(number, digitCh)
	for v := range digitCh {
		sum += v * v * v
	}

	cubesOp <- sum
}

func main() {
	number := 589
	squaresChan := make(chan int)
	cubesChan := make(chan int)

	go calcSquares(number, squaresChan)
	go calcCubes(number, cubesChan)

	squares, cubes := <-squaresChan, <-cubesChan
	fmt.Println("Final output", squares+cubes)

}
