package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name          string
		a             float64
		b             float64
		expected      float64
		expectedError string
	}{
		{"normal division", 6.0, 3.0, 2.0, ""},
		{"divide by zero", 5.0, 0.0, 0.0, "division by zero is not allowed"},
		{"negative numbers", -10.0, 2.0, -5.0, ""},
		{"zero dividend", 0.0, 5.0, 0.0, ""},
		{"fraction result", 5.0, 2.0, 2.5, ""},
		{"both negative", -15.0, -3.0, 5.0, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			assert.Equal(t, tt.expected, got)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				return
			}

			assert.Nil(t, err)
		})
	}
}
