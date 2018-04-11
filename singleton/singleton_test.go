package singleton

import (
	"testing"
)

func BenchmarkSyncOnce(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleTonOfSyncOnce()
	}
	b.ReportAllocs()
}

func BenchmarkMutex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleTonOfMutex()
	}
	b.ReportAllocs()
}

func BenchmarkMutexLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleTonOfMutexLock()
	}
	b.ReportAllocs()
}

func BenchmarkMutexPal(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SingleTonOfMutex()
		}
	})
	b.ReportAllocs()
}

func BenchmarkSyncOncePal(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SingleTonOfSyncOnce()
		}
	})
	b.ReportAllocs()
}

func BenchmarkMutexLockPal(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SingleTonOfMutexLock()
		}
	})
	b.ReportAllocs()
}
