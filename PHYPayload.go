package lora

import (
	"encoding/binary"
)

// PHYPayload ...
type PHYPayload struct {
	MHDR       *mhdr
	MACPayload []byte
	MIC        uint32
}

// ParsePHYPayload ...
func ParsePHYPayload(d []byte) (*PHYPayload, error) {
	mhdr := ParseMHDR(d[0])
	macpayload := d[1 : len(d)-4]
	mic := binary.LittleEndian.Uint32(d[len(d)-4 : len(d)])
	return &PHYPayload{mhdr, macpayload, mic}, nil
}
