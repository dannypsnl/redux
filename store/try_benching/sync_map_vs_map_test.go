package main

import (
	"sync"
	"testing"
)

func mapInsert() {
	m := map[int]int{}
	for i := 0; i < 100; i++ {
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
	for i := 0; i < 100; i++ {
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
	for i := 0; i < 100; i++ {
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

func BenchmarkMapRead(b *testing.B) {
	v := 0
	m := map[int]int{}
	for i := 0; i < 10000; i++ {
		m[i] = 1
	}
	for n := 0; n < b.N; n++ {
		v = m[8000]
	}
	print(v)
}

func BenchmarkSyncMapRead(b *testing.B) {
	smap := sync.Map{}
	for i := 0; i < 10000; i++ {
		smap.Store(i, 1)
	}
	for n := 0; n < b.N; n++ {
		smap.Load(8000)
	}
}
