package state_test

import (
	"os"
	"testing"

	"github.com/zhengjianfeng1103/fbc/libs/tendermint/types"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidencesGlobal()
	os.Exit(m.Run())
}
