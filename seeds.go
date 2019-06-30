package fraud

import (
	cryptoRand "crypto/rand"
	"sync"

	"golang.org/x/exp/rand"
)

const SeedSize = 64

var (
	internalSource      = rand.NewSource(CryptoSeed())
	internalSourceMutex sync.Mutex
)

func CryptoSeed() uint64 {
	var seed uint64
	var buf [SeedSize / 8]byte
	cryptoRand.Read(buf[:])
	for i := 0; i < len(buf); i++ {
		seed = (seed << 8) | uint64(buf[i])
	}
	return seed
}

func Uint64() uint64 {
	internalSourceMutex.Lock()
	n := internalSource.Uint64()
	internalSourceMutex.Unlock()
	return n
}

func Seed() uint64 {
	return Uint64()
}

func Rand() *rand.Rand {
	return rand.New(rand.NewSource(Seed()))
}
