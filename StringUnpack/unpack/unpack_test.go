package unpack

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	type args struct {
		str   string
		isRaw bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name:    "Valid string provided. not Raw",
			args:    args{"\n4a4b1c2d5e0fgh4i0", false},
			want:    `"\n\n\n\naaaabccdddddfghhhh"`,
			wantErr: nil,
		},
		{
			name:    "String contains number >9. Not raw",
			args:    args{"a4b21", false},
			wantErr: errors.New("error. String cannot contain numbers >9"),
		},
		{
			name:    "Valid string provided. Raw",
			args:    args{`\04\24\\\3b3c1d0\\2a0\4b\2\10`, true},
			want: `00002222\3bbbc\\4b2`,
		},
		{
			name:    "String starting with digit",
			args:    args{`4jksdfio`, true},
			wantErr: errors.New("error. String cannot start with digit"),
		},
		{
			name:    "String contains number >9. Raw",
			args:    args{`a4b21`, true},
			wantErr: errors.New("error. String cannot contain numbers >9"),
		},
		{
			name:    "Escape letters in raw mode. Raw",
			args:    args{`\a4b21`, true},
			wantErr: errors.New("error. cannot escape letter in escaping mode"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unpack(tt.args.str, tt.args.isRaw)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Unpack() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Unpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
