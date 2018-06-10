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

func MultiMergeSort(data []float64, sem chan struct{}) []float64 {
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
			ldata = MultiMergeSort(data[:middle], sem)
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
			rdata = MultiMergeSort(data[middle:], sem)
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
	size := 16777216
	sem := make(chan struct{}, 4)

	s := make([]float64, size)
	for i := 0; i < cap(s); i++ {
		s[i] = rand.Float64() * float64(size)
	}

	start := time.Now()
	multiResult := MultiMergeSort(s, sem)
	fmt.Println(time.Since(start))

	start = time.Now()
	singleResult := SingleMergeSort(s)
	fmt.Println(time.Since(start))

	fmt.Println(reflect.DeepEqual(singleResult, multiResult))
}
