package main

func SingleMergeSort(data []float64) []float64 {
	if len(data) < 2 {
		return data
	}
	middle := len(data) / 2
	return Merge(SingleMergeSort(data[:middle]), SingleMergeSort(data[middle:]))
}
