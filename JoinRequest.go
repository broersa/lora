package lora

import (
	"bytes"
	"errors"

	"github.com/jacobsa/crypto/cmac"
)

// JoinRequest ...
type JoinRequest struct {
	mhdr     *MHDR
	appeui   []byte
	deveui   []byte
	devnonce []byte
	mic      []byte
}

// NewJoinRequest ...
func NewJoinRequest(appkey []byte, joinrequest []byte) (*JoinRequest, error) {
	mhdr, err := NewMHDRFromByte(joinrequest[0])
	if err != nil {
		return nil, err
	}
	returnvalue := &JoinRequest{mhdr, joinrequest[1:9], joinrequest[9:17], joinrequest[17:19], joinrequest[19:23]}
	valid, err := returnvalue.validateJoinRequest(appkey)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, errors.New("MIC validation failed")
	}
	return returnvalue, nil
}

// GetMHDR ...
func (joinrequest *JoinRequest) GetMHDR() *MHDR {
	return joinrequest.mhdr
}

// GetDevNonce ...
func (joinrequest *JoinRequest) GetDevNonce() []byte {
	return joinrequest.devnonce
}

func (joinrequest *JoinRequest) validateJoinRequest(appkey []byte) (bool, error) {
	b0 := new(bytes.Buffer)
	b0.WriteByte(joinrequest.mhdr)
	b0.Write(joinrequest.appeui)
	b0.Write(joinrequest.deveui)
	b0.Write(joinrequest.devnonce)
	hash, err := cmac.New(appkey)
	if err != nil {
		return false, err
	}
	_, err = hash.Write(b0.Bytes())
	if err != nil {
		return false, err
	}
	calculatedMIC := hash.Sum([]byte{})[0:4]
	return calculatedMIC == joinrequest.mic, nil
}
