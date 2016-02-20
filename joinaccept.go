package lora

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/binary"
	"errors"

	"github.com/jacobsa/crypto/cmac"
)

// JoinAccept ...
type JoinAccept struct {
	mhdr       *MHDR
	appnonce   []byte
	netid      []byte
	devaddr    []byte
	dlsettings byte
	rxdelay    byte
	cflist     []byte
	mic        []byte
}

// NewJoinAccept ...
func NewJoinAccept(appkey []byte, nwkid byte, nwkaddr uint32) (*JoinAccept, error) {
	returnvalue := &JoinAccept{}
	mhdr, err := NewMHDRFromValues(MTypeJoinAccept, MajorLoRaWANR1)
	if err != nil {
		return nil, err
	}
	returnvalue.mhdr = mhdr
	appnonce, err := returnvalue.getAppNonce()
	if err != nil {
		return nil, err
	}
	returnvalue.appnonce = appnonce
	netid, err := returnvalue.getNetID()
	if err != nil {
		return nil, err
	}
	returnvalue.netid = netid
	devaddr, err := returnvalue.getDevAddr(nwkid)
	if err != nil {
		return nil, err
	}
	returnvalue.devaddr = devaddr
	rxdelay, err := returnvalue.getRxDelay()
	if err != nil {
		return nil, err
	}
	returnvalue.rxdelay = rxdelay
	cflist, err := returnvalue.getCFList()
	if err != nil {
		return nil, err
	}
	returnvalue.cflist = cflist
	mic, err := returnvalue.calcMIC(appkey)
	if err != nil {
		return nil, err
	}
	returnvalue.mic = mic
	return returnvalue, nil
}

// Marshal ...
func (joinaccept *JoinAccept) Marshal(appkey []byte) ([]byte, error) {
	block, err := aes.NewCipher(appkey)
	if err != nil {
		return nil, err
	}
	b0 := new(bytes.Buffer)
	b0.Write(joinaccept.appnonce)
	b0.Write(joinaccept.netid)
	b0.Write(joinaccept.devaddr)
	b0.WriteByte(joinaccept.dlsettings)
	b0.WriteByte(joinaccept.rxdelay)
	//	b0.Write(joinaccept.cflist)
	b0.Write(joinaccept.mic)
	plaintext := b0.Bytes()
	bs := block.BlockSize()
	if len(plaintext)%bs != 0 {
		return nil, errors.New("Encrypt Need a multiple of the blocksize")
	}
	ciphertext := make([]byte, len(plaintext))
	ct := ciphertext
	for len(plaintext) > 0 {
		block.Decrypt(ciphertext, plaintext)
		plaintext = plaintext[bs:]
		ciphertext = ciphertext[bs:]
	}
	b1 := new(bytes.Buffer)
	b1.WriteByte(joinaccept.mhdr.Marshal())
	b1.Write(ct)
	b1.Write(joinaccept.mic)
	return b1.Bytes(), nil
}

func (joinaccept *JoinAccept) getAppNonce() ([]byte, error) {
	returnvalue := make([]byte, 3)
	_, err := rand.Read(returnvalue)
	if err != nil {
		return nil, err
	}
	return returnvalue, nil
}

// getNetID: Some constant? 0 for now
func (joinaccept *JoinAccept) getNetID() ([]byte, error) {
	b := make([]byte, 3)
	b[0] = 0
	b[1] = 0
	b[2] = 0
	return b, nil
}

func (joinaccept *JoinAccept) getDevAddr(nwkid byte, nwkaddr uint32) ([]byte, error) {
	a := (uint32)(nwkid*16777216) + nwkaddr
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, a)
	return b, nil
}

func (joinaccept *JoinAccept) getDLSettings() (byte, error) {
	return 0 + 0 + 7, nil // 7 = max datarate ?
}

func (joinaccept *JoinAccept) getRxDelay() (byte, error) {
	return 0, nil
}

func (joinaccept *JoinAccept) getCFList() ([]byte, error) {
	b := make([]byte, 16)
	b[0] = 0
	b[1] = 0
	b[2] = 0

	b[3] = 0
	b[4] = 0
	b[5] = 0

	b[6] = 0
	b[7] = 0
	b[8] = 0

	b[9] = 0
	b[10] = 0
	b[11] = 0

	b[12] = 0
	b[13] = 0
	b[14] = 0

	b[15] = 0

	return b, nil
}

func (joinaccept *JoinAccept) calcMIC(appkey []byte) ([]byte, error) {
	b0 := new(bytes.Buffer)
	b0.WriteByte(joinaccept.mhdr.Marshal())
	b0.Write(joinaccept.appnonce)
	b0.Write(joinaccept.netid)
	b0.Write(joinaccept.devaddr)
	b0.WriteByte(joinaccept.dlsettings)
	b0.WriteByte(joinaccept.rxdelay)
	//	b0.Write(joinaccept.cflist)

	hash, err := cmac.New(appkey)
	if err != nil {
		return nil, err
	}

	_, err = hash.Write(b0.Bytes())
	if err != nil {
		return nil, err
	}

	calculatedMIC := hash.Sum([]byte{})[0:4]
	return calculatedMIC, nil
}
