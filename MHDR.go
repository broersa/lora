package lora

import "errors"

const (
	// MajorLoRaWANR1 major
	MajorLoRaWANR1 byte = 0
	// MTypeJoinRequest messagetype
	MTypeJoinRequest byte = 0
	// MTypeJoinAccept messagetype
	MTypeJoinAccept byte = 1
	// MTypeUnconfirmedDataUp messagetype
	MTypeUnconfirmedDataUp byte = 2
	// MTypeUnconfirmedDataDown messagetype
	MTypeUnconfirmedDataDown byte = 3
	// MTypeConfirmedDataUp messagetype
	MTypeConfirmedDataUp byte = 4
	// MTypeConfirmedDataDown messagetype
	MTypeConfirmedDataDown byte = 5
)

// MHDR message
type MHDR struct {
	mtype byte
	major byte
}

// NewMHDRFromByte ...
func NewMHDRFromByte(mhdr byte) (*MHDR, error) {
	returnvalue := &MHDR{mtype: b >> 5, major: b & 0x3}
	if !(returnvalue.mtype == MTypeJoinRequest ||
		returnvalue.mtype == MTypeConfirmedDataUp ||
		returnvalue.mtype == MTypeUnconfirmedDataUp) {
		return nil, errors.New("Not a valid MType")
	}
	if mhdr.Major != MajorLoRaWANR1 {
		return nil, errors.New("Not a valic Major")
	}
	return returnvalue, nil
}

// NewMHDRFromValues ...
func NewMHDRFromValues(mtype byte, major byte) (*MHDR, error) {
	returnvalue := &MHDR{mtype: b >> 5, major: b & 0x3}
	if !(returnvalue.mtype == MTypeJoinAccept ||
		returnvalue.mtype == MTypeConfirmedDataDown ||
		returnvalue.mtype == MTypeUnconfirmedDataDown) {
		return nil, errors.New("Not a valid MType")
	}
	if major != MajorLoRaWANR1 {
		return nil, errors.New("Not a valic Major")
	}
	return returnvalue, nil
}

// Marshal ...
func (mhdr *MHDR) Marshal() byte {
	return (mhdr.MType << 5) + mhdr.Major
}

// IsJoinRequest ...
func (mhdr *MHDR) IsJoinRequest() bool {
	return mhdr.MType == MTypeJoinRequest
}

// IsJoinAccept ...
func (mhdr *MHDR) IsJoinAccept() bool {
	return mhdr.MType == MTypeJoinAccept
}
