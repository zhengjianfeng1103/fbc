package secp256k1

import (
	"io"
	"testing"

	"github.com/zhengjianfeng1103/fbc/libs/tendermint/crypto"
	"github.com/zhengjianfeng1103/fbc/libs/tendermint/crypto/internal/benchmarking"
)

func BenchmarkKeyGeneration(b *testing.B) {
	benchmarkKeygenWrapper := func(reader io.Reader) crypto.PrivKey {
		return genPrivKey(reader)
	}
	benchmarking.BenchmarkKeyGeneration(b, benchmarkKeygenWrapper)
}

func BenchmarkSigning(b *testing.B) {
	priv := GenPrivKey()
	benchmarking.BenchmarkSigning(b, priv)
}

func BenchmarkVerification(b *testing.B) {
	priv := GenPrivKey()
	benchmarking.BenchmarkVerification(b, priv)
}
