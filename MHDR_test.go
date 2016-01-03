package lora

import (
	"testing"
)

func TestParseMHDR(t *testing.T) {
	m := ParseMHDR(255)
	if m.MType != 1 {
		t.Fail()
	}
}
