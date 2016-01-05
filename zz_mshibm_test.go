// 5 january 2016
package main

import (
	"testing"
)

const mshipath = "doc/mshi/"

func testOpenMSHI(stage int, b *testing.B) {
	m := new(MSHI)
	m.dir = mshipath

	n := 1
	if stage == 0 {
		n = b.N
		b.ResetTimer()
	}
	for i := 0; i < n; i++ {
		m.containers = nil
		m.assets = nil
		err := m.readAllIndices()
		if err != nil {
			b.Fatal(err)
		}
	}
	if stage == 0 {
		return
	}

	if stage == 1 {
		n = b.N
		b.ResetTimer()
	}
	for i := 0; i < n; i++ {
		// no need to reset anything; collectTopics() does it
		m.collectTopics()
	}
	if stage == 1 {
		return
	}

	if stage == 2 {
		n = b.N
		b.ResetTimer()
	}
	for i := 0; i < n; i++ {
		m.books = nil
		m.orphans = nil
		// we can't reset children here, alas
		m.buildHierarchy()
	}
	if stage == 2 {
		return
	}

	if stage == 3 {
		n = b.N
		b.ResetTimer()
	}
	for i := 0; i < n; i++ {
		// again, can't reset state here
		m.sortAllChildren()
	}
}

func BenchmarkStage0(b *testing.B) {
	testOpenMSHI(0, b)
}

func BenchmarkStage1(b *testing.B) {
	testOpenMSHI(1, b)
}

func BenchmarkStage2(b *testing.B) {
	testOpenMSHI(2, b)
}

func BenchmarkStage3(b *testing.B) {
	testOpenMSHI(3, b)
}
