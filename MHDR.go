package lora

//	"fmt"

const (
	JoinRequest         byte = 0
	JoinAccept          byte = 1
	UnconfirmedDataUp   byte = 2
	UnconfirmedDataDown byte = 3
	ConfirmedDataUp     byte = 4
	ConfirmedDataDown   byte = 5
)

const (
	LoRaWANR1 byte = 0
)

type MHDR struct {
	MType byte
	Major byte
}

func ParseMHDR(b byte) (*MHDR, error) {
	return &MHDR{b >> 5, b & 0x3}, nil
}

func MarshallMHDR(mhdr *MHDR) (byte, error) {
	return (mhdr.MType << 5) + mhdr.Major, nil
}
