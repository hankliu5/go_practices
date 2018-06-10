package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

func Merge(ldata []float64, rdata []float64) (result []float64) {
	result = make([]float64, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx] < rdata[ridx]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}
	return
}

func MultiMergeSort(data []float64, res chan []float64) {
	if len(data) < 2 {
		res <- data
		return
	}

	leftChan := make(chan []float64)
	rightChan := make(chan []float64)
	middle := len(data) / 2

	go MultiMergeSort(data[:middle], leftChan)
	go MultiMergeSort(data[middle:], rightChan)
	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	res <- Merge(ldata, rdata)
	return
}

func MultiMergeSortWithSem(data []float64, sem chan struct{}) []float64 {
	if len(data) < 2 {
		return data
	}

	middle := len(data) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var ldata []float64
	var rdata []float64

	select {
	case sem <- struct{}{}:
		go func() {
			ldata = MultiMergeSortWithSem(data[:middle], sem)
			<-sem
			wg.Done()
		}()
	default:
		ldata = SingleMergeSort(data[:middle])
		wg.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			rdata = MultiMergeSortWithSem(data[middle:], sem)
			<-sem
			wg.Done()
		}()
	default:
		rdata = SingleMergeSort(data[middle:])
		wg.Done()
	}

	wg.Wait()
	return Merge(ldata, rdata)
}

func SingleMergeSort(data []float64) []float64 {
	if len(data) < 2 {
		return data
	}
	middle := len(data) / 2
	return Merge(SingleMergeSort(data[:middle]), SingleMergeSort(data[middle:]))
}

func main() {
	size := 2097152
	sem := make(chan struct{}, 4)

	fmt.Println("generate numbers...")
	s := make([]float64, size)
	for i := 0; i < cap(s); i++ {
		s[i] = rand.Float64() * float64(size)
	}

	fmt.Println("running multithread without limited number of threads")
	start := time.Now()
	res := make(chan []float64)
	go MultiMergeSort(s, res)
	multiResult := <-res
	fmt.Println(time.Since(start))

	fmt.Println("running multithread with limited number of threads")
	start = time.Now()
	multiResultWithSem := MultiMergeSortWithSem(s, sem)
	fmt.Println(time.Since(start))

	fmt.Println("running single thread")
	start = time.Now()
	singleResult := SingleMergeSort(s)
	fmt.Println(time.Since(start))

	fmt.Println("Verifying the answer")
	fmt.Println(reflect.DeepEqual(singleResult, multiResult) &&
		reflect.DeepEqual(singleResult, multiResultWithSem))
}
