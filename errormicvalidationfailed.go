package lora

// ErrorMICValidationFailed ...
type ErrorMICValidationFailed struct {
}

func (errormicvalidationfailed *ErrorMICValidationFailed) Error() string {
	return "MIC validation failed"
}

// NewErrorMICValidationFailed ...
func NewErrorMICValidationFailed() *ErrorMICValidationFailed {
	return &ErrorMICValidationFailed{}
}
