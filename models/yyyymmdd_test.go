package models

import (
	"testing"
)

func TestYyyyMmDdDashed_UnmarshalJSON(t *testing.T) {
	type args struct {
		inputDateString string
	}
	tests := []struct {
		name     string
		expected YyyyMmDdDashed
		args     args
		wantErr  bool
	}{
		{
			name:     "2021-05-01 is a valid YyyyMmDdDashed type",
			expected: NewYyyyMmDdDashed(2021, 5, 1),
			args: args{
				inputDateString: "2021-05-01",
			},
			wantErr: false,
		},
		{
			name:     "20210501 is NOT a valid YyyyMmDdDashed type",
			expected: NewYyyyMmDdDashed(2021, 5, 1),
			args: args{
				inputDateString: "20210501",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.expected.UnmarshalJSON([]byte(tt.args.inputDateString)); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
