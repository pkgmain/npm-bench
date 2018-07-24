package main

import "testing"

func BenchmarkNpm(b *testing.B) {
	for n := 0; n < b.N; n++ {
		err := npm()
		if err != nil {
			b.Fatal(err)
		}
	}
}
