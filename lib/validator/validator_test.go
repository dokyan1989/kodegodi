package validator

import (
	"fmt"
	"testing"
)

func TestIsJSONString(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "",
			want:  false,
		},
		{
			input: "1",
			want:  true,
		},
		{
			input: `"abc"`,
			want:  true,
		},
		{
			input: "{}",
			want:  true,
		},
		{
			input: `{a: 1}`,
			want:  false,
		},
		{
			input: `{"a": "1"`,
			want:  false,
		},
		{
			input: `{"a": "1"}`,
			want:  true,
		},
		{
			input: `[{"a": "1"}]`,
			want:  true,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Input: %v", tt.input)
		t.Run(name, func(t *testing.T) {
			if got := IsJSONString(tt.input); got != tt.want {
				t.Errorf("IsJSONString() = %v, want %v", got, tt.want)
			}
		})
	}
}
