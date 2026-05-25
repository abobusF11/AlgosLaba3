package usecase

import (
	"fmt"

	"brackets-lab/internal/domain"
)

// Validator checks bracket sequences.
type Validator struct {
	pairs        []domain.BracketPair
	openToClose  map[rune]rune
	closeToOpen  map[rune]rune
	ignoreOthers bool
}

// Option configures Validator.
type Option func(*Validator)

// WithAngleBrackets enables validation for angle brackets: < and >.
func WithAngleBrackets() Option {
	return func(v *Validator) {
		v.addPair('<', '>')
	}
}

// WithIgnoreOthers configures whether Validator should ignore non-bracket symbols.
func WithIgnoreOthers(ignore bool) Option {
	return func(v *Validator) {
		v.ignoreOthers = ignore
	}
}

// NewValidator creates Validator with default bracket pairs: (), [], {}.
func NewValidator(opts ...Option) *Validator {
	v := &Validator{
		openToClose:  make(map[rune]rune),
		closeToOpen:  make(map[rune]rune),
		ignoreOthers: true,
	}

	v.addPair('(', ')')
	v.addPair('[', ']')
	v.addPair('{', '}')

	for _, opt := range opts {
		opt(v)
	}

	return v
}

// Validate checks one string and returns validation details.
func (v *Validator) Validate(input string) domain.ValidationResult {
	result := domain.ValidationResult{
		Input:    input,
		Valid:    true,
		ErrorPos: -1,
	}

	stack := make([]stackItem, 0, len(input))

	for pos, ch := range []rune(input) {
		if _, ok := v.openToClose[ch]; ok {
			stack = append(stack, stackItem{
				bracket: ch,
				pos:     pos,
			})

			result.OpenCount++

			if len(stack) > result.MaxDepth {
				result.MaxDepth = len(stack)
			}

			continue
		}

		if expectedOpen, ok := v.closeToOpen[ch]; ok {
			result.CloseCount++

			if len(stack) == 0 {
				return invalidResult(
					result,
					pos,
					fmt.Sprintf("лишняя закрывающая скобка %q", ch),
				)
			}

			last := stack[len(stack)-1]

			if last.bracket != expectedOpen {
				expectedClose := v.openToClose[last.bracket]

				return invalidResult(
					result,
					pos,
					fmt.Sprintf(
						"ожидалась закрывающая скобка %q для %q, получена %q",
						expectedClose,
						last.bracket,
						ch,
					),
				)
			}

			stack = stack[:len(stack)-1]
			result.PairsCount++

			continue
		}

		if !v.ignoreOthers {
			return invalidResult(
				result,
				pos,
				fmt.Sprintf("недопустимый символ %q", ch),
			)
		}
	}

	if len(stack) > 0 {
		last := stack[len(stack)-1]
		expectedClose := v.openToClose[last.bracket]

		return invalidResult(
			result,
			last.pos,
			fmt.Sprintf("не хватает закрывающей скобки %q", expectedClose),
		)
	}

	return result
}

// ValidateMany checks several strings and returns results in the same order.
func (v *Validator) ValidateMany(inputs []string) []domain.ValidationResult {
	results := make([]domain.ValidationResult, 0, len(inputs))

	for _, input := range inputs {
		results = append(results, v.Validate(input))
	}

	return results
}

func (v *Validator) addPair(open, close rune) {
	v.pairs = append(v.pairs, domain.BracketPair{
		Open:  open,
		Close: close,
	})

	v.openToClose[open] = close
	v.closeToOpen[close] = open
}

type stackItem struct {
	bracket rune
	pos     int
}

func invalidResult(
	result domain.ValidationResult,
	pos int,
	reason string,
) domain.ValidationResult {
	result.Valid = false
	result.ErrorPos = pos
	result.ErrorReason = reason

	return result
}
