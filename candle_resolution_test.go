package finnhub

import (
	"testing"
)

func TestCandleResolution_String(t *testing.T) {
	tests := []struct {
		name string
		d    CandleResolution
		want string
	}{
		{
			name: "valid",
			d: CandleResolution5Second,
			want: "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}