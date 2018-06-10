package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

const XMAX float64 = 1.0
const XMIN float64 = 0.0
const MAX_PSZ int = 10
const TSCALE float64 = 1.0

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: npoints npebs time_finish nthreads")
	}

	npoints, err := strconv.Atoi(os.Args[1])
	npebs, err := strconv.Atoi(os.Args[2])
	// end_time, err := strconv.ParseFloat(os.Args[3], 64)
	nthreads, err := strconv.Atoi(os.Args[4])

	if err != nil {
		fmt.Println("Convert number incorrectly")
		os.Exit(1)
	}

	if npoints%nthreads != 0 {
		fmt.Println("threads cannot separate npoints evenly.")
		os.Exit(1)
	}

	narea := npoints * npoints

	u_i0 := make([]float64, narea)
	u_i1 := make([]float64, narea)
	pebs := make([]float64, narea)
	// u_cpu := make([]float64, narea)

	// h := (XMAX - XMIN) / float64(npoints)
	initPebbles(pebs, npebs, npoints)
	initMaps(u_i0, u_i1, pebs, npoints)
	fmt.Println(pebs)
}

func initPebbles(pebs []float64, npebs int, npoints int) {
	for k := 0; k < npebs; k++ {
		i := rand.Int()%(npoints-4) + 2
		j := rand.Int()%(npoints-4) + 2
		sz := rand.Int() % MAX_PSZ
		idx := j + i*npoints
		pebs[idx] = float64(sz)
	}
}

func initMaps(u0 []float64, u1 []float64, pebs []float64, npoints int) {
	for i := 0; i < npoints; i++ {
		for j := 0; j < npoints; j++ {
			idx := j + i*npoints
			scale := f(pebs[idx], 0.0)
			u0[idx] = scale
			u1[idx] = scale
		}
	}
}

func f(p float64, t float64) float64 {
	return -math.Exp(-TSCALE*t) * p
}
