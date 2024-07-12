package main

import (
	"sync"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	numEmployees := 100
	ch := make(chan interface{}, numEmployees)
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		for j := 1; j <= numEmployees; j++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				createEmployee(id, ch)
			}(j)
		}
	}

	wg.Wait()
	close(ch)
	for range ch {
	}

	ch = make(chan interface{}, numEmployees)
	wg = sync.WaitGroup{}
}
