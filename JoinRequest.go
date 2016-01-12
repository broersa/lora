package lora

import (
	"bytes"
	"encoding/binary"

	"github.com/jacobsa/crypto/cmac"
)

type joinrequest struct {
	AppEUI   uint64
	DevEUI   uint64
	DevNonce uint16
}

func ParseJoinRequest(macpayload []byte) (*joinrequest, error) {
	return &joinrequest{binary.LittleEndian.Uint64(macpayload[0:8]),
		binary.LittleEndian.Uint64(macpayload[8:16]),
		binary.LittleEndian.Uint16(macpayload[16:18])}, nil
}

func ValidateJoinRequest(appkey []byte, mhdr byte, macpayload []byte, mic uint32) (bool, error) {
	b0 := new(bytes.Buffer)
	b0.WriteByte(mhdr)
	b0.Write(macpayload)
	//binary.Write(b0, binary.LittleEndian, appeui)
	//binary.Write(b0, binary.LittleEndian, deveui)
	//binary.Write(b0, binary.LittleEndian, devnonce)

	hash, err := cmac.New(appkey)
	if err != nil {
		return false, err
	}

	_, err = hash.Write(b0.Bytes())
	if err != nil {
		return false, err
	}
	calculatedMIC := binary.LittleEndian.Uint32(hash.Sum([]byte{})[0:4])
	return calculatedMIC == mic, nil
}
