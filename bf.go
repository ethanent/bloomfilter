package bloomfilter

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

type BloomFilter struct {
	filter *BitVector
	salts  [][]byte
}

func NewBloomFilter(size int, hashCount int) (*BloomFilter, error) {
	bf := &BloomFilter{}

	bf.filter = NewBitVector(size)

	for i := 0; i < hashCount; i++ {
		salt := make([]byte, 9)

		if _, err := rand.Read(salt); err != nil {
			return nil, err
		}

		bf.salts = append(bf.salts, salt)
	}

	return bf, nil
}

func (f *BloomFilter) dataToPositions(d []byte) []int {
	positions := []int{}

	for _, salt := range f.salts {
		hashData := append(d, salt...)

		hr := sha256.Sum256(hashData)

		// Convert hr [32]byte to hri int

		hrv := NewBitVectorFromData(hr[:])

		hri := hrv.Int()

		// Note bit position: hri % f.filter.Size()

		positions = append(positions, hri%f.filter.Size())
	}

	return positions
}

func (f *BloomFilter) Add(d []byte) {
	for _, pos := range f.dataToPositions(d) {
		// Set bit at position
		f.filter.SetBit(pos)
	}
}

func (f *BloomFilter) Probe(d []byte) bool {
	for _, pos := range f.dataToPositions(d) {
		// Check bit at position
		if f.filter.GetBit(pos) == 0 {
			return false
		}
	}

	return true
}

func (f *BloomFilter) Print() {
	fmt.Printf("BloomFilter [Hashes: %d]:\n", len(f.salts))

	if f.filter.Size() <= 64 {
		fmt.Print(" Filter: ")
		f.filter.Print()
	}

	fmt.Printf(" Saturation: %.2f%%\n", f.filter.Saturation()*100)
}
