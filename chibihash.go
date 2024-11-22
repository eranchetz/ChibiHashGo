package chibihash

// ChibiHash64 computes a small, fast 64-bit hash for a given key, length, and seed.
func ChibiHash64(key []byte, length int, seed uint64) uint64 {
	const (
		P1 = uint64(0x2B7E151628AED2A5)
		P2 = uint64(0x9E3793492EEDC3F7)
		P3 = uint64(0x3243F6A8885A308D)
	)

	h := [4]uint64{P1, P2, P3, seed}
	k := key
	l := length

	for l >= 32 {
		for i := 0; i < 4; i++ {
			lane := load64LE(k[:8])
			h[i] ^= lane
			h[i] *= P1
			h[(i+1)&3] ^= (lane << 40) | (lane >> 24)
			k = k[8:]
		}
		l -= 32
	}

	h[0] += uint64(length)<<32 | uint64(length)>>32
	if l&1 != 0 {
		h[0] ^= uint64(k[0])
		l--
		k = k[1:]
	}
	h[0] *= P2
	h[0] ^= h[0] >> 31

	for i := 1; l >= 8; i++ {
		h[i] ^= load64LE(k[:8])
		h[i] *= P2
		h[i] ^= h[i] >> 31
		k = k[8:]
		l -= 8
	}

	for i := 0; l > 0; i++ {
		h[i] ^= uint64(k[0]) | (uint64(k[1]) << 8)
		h[i] *= P3
		h[i] ^= h[i] >> 31
		k = k[2:]
		l -= 2
	}

	x := seed
	x ^= h[0] * ((h[2] >> 32) | 1)
	x ^= h[1] * ((h[3] >> 32) | 1)
	x ^= h[2] * ((h[0] >> 32) | 1)
	x ^= h[3] * ((h[1] >> 32) | 1)

	x ^= x >> 27
	x *= 0x3C79AC492BA7B653
	x ^= x >> 33
	x *= 0x1C69B3F74AC4AE35
	x ^= x >> 27

	return x
}

func load64LE(p []byte) uint64 {
	return uint64(p[0]) |
		uint64(p[1])<<8 |
		uint64(p[2])<<16 |
		uint64(p[3])<<24 |
		uint64(p[4])<<32 |
		uint64(p[5])<<40 |
		uint64(p[6])<<48 |
		uint64(p[7])<<56
}
