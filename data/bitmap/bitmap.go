package bitmap

const (
	constBitMapMask byte   = 1
	constBitsInByte uint64 = 8
	constInitByte   byte   = 0
)

type bitmap struct {
	bits []byte
}

// diveup dives and rounds up a/b
func diveup(a, b uint64) uint64 {
	if a%b == 0 {
		return a / b
	}
	return (a / b) + 1
}

// newBitMap returns a new bitmap with nbits
func newBitMap(nbits uint64) *bitmap {
	var bm bitmap

	nbytes := diveup(nbits, constBitsInByte)
	bm.bits = make([]byte, nbytes)

	var i uint64
	for i = 0; i < nbytes; i++ {
		bm.bits[i] = constInitByte
	}

	return &bm
}

// nbytes returns the number of byte in bitmap
func (bm *bitmap) nbytes() uint64 {
	return (uint64)(len(bm.bits))
}

// nbytes return the number of bits in bitmap
func (bm *bitmap) nbits() uint64 {
	return (constBitsInByte * bm.nbytes())
}

// resize resizes the bitmap and returns the nbits after resizing
func (bm *bitmap) resize(nbits uint64) uint64 {
	renbytes := diveup(nbits, constBitsInByte)

	if renbytes < bm.nbytes() {
		bm.bits = bm.bits[:renbytes]
	} else {
		rebits := make([]byte, renbytes)
		copy(rebits, bm.bits)
		bm.bits = rebits
	}

	return bm.nbits()
}

// test returns whether bitmap[bitIdx] == true
// test returns false  when bitIdx is out of range of bitmap
// O(1)
func (bm *bitmap) test(bitIdx uint64) bool {
	byteIdx := bitIdx / constBitsInByte
	idxInByte := bitIdx % constBitsInByte

	if !(byteIdx < bm.nbytes()) {
		return false // out of range, return false
	}

	return ((bm.bits[byteIdx] & (constBitMapMask << idxInByte)) != 0)
}

// set puts the value on bitmap[bitIdx]
// set does nothing when bitIdx is out of range of bitmap
// O(1)
func (bm *bitmap) set(bitIdx uint64, value bool) {
	byteIdx := bitIdx / constBitsInByte
	idxInByte := bitIdx % constBitsInByte

	if !(byteIdx < bm.nbytes()) {
		return // out of range, do nothing
	}

	if value == true {
		bm.bits[byteIdx] |= (constBitMapMask << idxInByte)
	} else {
		bm.bits[byteIdx] &= (^(constBitMapMask << idxInByte))
	}
}

// rangeSet set the same value on the consecutive space in bitmap
// O(length)
func (bm *bitmap) _rangeSet(startBitIdx, length uint64, value bool) {
	var i uint64
	for i = 0; i < length; i++ {
		bm.set(startBitIdx+i, value)
	}
}

// rangeAlloc allocates the consecutive space in bitmap and returns the starting bit index
// worst O(nbits)
func (bm *bitmap) rangeAlloc(length uint64) (startBitIdx uint64, ok bool) {
	var (
		count  uint64 = 0
		bitIdx uint64 = 0
	)

	for count < length {
		if !(bitIdx < bm.nbits()) {
			return 0, false // out of range, allocate failed
		}
		if bm.test(bitIdx) == false { // a free bit
			count++
		} else { // discontinuous
			count = 0 // re count it
		}
		bitIdx++
	}

	bm._rangeSet(bitIdx-length, length, true)

	return startBitIdx, true
}

// rangeFree frees the consecutive space in bitmap
// O(length)
func (bm *bitmap) rangeFree(startBitIdx, length uint64) {
	bm._rangeSet(startBitIdx, length, false)
}
