package main

import(
	"fmt"
	"math"
	"encoding/json"
    "io/ioutil"
    "os"
)

type SLAU struct {
	A [][]float64 `json:"a"`
	B []float64   `json:"b"`
}

var slau SLAU
var x[]float64

func main(){
    jsonFile, err := os.Open("input.json")
    if err != nil {
        fmt.Println(err)
    }

    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &slau)

	if solve() {
		var j int = 0
		for i := len(slau.B)-1; i >= 0; i-- {
			fmt.Printf("x[%v]: %v \n", j+1, x[i])
			j++
		}
	}
}

// 9 18 10 -16

func solve() bool{
	for i := 0; i < len(slau.A); i++ {
		if slau.A[i][i] == 0 {
			// Делаем перестановку столбцов, подставляем наибольший коэффициент в диоганальный элемент
			var notZeroElementIndex int = 0
			for j:= i; j < len(slau.A[i]); j++ {
				if math.Abs(slau.A[i][j]) > 0 {
					notZeroElementIndex = j
					break // находим первый не нулевой элемент, не будем проходить все элементы
				}
			}
			fmt.Println("not zero element index: ", notZeroElementIndex)
			if notZeroElementIndex != 0 {
				swapColumns(i, notZeroElementIndex)
			}
		}

		if slau.A[i][i] == 0 {
			if slau.B[i] == 0 {
				fmt.Println("Infite solutions");
				return false;
			}else{
				fmt.Println("No solutions");
				return false;
			}
		}

		for l := 0; l < len(slau.A); l++ {
			if l != i { // проходимся по остальным строкам
				var k float64 = slau.A[l][i] / slau.A[i][i]
				for j := i; j < len(slau.A[i]); j++ {
					slau.A[l][j] = slau.A[l][j] - slau.A[i][j]*k
				}
				slau.B[l] = slau.B[l] - slau.B[i]*k
			}
		}
	}
	
	for i := len(slau.A)-1; i>=0; i-- {
		var summ float64 = 0
		for j := len(slau.A)-1; j>i; j-- {
			summ = summ + slau.A[i][j]
		}
		x = append(x, (slau.B[i] - summ)/slau.A[i][i])
		for l:= i; l>=0; l-- {
			slau.A[l][i] = slau.A[l][i] * (slau.B[i] - summ)/slau.A[i][i];
		}
	}
	return true
}

func swapColumns(j int, notZeroElementIndex int){
	for i := 0; i < len(slau.A); i++ {
		slau.A[i][j], slau.A[i][notZeroElementIndex] = slau.A[i][notZeroElementIndex], slau.A[i][j]
	}
}