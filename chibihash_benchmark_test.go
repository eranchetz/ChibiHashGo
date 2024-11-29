package chibihash

import (
	"testing"
)

func BenchmarkChibiHash64(b *testing.B) {
	testCases := []struct {
		name   string
		key    []byte
		length int
		seed   uint64
	}{
		{"Empty Key", []byte(""), 0, 0},
		{"Small Key", []byte("hi"), 2, 0},
		{"Numeric Key", []byte("123"), 3, 0},
		{"Alphabet Key", []byte("abcdefgh"), 8, 0},
		{"Greeting Key", []byte("Hello, world!"), 13, 0},
		{"Alphanumeric Key", []byte("qwertyuiopasdfghjklzxcvbnm123456"), 32, 0},
		{"Long Alphanumeric Key", []byte("qwertyuiopasdfghjklzxcvbnm123456789"), 35, 0},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Report allocations per operation
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				ChibiHash64(tc.key, tc.length, tc.seed)
			}
		})
	}
}
