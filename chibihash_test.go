package chibihash

import (
	"testing"
)

func TestChibiHash64(t *testing.T) {
	tests := []struct {
		name   string
		key    []byte
		length int
		seed   uint64
		want   uint64
	}{
		{"Empty Key", []byte{}, 0, 0, 0x9ea80f3b18e26cfb},
		{"Short Key", []byte("hello"), 5, 1, 0xd0784fe905a47cb8},
		{"Medium Key", []byte("This is a test key."), 18, 12345, 0xcb3cbd67539cf73d},
		{"Long Key", []byte("ChibiHash function test with a significantly longer key"), 51, 67890, 0xe873a4088e8e0a39},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChibiHash64(tt.key, tt.length, tt.seed)
			if got != tt.want {
				t.Errorf("ChibiHash64(%q, %d, %d) = %x; want %x", tt.key, tt.length, tt.seed, got, tt.want)
			}
		})
	}
}
