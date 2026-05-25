package usecase

import "testing"

func TestValidatorValidate(t *testing.T) {
	validator := NewValidator(WithAngleBrackets())

	tests := []struct {
		name  string
		input string
		valid bool
	}{
		{name: "round brackets", input: "()", valid: true},
		{name: "nested brackets", input: "([]{})", valid: true},
		{name: "angle brackets", input: "<([])>", valid: true},
		{name: "text with brackets", input: "a + (b * c)", valid: true},
		{name: "wrong order", input: "([)]", valid: false},
		{name: "extra closing", input: "())", valid: false},
		{name: "missing closing", input: "(()", valid: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.Validate(tt.input)

			if result.Valid != tt.valid {
				t.Fatalf("expected valid=%v, got valid=%v, reason=%q", tt.valid, result.Valid, result.ErrorReason)
			}
		})
	}
}
