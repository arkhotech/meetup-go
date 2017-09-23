package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	abc := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", ""}
	lenAbc := len(abc)
	numGoRutinas := lenAbc
	wg.Add(numGoRutinas)

	for i := 0; i < lenAbc; i++ {
		go imprimirRuna(abc[i], &wg)
	}

	wg.Wait()
}

func imprimirRuna(letra string, wg *sync.WaitGroup) {

	if letra != "" {
		fmt.Printf("%s", letra)
	}

	//Demorar
	result := 0.0
	for i := 0; i < 10000000; i++ {
		result += math.Pi * math.Sin(float64(i))
	}

	wg.Done()
}
