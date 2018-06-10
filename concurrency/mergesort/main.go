package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

func generateArray(numOfElements int) []float64 {
	s := make([]float64, numOfElements)
	for i := 0; i < cap(s); i++ {
		s[i] = rand.Float64() * float64(numOfElements)
	}
	return s
}

func main() {
	var size int
	if len(os.Args) < 2 {
		size = 2097152
	} else {
		size, _ = strconv.Atoi(os.Args[1])
	}

	fmt.Println("generate " + strconv.Itoa(size) + " numbers...")
	s := generateArray(size)

	fmt.Println("running multithread without limited number of threads")
	start := time.Now()
	multiResult := runMultiMergeSort(s)
	fmt.Println(time.Since(start))

	fmt.Println("running multithread with limited number of threads")
	start = time.Now()
	multiResultWithSem := runMultiMergesortWithSem(s)
	fmt.Println(time.Since(start))

	fmt.Println("running single thread")
	start = time.Now()
	singleResult := SingleMergeSort(s)
	fmt.Println(time.Since(start))

	fmt.Println("Verifying the answer")
	fmt.Println(reflect.DeepEqual(singleResult, multiResult) &&
		reflect.DeepEqual(singleResult, multiResultWithSem))
}
