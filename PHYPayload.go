package lora

import (
	"encoding/binary"
)

type phypayload struct {
	MHDR       *mhdr
	MACPayload []byte
	MIC        uint32
}

func ParsePHYPayload(d []byte) *phypayload {
	mhdr := ParseMHDR(d[0])
	macpayload := d[1 : len(d)-4]
	mic := binary.LittleEndian.Uint32(d[len(d)-4 : len(d)])
	return &phypayload{mhdr, macpayload, mic}
}
