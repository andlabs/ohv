// 5 january 2016
package main

import (
	"testing"
)

const mshipath = "doc/mshi/"

func BenchmarkReadMSHI(b *testing.B) {
	m := new(MSHI)
	m.dir = mshipath
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.containers = nil
		m.assets = nil
		err := m.readAllIndices()
		if err != nil {
			b.Fatal(err)
		}
	}
}
