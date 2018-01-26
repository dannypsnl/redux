package main

import (
	"sync"
	"testing"
)

func mapInsert() {
	m := map[int]int{}
	for i := 0; i < 10000000; i++ {
		m[i] = 1
	}
}

func BenchmarkMapInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mapInsert()
	}
}

func syncMapInsertNoConc() {
	smap := sync.Map{}
	for i := 0; i < 10000000; i++ {
		smap.Store(i, 0)
	}
}

func BenchmarkSyncMapInsertNoConcurrency(b *testing.B) {
	for n := 0; n < b.N; n++ {
		syncMapInsertNoConc()
	}
}

func syncMapInsert() {
	var wg sync.WaitGroup
	smap := sync.Map{}
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			smap.Store(i, 0)
		}()
	}
	wg.Wait()
}

func BenchmarkSyncMapInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		syncMapInsert()
	}
}
