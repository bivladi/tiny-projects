package gordle

type corpusError string

// Error is the implementation of the error interface
func (e corpusError) Error() string {
	return string(e)
}