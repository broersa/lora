package lora

// ErrorMTypeValidationFailed ...
type ErrorMTypeValidationFailed struct {
}

func (errormtypevalidationfailed *ErrorMTypeValidationFailed) Error() string {
	return "MType validation failed"
}

// NewErrorMTypeValidationFailed ...
func NewErrorMTypeValidationFailed() *ErrorMTypeValidationFailed {
	return &ErrorMTypeValidationFailed{}
}
