package unpack

import (
	"errors"
	"testing"
)

func TestPack(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Valid string provided",
			args: args{"zzaaaabbbbcczccddddeeeed"},
			want: "z2a4b4c2z1c2d4e4d1",
		},
		{
			name: "Valid string. letter observed > 9 times",
			args: args{"aaaaaaaaaaaaaaaaaaaaaaaaavbbccc"},
			want: "a9a9a7v1b2c3",
		},
		{
			name:    "String with digit.",
			args:    args{"ksaf2"},
			wantErr: errors.New("error. Cannot use digits"),
		},
		{
			name: "Empty string",
			args: args{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pack(tt.args.str)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Pack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pack() = %v, want %v", got, tt.want)
			}
		})
	}
}
