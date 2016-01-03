package lora

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/jacobsa/crypto/cmac"
)

type joinrequest struct {
	AppEUI   uint64
	DevEUI   uint64
	DevNonce uint16
}

func ParseJoinRequest(mhdr byte, macpayload []byte, mic uint32) *joinrequest {
	return &joinrequest{binary.LittleEndian.Uint64(macpayload[0:8]),
		binary.LittleEndian.Uint64(macpayload[8:16]),
		binary.LittleEndian.Uint16(macpayload[16:18])}
}

func ValidateJoinRequest(appkey []byte, mhdr byte, x []byte /* appeui uint64, deveui uint64, devnonce uint16,*/, mic uint32) (bool, error) {
	b0 := new(bytes.Buffer)
	b0.WriteByte(mhdr)
	b0.Write(x)
	//binary.Write(b0, binary.LittleEndian, appeui)
	//binary.Write(b0, binary.LittleEndian, deveui)
	//binary.Write(b0, binary.LittleEndian, devnonce)

	hash, err := cmac.New(appkey)
	if err != nil {
		//log.Printf("Failed to initialize CMAC: %s", err.Error())
		return false, err
	}

	_, err = hash.Write(b0.Bytes())
	if err != nil {
		//log.Printf("Failed to hash data: %s", err.Error())
		return false, err
	}
	fmt.Printf("%v\n", hash.Sum([]byte{}))

	calculatedMIC := binary.LittleEndian.Uint32(hash.Sum([]byte{})[0:4])
	fmt.Printf("%d\n", calculatedMIC)
	return calculatedMIC == mic, nil
}
