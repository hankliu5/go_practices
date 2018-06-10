package main

import "fmt"
import "math/rand"

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

func MultiMergeSort(data []float64, r chan []float64) {
	if len(data) < 2 {
		r <- data
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
	r <- Merge(ldata, rdata)
	return
}

func SingleMergeSort(data []float64) []float64 {
	if len(data) < 2 {
		return data
	}
	middle := len(data) / 2
	return Merge(SingleMergeSort(data[:middle]), SingleMergeSort(data[middle:]))
}

func main() {
	size := 10
	s := make([]float64, size)
	for i := 0; i < cap(s); i++ {
		s[i] = rand.Float64() * float64(size)
	}
	result := make(chan []float64)
	go MultiMergeSort(s, result)

	fmt.Println("Multithreading...")
	r := <-result
	fmt.Println(r)

	close(result)

	fmt.Println("Singlethreading...")
	singleResult := SingleMergeSort(s)
	fmt.Println(singleResult)

}
