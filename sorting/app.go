package main

import (
    "fmt"
    "os"
	"bufio"
	"log"
	"strconv"
)

func main(){
    input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var numbers[]int
	for scanner.Scan() {
		var myInt, _ = strconv.Atoi(scanner.Text())
		numbers = append(numbers, myInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort(numbers)

	output, err := os.Open("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	myWriter := bufio.NewWriter(output)

	for _, number := range numbers {
		myWriter.WriteString(strconv.Itoa(number))    // запись строки
		myWriter.WriteString("\n")   // перевод строки
		fmt.Println(number)
	}
}

func sort(ar []int){
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				swap(ar, i, j)
			}
		}
	}
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}