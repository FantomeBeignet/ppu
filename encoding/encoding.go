package encoding

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	"slices"
	"strings"

	"github.com/kklash/wordlist4096"
	"golang.org/x/crypto/argon2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	WORD_LEN_BIT = wordlist4096.BitsPerWord
)

var (
	one      = big.NewInt(1)
	wordMask = big.NewInt(int64((1 << WORD_LEN_BIT) - 1))
)

type EncodedPassphrase struct {
	value    *big.Int
	numWords uint
}

func maxBigIntWithLen(bitLen uint) *big.Int {
	maxPowTwo := new(big.Int).Lsh(one, bitLen)
	return maxPowTwo.Sub(maxPowTwo, one)
}

func NewFromSeed(seed, salt []byte, numWords uint) *EncodedPassphrase {
	bitLen := numWords * WORD_LEN_BIT
	hashedSeed := argon2.IDKey(seed, nil, 1, 64*1024, 4, uint32(bitLen))
	val := new(big.Int).SetBytes(hashedSeed)
	return &EncodedPassphrase{val, numWords}
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

func FromString(s string) (*EncodedPassphrase, error) {
	value, err := new(big.Int).SetString(s, 0)
	if err != true {
		return nil, errors.New("invalid input representation")
	}
	numWords := uint(math.Ceil(float64(value.BitLen()) / float64(WORD_LEN_BIT)))
	return &EncodedPassphrase{value, numWords}, nil
}

func FromWords(words []string) (*EncodedPassphrase, error) {
	value := new(big.Int)
	reversed := make([]string, len(words))
	copy(reversed, words)
	slices.Reverse(reversed)
	for _, w := range reversed {
		index, ok := wordlist4096.WordMap[strings.ToLower(w)]
		if !ok {
			return nil, errors.New(fmt.Sprintf("word %s not in wordlist", w))
		}
		value.Lsh(value, WORD_LEN_BIT)
		value.Or(value, big.NewInt(int64(index)))
	}
	numWords := uint(len(words))
	return &EncodedPassphrase{value, numWords}, nil
}

func (b *EncodedPassphrase) WordIndices() []uint16 {
	res := make([]uint16, b.numWords)
	valCopy := new(big.Int).Set(b.value)
	for i := range b.numWords {
		res[i] = uint16(new(big.Int).And(valCopy, wordMask).Uint64())
		valCopy.Rsh(valCopy, WORD_LEN_BIT)
	}
	return res
}

func (b *EncodedPassphrase) Words(capitalize bool) []string {
	indices := b.WordIndices()
	res := make([]string, b.numWords)
	for i, index := range indices {
		if capitalize {
			res[i] = cases.Title(language.English).String(wordlist4096.WordList[index%4096])
		} else {
			res[i] = wordlist4096.WordList[index%4096]
		}
	}
	return res
}

func (b *EncodedPassphrase) ToString(base int) string {
	return b.value.Text(base)
}
