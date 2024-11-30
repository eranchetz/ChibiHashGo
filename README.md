
# ChibiHashGo

ChibiHashGo is a simple, fast 64-bit hash function implemented in Go. It is designed for performance and ease of use in applications where efficient hashing is required.

## Features
- Minimal dependencies
- Optimized for performance
- Supports inputs of varying sizes (small keys to large datasets)
- Includes comprehensive benchmarks and GitHub Actions integration for CI

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


## Benchmarks

The following benchmarks demonstrate the performance of `ChibiHash64` for varying input sizes. These results were obtained using `go test -bench=. -benchmem`.

| Test Case                     | ns/op  | B/op | allocs/op |
|-------------------------------|--------|------|-----------|
| 5 Bytes                       | 4.215  | 0    | 0         |
| 100 Bytes                     | 5.155  | 0    | 0         |
| 4 KB                          | 44.78  | 0    | 0         |
| 10 MB                         | 888191 | 0    | 0         |

The function performs efficiently with no memory allocations (`allocs/op = 0`) across all test cases, ensuring minimal overhead.

## CI/CD Integration

GitHub Actions is set up to automatically run tests and benchmarks for every push or pull request.

### Workflow Configuration

The workflow file is located at `.github/workflows/go.yml` and includes:
- Installing dependencies (`go mod tidy`)
- Running tests (`go test ./... -v`)
- Running benchmarks (`go test -bench=. -benchmem`)

## Testing

To test the module locally:

```bash
go test ./... -v
```

To run benchmarks:

```bash
go test -bench=. -benchmem
```

## Contributing

Contributions are welcome! Fork the repository, make your changes, and submit a pull request. Ensure all tests pass and benchmarks are updated before submitting.

## License

This project is licensed under the [Unlicense](https://unlicense.org/).
