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

// NewBloomFilter creates a BloomFilter of the given size and creates hashCount random salts
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

// dataToPositions converts data to positions of bits within f.filter by performing a SHA256 for each f.salts
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

// Add saves d to the internal BitVector
func (f *BloomFilter) Add(d []byte) {
	for _, pos := range f.dataToPositions(d) {
		// Set bit at position
		f.filter.SetBit(pos)
	}
}

// Probe tests whether d might be in the set
func (f *BloomFilter) Probe(d []byte) bool {
	for _, pos := range f.dataToPositions(d) {
		// Check bit at position
		if f.filter.GetBit(pos) == 0 {
			return false
		}
	}

	return true
}

// Print logs information about the BloomFilter to STDOUT for debugging use
func (f *BloomFilter) Print() {
	fmt.Printf("BloomFilter [Hashes: %d]:\n", len(f.salts))

	if f.filter.Size() <= 64 {
		fmt.Print(" Filter: ")
		f.filter.Print()
	}

	fmt.Printf(" Saturation: %.2f%%\n", f.filter.Saturation()*100)
}
