package calc

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		b           float64
		want        float64
		expectError bool
	}{
		{"normal division", 6.0, 3.0, 2.0, false},
		{"divide by zero", 5.0, 0.0, 0.0, true},
		{"negative numbers", -10.0, 2.0, -5.0, false},
		{"zero dividend", 0.0, 5.0, 0.0, false},
		{"fraction result", 5.0, 2.0, 2.5, false},
		{"both negative", -15.0, -3.0, 5.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
