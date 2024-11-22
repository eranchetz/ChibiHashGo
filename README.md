
# ChibiHashGo

ChibiHashGo is a simple, fast 64-bit hash function implemented in Go. It's designed for performance and ease of use in applications where hashing is required.
Go port of N-R-K/ChibiHash. See the article ChibiHash: A small, fast 64-bit hash function for more information.

## Features
- Minimal dependencies
- Optimized for performance
- Provides a small, fast 64-bit hash function

## Installation

To install ChibiHashGo, use `go get`:

```bash
go get github.com/yourusername/ChibiHashGo
```

## Usage

Import the package and use `ChibiHash64` to compute hash values.

Example:

```go
package main

import (
    "fmt"
    "github.com/yourusername/ChibiHashGo"
)

func main() {
    key := []byte("example")
    seed := uint64(42)
    hash := chibihash.ChibiHash64(key, len(key), seed)
    fmt.Printf("Hash: %x\n", hash)
}
```

## API Reference

### `func ChibiHash64(key []byte, length int, seed uint64) uint64`

- **key**: The data to hash as a byte slice.
- **length**: The length of the data.
- **seed**: A 64-bit seed value for hashing.
- **Returns**: A 64-bit unsigned integer hash.

## Testing

To test the module:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the [Unlicense](https://unlicense.org/).
