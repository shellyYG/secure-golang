package null_byte

import "strings"

type Validator struct{}

func NewValidator() *Validator{
	return &Validator{}
}
// Note that http://localhost:8084?input=%00 is a null byte 
// It does not look like a null byte, because in http request, we need to encode it in UTF-8 format in order to send it
func (iv *Validator) ContainsNullByte(input string) bool {
	if input == "" {
		return false
	}
	return strings.Contains(input,"\x00")	
}