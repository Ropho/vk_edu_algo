package sem_4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtraLetter(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Positive #1",
			args: args{
				first:  "uio",
				second: "oeui",
			},
			want: 'e',
		},
		{
			name: "Positive #2",
			args: args{
				first:  "fe",
				second: "efo",
			},
			want: 'o',
		},
		{
			name: "Positive #3",
			args: args{
				first:  "ab",
				second: "ab",
			},
			want: 0,
		},
		{
			name: "Positive #4",
			args: args{
				first:  "bbb",
				second: "bbbb",
			},
			want: 'b',
		},
		{
			name: "Positive #5",
			args: args{
				first:  "",
				second: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := ExtraLetter(tt.args.first, tt.args.second)

			require.Equal(t, tt.want, got)
		})
	}
}
