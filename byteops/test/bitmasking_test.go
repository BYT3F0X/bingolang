package test

import (
	"github.com/BYT3F0X/bingolang/byteops"
	"github.com/BYT3F0X/bingolang/util"
	"testing"
)

func TestBitMask(t *testing.T) {
	var a uint8
	if end, _ := util.DetermineEndianness(); end == util.LittleEndian {
		a = 1
		if byteops.QueryBit(a, 0) == false {
			t.Failed()
		} else if a = 2; byteops.QueryBit(a, 1) == false {
			t.Failed()
		} else if a = 4; byteops.QueryBit(a, 2) == false {
			t.Failed()
		} else if a = 8; byteops.QueryBit(a, 3) == false {
			t.Failed()
		} else if a = 16; byteops.QueryBit(a, 4) == false {
			t.Failed()
		} else if a = 32; byteops.QueryBit(a, 5) == false {
			t.Failed()
		} else if a = 64; byteops.QueryBit(a, 6) == false {
			t.Failed()
		} else if a = 128; byteops.QueryBit(a, 7) == false {
			t.Failed()
		}
	} else {
		t.Skipf("skipped because the system is not little endian")
	}
}
