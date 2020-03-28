package lib

import "testing"

func TestBinHash(t *testing.T) {
	hash := BinHash("CAR1072")

	if hash != 0xD6C2EDDF {
		t.Errorf("Expected BinHash(CAR1072) = 0xD6C2EDDF, got 0x%08X", hash)
		t.FailNow()
	}
}

func TestVltHash(t *testing.T) {
	hash := VltHash("car1072")

	if hash != 0x6618FEF9 {
		t.Errorf("Expected VltHash(car1072) = 0x6618FEF9, got 0x%08X", hash)
		t.FailNow()
	}
}
