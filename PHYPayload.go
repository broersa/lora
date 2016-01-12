package lora

import (
	"bytes"
	"encoding/binary"
)

// PHYPayload ...
type PHYPayload struct {
	MHDR       *MHDR
	MACPayload []byte
	MIC        uint32
}

// ParsePHYPayload ...
func ParsePHYPayload(d []byte) (*PHYPayload, error) {
	mhdr, err := ParseMHDR(d[0])
	if err != nil {
		return nil, err
	}
	macpayload := d[1 : len(d)-4]
	mic := binary.LittleEndian.Uint32(d[len(d)-4 : len(d)])
	return &PHYPayload{mhdr, macpayload, mic}, nil
}

func MarshallPHYPayload(mhdr byte, macpayload []byte, mic []byte) ([]byte, error) {
	b0 := new(bytes.Buffer)
	b0.WriteByte(mhdr)
	b0.Write(macpayload)
	b0.Write(mic)

	return b0.Bytes(), nil
}
