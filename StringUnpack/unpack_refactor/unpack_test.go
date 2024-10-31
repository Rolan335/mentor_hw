package unpack_refactor

import (
	"errors"
	"testing"
)

func BenchmarkUnpackRawMode(b *testing.B) {
	input := `\04\24\\\3b3c1d0\\2a0\4b\2\10`
	for i := 0; i < b.N; i++ {
		_, err := Unpack(input, true)
		if err != nil {
			b.Fatal(err)
		}
	}
}
//BenchmarkUnpackRawMode-12    	 1372424	       863.6 ns/op	     144 B/op	       9 allocs/op


func BenchmarkPack(b *testing.B) {
	input := `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbjjjjjkuwqorjoajsfnjjjjjjoooooiiiiiewklllllllfnnnnnnndjdddjkll`
	for i := 0; i < b.N; i++ {
		_, err := Pack(input)
		if err != nil {
			b.Fatal(err)
		}
	}
}
//BenchmarkPack-12          100000              1848 ns/op             776 B/op         12 allocs/op

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
			name: "Empty string provided. Any mode.",
			args: args{"", false},
			want: "",
		},
		{
			name:    "Valid string provided. not Raw",
			args:    args{"\n4a4b1c2d5e0fgh4i0", false},
			want:    `"\n\n\n\naaaabccdddddfghhhh"`,
			wantErr: nil,
		},
		{
			name:    "String contains number >9. Not raw",
			args:    args{"a4b21", false},
			wantErr: errors.New("cannot contain numbers >9 or 00, 01, 02, etc"),
		},
		{
			name: "Valid string provided. Raw",
			args: args{`\04\24\\\3b3c1d0\\2a0\4b\2\10`, true},
			want: `00002222\3bbbc\\4b2`,
		},
		{
			name:    "String starting with digit",
			args:    args{`4jksdfio`, true},
			wantErr: errors.New("cannot start with digit"),
		},
		{
			name:    "String contains number >9. Raw",
			args:    args{`a4b21`, true},
			wantErr: errors.New("cannot contain numbers >9 or 00, 01, 02, etc"),
		},
		{
			name:    "Escape letters at the end in raw mode. Raw",
			args:    args{`b2\1\n`, true},
			wantErr: errors.New("cannot escape letter in escaping mode"),
		},
		{
			name:    "Escape letters at the start in raw mode. Raw",
			args:    args{`\nb2\1`, true},
			wantErr: errors.New("cannot escape letter in escaping mode"),
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
