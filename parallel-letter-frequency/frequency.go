// Package letter counts the frequency of letters in texts using parallel
// computation.
package letter

import (
	"sync"
)

// FreqMap counts how many times a rune has appeared.
type FreqMap map[rune]int

// Frequency loops through runes in a string and updates the values in the
// FreqMap accordingly.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency creates a FreqMap using goroutines.
func ConcurrentFrequency(l []string) FreqMap {
	n := len(l)
	jobs := make(chan int, n)
	fmaps := make(chan FreqMap, n)

	var wg sync.WaitGroup
	wg.Add(n)
	for _, s := range l {
		go calculateFrequency(s, jobs, fmaps, &wg)
	}
	wg.Wait()

	defer close(fmaps)
	defer close(jobs)

	return reduce(n, fmaps)
}

// Internal function to concurrently calculate the frequency.
func calculateFrequency(s string, jobs chan int, fmaps chan FreqMap, wg *sync.WaitGroup) {
	defer wg.Done()
	jobs <- 1
	fmaps <- Frequency(s)
	<-jobs
}

// Internal function to reduce a channel of maps, given the total jobs.
func reduce(n int, fmaps chan FreqMap) FreqMap {
	res := FreqMap{}
	for i := 0; i < n; i++ {
		m := <-fmaps
		for k, v := range m {
			res[k] += v
		}
	}
	return res
}
