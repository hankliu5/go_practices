package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	size := 2097152

	fmt.Println("generate numbers...")
	s := make([]float64, size)
	for i := 0; i < cap(s); i++ {
		s[i] = rand.Float64() * float64(size)
	}

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
