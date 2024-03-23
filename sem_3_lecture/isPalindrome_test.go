package sem3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindromeStack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive #1",
			args: args{
				s: "aaadfdaaa",
			},
			want: true,
		},
		{
			name: "Positive #2",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "Positive #3",
			args: args{
				s: "qwerty",
			},
			want: false,
		},
		{
			name: "Positive #4",
			args: args{
				s: "😀🤣😇🥰🥰😇🤣😀",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := IsPalindromeStack(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestIsPalindromePointers(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive #1",
			args: args{
				s: "aaadfdaaa",
			},
			want: true,
		},
		{
			name: "Positive #2",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "Positive #3",
			args: args{
				s: "qwerty",
			},
			want: false,
		},
		{
			name: "Positive #4",
			args: args{
				s: "😀🤣😇🥰🥰😇🤣😀",
			},
			want: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := IsPalindromePointers(tt.args.s)
			require.Equal(t, tt.want, got)

		})
	}
}
