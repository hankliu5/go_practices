package main

import (
	"sync"
)

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

func RunMultiMergesortWithSem(data []float64) []float64 {
	sem := make(chan struct{}, 4)
	return MultiMergeSortWithSem(data, sem)
}
