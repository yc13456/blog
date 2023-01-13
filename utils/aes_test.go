package utils

import (
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	key :="zxpT6X3zjvlp5GKU5fV2cR3PGUZwFt11"
	enstr :=AesEncrypt("sdffdf",key)
	t.Logf("encrypt strings : %s", enstr)

	oristr:=AesDecrypt(enstr,key)
	t.Logf("origin strings %s",oristr)
}
