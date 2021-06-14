# bloomfilter
> Go Bloom Filter implementation, with helper types such as BitVector

## Install

```shell
go get github.com/ethanent/bloomfilter
```

## Doc

```go
package bloomfilter // import "github.com/ethanent/bloomfilter"

type BitVector struct {
        // Has unexported fields.
}

func NewBitVector(size int) *BitVector
    NewBitVector creates a zero-filled BitVector

func NewBitVectorFromData(data []byte) *BitVector
    NewBitVectorFromData creates a BitVector from data

func (v *BitVector) ClrBit(idx int)
    ClrBit clears the bit at index idx

func (v *BitVector) GetBit(idx int) int
    GetBit returns 1 or 0, depending on the bit value at index idx

func (v *BitVector) Int() int
    Int converts the BitVector to an integer

func (v *BitVector) Print()
    Print logs information about the BitVector to STDOUT for debugging use

func (v *BitVector) Saturation() float64
    Saturation returns the proportion of bits within v which are set

func (v *BitVector) SetBit(idx int)
    SetBit sets the bit at index idx

func (v *BitVector) Size() int
    Size returns the size of v

type BloomFilter struct {
        // Has unexported fields.
}

func NewBloomFilter(size int, hashCount int) (*BloomFilter, error)
    NewBloomFilter creates a BloomFilter of the given size and creates hashCount
    random salts

func (f *BloomFilter) Add(d []byte)
    Add saves d to the internal BitVector

func (f *BloomFilter) Print()
    Print logs information about the BloomFilter to STDOUT for debugging use

func (f *BloomFilter) Probe(d []byte) bool
    Probe tests whether d might be in the set
```
