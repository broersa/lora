package lora

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
	returnvalue := &MHDR{mtype: mhdr >> 5, major: mhdr & 0x3}
	if !(returnvalue.mtype == MTypeJoinRequest ||
		returnvalue.mtype == MTypeConfirmedDataUp ||
		returnvalue.mtype == MTypeUnconfirmedDataUp) {
		return nil, NewErrorMTypeValidationFailed()
	}
	if returnvalue.major != MajorLoRaWANR1 {
		return nil, NewErrorMajorValidationFailed()
	}
	return returnvalue, nil
}

// NewMHDRFromValues ...
func NewMHDRFromValues(mtype byte, major byte) (*MHDR, error) {
	returnvalue := &MHDR{mtype: mtype, major: major}
	if !(mtype == MTypeJoinAccept ||
		returnvalue.mtype == MTypeConfirmedDataDown ||
		returnvalue.mtype == MTypeUnconfirmedDataDown) {
		return nil, NewErrorMTypeValidationFailed()
	}
	if major != MajorLoRaWANR1 {
		return nil, NewErrorMajorValidationFailed()
	}
	return returnvalue, nil
}

// Marshal ...
func (mhdr *MHDR) Marshal() byte {
	return (mhdr.mtype << 5) + mhdr.major
}

// IsJoinRequest ...
func (mhdr *MHDR) IsJoinRequest() bool {
	return mhdr.mtype == MTypeJoinRequest
}

// IsJoinAccept ...
func (mhdr *MHDR) IsJoinAccept() bool {
	return mhdr.mtype == MTypeJoinAccept
}
