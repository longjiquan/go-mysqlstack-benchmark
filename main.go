package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func single(n int) {
	fmt.Println("single worker bench")

	e := queryRow()
	fmt.Printf("warmup, id: %d, name: %s\n", e.id, e.name)

	{
		start := time.Now()
		for i := 0; i < n; i++ {
			queryRow()
		}
		cost := time.Since(start)
		fmt.Printf("single row query, n: %d, time cost: %s, rps: %f\n", n, cost.String(), float64(n)/cost.Seconds())
	}

	{
		start := time.Now()
		for i := 0; i < n; i++ {
			queryMultiRow()
		}
		cost := time.Since(start)
		fmt.Printf("multiple row query, n: %d, time cost: %s, rps: %f\n", n, cost.String(), float64(n)/cost.Seconds())
	}
}

func work(wg *sync.WaitGroup, fn func()) {
	defer wg.Done()
	fn()
}

func multiple(nWorker, nPerWorker int) {
	fmt.Printf("multiple worker bench, n_worker: %d, n_per_worker: %d\n", nWorker, nPerWorker)

	e := queryRow()
	fmt.Printf("warmup, id: %d, name: %s\n", e.id, e.name)

	{
		var wg sync.WaitGroup
		start := time.Now()
		for i := 0; i < nWorker; i++ {
			wg.Add(1)
			go work(&wg, func() {
				for i := 0; i < nPerWorker; i++ {
					queryRow()
				}
			})
		}
		wg.Wait()
		cost := time.Since(start)
		fmt.Printf("single row query, n_worker: %d, n_per_worker: %d, time cost: %s, rps: %f\n", nWorker, nPerWorker, cost.String(), float64(nWorker*nPerWorker)/cost.Seconds())
	}

	{
		var wg sync.WaitGroup
		start := time.Now()
		for i := 0; i < nWorker; i++ {
			wg.Add(1)
			go work(&wg, func() {
				for i := 0; i < nPerWorker; i++ {
					queryMultiRow()
				}
			})
		}
		wg.Wait()
		cost := time.Since(start)
		fmt.Printf("single row query, n_worker: %d, n_per_worker: %d, time cost: %s, rps: %f\n", nWorker, nPerWorker, cost.String(), float64(nWorker*nPerWorker)/cost.Seconds())
	}
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}

	nCPU := runtime.NumCPU()
	n := 1000000

	single(n)
	multiple(nCPU, n)
}
