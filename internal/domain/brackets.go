package domain

type BracketPair struct {
	Open  rune
	Close rune
}

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

func (r ValidationResult) Status() string {
	if r.Valid {
		return "true"
	}

	return "false"
}
