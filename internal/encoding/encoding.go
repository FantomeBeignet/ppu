package encoding

import (
	"crypto/rand"
	"math/big"

	"github.com/kklash/wordlist4096"
)

type EncodedPassphrase struct {
	value    *big.Int
	numWords uint
}

const (
	WORD_LEN_BIT = wordlist4096.BitsPerWord
)

var (
	one      = big.NewInt(1)
	wordMask = big.NewInt(int64((1 << WORD_LEN_BIT) - 1))
)

func maxBigIntWithLen(bitLen uint) *big.Int {
	maxPowTwo := new(big.Int).Lsh(one, bitLen)
	return maxPowTwo.Sub(maxPowTwo, one)
}

func NewRandom(numWords uint) (*EncodedPassphrase, error) {
	bitLen := numWords * WORD_LEN_BIT
	maxInt := maxBigIntWithLen(bitLen)
	val, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return nil, err
	}
	return &EncodedPassphrase{val, numWords}, nil
}

func (b *EncodedPassphrase) ToWordIndices() []uint16 {
	res := make([]uint16, b.numWords)
	for i := range b.numWords {
		res[i] = uint16(new(big.Int).And(b.value, wordMask).Uint64())
		b.value.Rsh(b.value, WORD_LEN_BIT)
	}
	return res
}

func (b *EncodedPassphrase) ToWords(capitalize bool) []string {
	indices := b.ToWordIndices()
	res := make([]string, b.numWords)
	for i, index := range indices {
		res[i] = wordlist4096.WordList[index]
	}
	return res
}
