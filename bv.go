package bloomfilter

import "fmt"

type BitVector struct {
	data []byte
	size int
}

// getIdxs returns byte and within-byte bit indexes
func (v *BitVector) getIdxs(idx int) (byteIdx int, bitIdx int) {
	if idx > v.Size() {
		panic("BitVector access index out of range")
	}

	return idx / 8, idx % 8
}

// NewBitVector creates a zero-filled BitVector
func NewBitVector(size int) *BitVector {
	d := make([]byte, size/8+1)

	return &BitVector{data: d, size: size}
}

// NewBitVectorFromData creates a BitVector from data
func NewBitVectorFromData(data []byte) *BitVector {
	return &BitVector{data: data, size: len(data)}
}

// SetBit sets the bit at index idx
func (v *BitVector) SetBit(idx int) {
	byteIdx, bitIdx := v.getIdxs(idx)

	v.data[byteIdx] = v.data[byteIdx] | (1 << bitIdx)
}

// ClrBit clears the bit at index idx
func (v *BitVector) ClrBit(idx int) {
	byteIdx, bitIdx := v.getIdxs(idx)

	v.data[byteIdx] = v.data[byteIdx] & ^(1 << bitIdx)
}

// GetBit returns 1 or 0, depending on the bit value at index idx
func (v *BitVector) GetBit(idx int) int {
	byteIdx, bitIdx := v.getIdxs(idx)

	return int((v.data[byteIdx] >> bitIdx) & 1)
}

// Size returns the size of v
func (v *BitVector) Size() int {
	return v.size
}

// Saturation returns the proportion of bits within v which are set
func (v *BitVector) Saturation() float64 {
	ones := 0
	tot := 0

	for i := 0; i < v.Size(); i++ {
		ones += v.GetBit(i)
		tot += 1
	}

	return float64(ones) / float64(tot)
}

// Int converts the BitVector to an integer
func (v *BitVector) Int() int {
	r := 0
	cp := 1

	for i := 0; i < v.Size(); i++ {
		r += cp * v.GetBit(i)

		cp *= 2
	}

	return r
}

// Print logs information about the BitVector to STDOUT for debugging use
func (v *BitVector) Print() {
	fmt.Printf("BitVector [Size: %d] {", v.Size())

	for i := 0; i < v.Size(); i++ {
		fmt.Printf("%d", v.GetBit(i))

		if (i+1)%8 == 0 && i != 0 && i != v.Size()-1 {
			fmt.Print(" ")
		}
	}

	fmt.Print("}\n")
}
