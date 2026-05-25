package domain

// BracketPair describes one opening and closing bracket pair.
type BracketPair struct {
	Open  rune
	Close rune
}

// ValidationResult contains the result of bracket sequence validation.
type ValidationResult struct {
	Input       string
	Valid       bool
	ErrorPos    int
	ErrorReason string
	PairsCount  int
	MaxDepth    int
	OpenCount   int
	CloseCount  int
}

// Status returns a user-friendly validation status.
func (r ValidationResult) Status() string {
	if r.Valid {
		return "true"
	}

	return "false"
}
