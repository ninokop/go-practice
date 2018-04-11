package singleton

import (
	"sync"
)

var (
	singleTon *SingleTon
	mux       sync.Mutex
	once      sync.Once
)

type SingleTon struct{}

/* Benchmark says use dubbole check please
BenchmarkSyncOnce-4       	200000000	         6.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutex-4          	300000000	         5.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexLock-4      	20000000	        58.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexPal-4       	500000000	         3.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncOncePal-4    	300000000	         4.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexLockPal-4   	10000000	       159 ns/op	       0 B/op	       0 allocs/op
*/

func SingleTonOfMutex() *SingleTon {
	if singleTon == nil {
		mux.Lock()
		defer mux.Unlock()
		if singleTon == nil { // dubbole check
			singleTon = &SingleTon{}
		}
	}
	return singleTon
}

func SingleTonOfSyncOnce() *SingleTon {
	once.Do(func() {
		singleTon = &SingleTon{}
	})
	return singleTon
}

func SingleTonOfMutexLock() *SingleTon {
	mux.Lock()
	defer mux.Unlock()
	if singleTon == nil {
		singleTon = &SingleTon{}
	}
	return singleTon
}
