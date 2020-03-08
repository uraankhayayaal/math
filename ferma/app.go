package main

import(
	"fmt"
	"math"
	"time"
)

var n int = 3
var M int = 100

func main() {

	go solve(1,25)
	go solve(25,50)
	go solve(50,75)
	go solve(75,100)

	//solve(1,100)

	fmt.Scanln()
}

func solve(start int, end int){
	t_start:= time.Now()
	for a:=start; a<end; a++ {
		for b:=1; b<M; b++ {
			for c:=1; c<M; c++ {
				for d:=1; d<M; d++ {
					var left = math.Pow(float64(a), float64(n)) + math.Pow(float64(b), float64(n)) + math.Pow(float64(d), float64(d))
					var rigth = math.Pow(float64(c), float64(n))
					if left == rigth {
						fmt.Println(a, b, c, d);
					}
				}
			}
		}
	}
	t_end := time.Now()
	elapsed := t_end.Sub(t_start)
	fmt.Println("Script Time:", elapsed);
}