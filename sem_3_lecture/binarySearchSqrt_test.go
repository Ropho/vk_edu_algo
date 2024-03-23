package sem3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearchSqrt(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive #1",
			args: args{
				target: 4,
			},
			want: 2,
		},
		{
			name: "Positive #2",
			args: args{
				target: 10,
			},
			want: 3,
		},
		{
			name: "Positive #3",
			args: args{
				target: 14,
			},
			want: 4,
		},
		{
			name: "Positive #5",
			args: args{
				target: 0,
			},
			want: 0,
		},
		{
			name: "Positive #6",
			args: args{
				target: 9781500,
			},
			want: 3128,
		},
		{
			name: "Positive #7",
			args: args{
				target: 9781132,
			},
			want: 3127,
		},
		{
			name: "Negative #1",
			args: args{
				target: -10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := BinarySearchSqrt(tt.args.target)
			require.Equal(t, tt.want, got)

		})
	}
}
