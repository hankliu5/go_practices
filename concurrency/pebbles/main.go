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
const VSQR float64 = 0.1

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: npoints npebs time_finish nthreads")
	}

	npoints, err := strconv.Atoi(os.Args[1])
	npebs, err := strconv.Atoi(os.Args[2])
	end_time, err := strconv.ParseFloat(os.Args[3], 64)
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
	u_cpu := make([]float64, narea)

	h := (XMAX - XMIN) / float64(npoints)
	initPebbles(pebs, npebs, npoints)
	initMaps(u_i0, u_i1, pebs, npoints)

	runSimulation(u_cpu, u_i0, u_i1, pebs, npoints, h, end_time)

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
			u0[idx] = pebs[idx]
			u1[idx] = pebs[idx]
		}
	}
}

func f(p float64, t float64) float64 {
	return -math.Exp(-TSCALE*t) * p
}

func runSimulation(u []float64, u0 []float64, u1 []float64, pebs []float64, npoints int, h float64, end_time float64) {
	un := make([]float64, npoints*npoints)
	uc := make([]float64, npoints*npoints)
	uo := make([]float64, npoints*npoints)

	for i := 0; i < npoints*npoints; i++ {
		uo[i] = u0[i]
		uc[i] = u1[i]
	}
	t := 0.0
	dt := h / 2.0

	for t < end_time {
		for i := 0; i < npoints; i++ {
			for j := 0; j < npoints; j++ {
				idx := j + i*npoints

				if i == 0 || i == npoints-1 || j == 0 || j == npoints-1 {
					un[idx] = 0.0
				} else {
					un[idx] = 2*uc[idx] - uo[idx] + VSQR*(dt*dt)*((uc[idx-1]+uc[idx+1]+
						uc[idx+npoints]+uc[idx-npoints]+0.25*
						(uc[idx-npoints-
							1]+
							uc[idx+npoints-
								1]+
							uc[idx-npoints+
								1]+
							uc[idx+npoints+
								1])-
						5*uc[idx])/(h*h)+
						f(pebs[idx], t))
				}
			}
		}
		for i := 0; i < npoints*npoints; i++ {
			uo[i] = uc[i]
			uc[i] = un[i]
		}
		t += dt
	}

	for i := 0; i < npoints*npoints; i++ {
		u[i] = un[i]
	}
}
