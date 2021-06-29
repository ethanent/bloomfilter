# bloomfilter

> Go Bloom Filter implementation, with helper types such as BitVector

[![Go Reference](https://pkg.go.dev/badge/github.com/ethanent/bloomfilter.svg)](https://pkg.go.dev/github.com/ethanent/bloomfilter)

## Install

```shell
go get github.com/ethanent/bloomfilter
```

## Example

```go
bf := bloomfilter.NewBloomFilter(100000, 3)

bf.Add([]byte("hello"))
bf.Add([]byte("morning"))

bf.Probe([]byte("testing"))
// => false
// (definitely not in set)

bf.Probe([]byte("hello"))
// => true
// (possibly in set)
```
