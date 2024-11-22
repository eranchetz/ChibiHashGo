package chibihash

import (
	"fmt"
)

// This example file demonstrates how to use the ChibiHash64 function
func main() {
	// Example 1: Hash an empty key with a seed of 0
	emptyKey := []byte("")
	seed := uint64(0)
	hash := ChibiHash64(emptyKey, len(emptyKey), seed)
	fmt.Printf("Hash of empty key with seed 0: %x\n", hash)

	// Example 2: Hash a short string with a small seed
	shortKey := []byte("hello")
	seed = uint64(1)
	hash = ChibiHash64(shortKey, len(shortKey), seed)
	fmt.Printf("Hash of 'hello' with seed 1: %x\n", hash)

	// Example 3: Hash a medium-length string with a random seed
	mediumKey := []byte("This is a test key.")
	seed = uint64(12345)
	hash = ChibiHash64(mediumKey, len(mediumKey), seed)
	fmt.Printf("Hash of 'This is a test key.' with seed 12345: %x\n", hash)

	// Example 4: Hash a longer string with a different seed
	longKey := []byte("ChibiHash function test with a significantly longer key")
	seed = uint64(67890)
	hash = ChibiHash64(longKey, len(longKey), seed)
	fmt.Printf("Hash of longer key with seed 67890: %x\n", hash)

	// Example 5: Hash with a specific seed and message
	specificKey := []byte("hash with a different seed")
	seed = uint64(54321)
	hash = ChibiHash64(specificKey, len(specificKey), seed)
	fmt.Printf("Hash of 'hash with a different seed' with seed 54321: %x\n", hash)
}
