package lora

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/binary"
	"errors"

	"github.com/jacobsa/crypto/cmac"
)

func JoinAcceptGetAppNonce() ([]byte, error) {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func JoinAcceptGetNetID(nwkid byte) ([]byte, error) {
	b := make([]byte, 3)
	b[0] = 0
	b[1] = 0
	b[2] = 0
	return b, nil
}

func JoinAcceptGetDevAddr(nwkid byte) (uint32, error) {
	b := make([]byte, 4)
	b[0] = 0
	b[1] = 0
	b[2] = 0
	b[3] = 0
	return binary.LittleEndian.Uint32(b), nil
}

func JoinAcceptGetDLSettings() (byte, error) {
	return 0 + 0 + 7, nil // 7 = max datarate ?
}

func JoinAcceptGetRxDelay() (byte, error) {
	return 0, nil
}

func JoinAcceptGetCFList() ([]byte, error) {
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

func JoinAcceptCalcMIC(appkey []byte, mhdr byte, appnonce []byte, netid []byte, devaddr uint32, dlsettings byte, rxdelay byte, cflist []byte) ([]byte, error) {
	b0 := new(bytes.Buffer)
	b0.WriteByte(mhdr)
	b0.Write(appnonce)
	b0.Write(netid)
	binary.Write(b0, binary.LittleEndian, devaddr)
	b0.WriteByte(dlsettings)
	b0.WriteByte(rxdelay)
	//b0.Write(cflist)

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

func JoinAcceptEncrypt(appkey []byte, appnonce []byte, netid []byte, devaddr uint32, dlsettings byte, rxdelay byte, cflist []byte, mic []byte) ([]byte, error) {
	block, err := aes.NewCipher(appkey)
	if err != nil {
		return nil, err
	}
	b0 := new(bytes.Buffer)
	b0.Write(appnonce)
	b0.Write(netid)
	binary.Write(b0, binary.LittleEndian, devaddr)
	b0.WriteByte(dlsettings)
	b0.WriteByte(rxdelay)
	//b0.Write(cflist)
	b0.Write(mic)
	plaintext := b0.Bytes()
	//pt := plaintext

	bs := block.BlockSize()
	if len(plaintext)%bs != 0 {
		return nil, errors.New("JoinAccept: Encrypt Need a multiple of the blocksize")
	}

	ciphertext := make([]byte, len(plaintext))
	ct := ciphertext
	for len(plaintext) > 0 {
		block.Decrypt(ciphertext, plaintext)
		plaintext = plaintext[bs:]
		ciphertext = ciphertext[bs:]
	}
	return ct, nil
}
