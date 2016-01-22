package lora

// ErrorMajorValidationFailed ...
type ErrorMajorValidationFailed struct {
}

func (errormajorvalidationfailed *ErrorMajorValidationFailed) Error() string {
	return "Major validation failed"
}

// NewErrorMajorValidationFailed ...
func NewErrorMajorValidationFailed() *ErrorMajorValidationFailed {
	return &ErrorMajorValidationFailed{}
}
