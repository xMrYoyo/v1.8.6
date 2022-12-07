package converter

type directStringPubkeyConverter struct{}

// NewDirectStringPubkeyConverter creates a new instance of a direct conversion from string to byte converter
func NewDirectStringPubkeyConverter() *directStringPubkeyConverter {
	return &directStringPubkeyConverter{}
}

// Decode decodes the string as its representation in bytes
func (converter *directStringPubkeyConverter) Decode(humanReadable string) ([]byte, error) {
	return []byte(humanReadable), nil
}

// Encode encodes a byte array in its string representation
func (converter *directStringPubkeyConverter) Encode(pkBytes []byte) (string, error) {
	return string(pkBytes), nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (converter *directStringPubkeyConverter) IsInterfaceNil() bool {
	return converter == nil
}
