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
func NewJoinRequest(joinrequest []byte) (*JoinRequest, error) {
	mhdr, err := NewMHDRFromByte(joinrequest[0])
	if err != nil {
		return nil, err
	}
	if !mhdr.IsJoinRequest() {
		return nil, errors.New("Not a join request message")
	}
	returnvalue := &JoinRequest{mhdr, joinrequest[1:9], joinrequest[9:17], joinrequest[17:19], joinrequest[19:23]}
	return returnvalue, nil
}

// NewJoinRequestValidated ...
func NewJoinRequestValidated(appkey []byte, joinrequest []byte) (*JoinRequest, error) {
	returnvalue, err := NewJoinRequest(joinrequest)
	if err != nil {
		return nil, err
	}
	valid, err := returnvalue.validateJoinRequest(appkey)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, NewErrorMICValidationFailed()
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
	b0.WriteByte(joinrequest.mhdr.Marshal())
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
	return bytes.Equal(calculatedMIC, joinrequest.mic), nil
}
