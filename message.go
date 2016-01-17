package lora

import "errors"

// GetMessageType ...
func GetMessageType(data []byte) (byte, error) {
	mhdr := ParseMHDR(data[0])
	if mhdr.Major != LoRaWANR1 {
		return 0, errors.New("Invalid message major")
	}
	return mhdr.MType, nil
}
