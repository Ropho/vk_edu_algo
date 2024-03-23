package sem3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleTask(t *testing.T) {
	type args struct {
		copyNum int
		first   int
		second  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive #1",
			args: args{
				copyNum: 4,
				first:   1,
				second:  1,
			},
			want: 3,
		},
		{
			name: "Positive #2",
			args: args{
				copyNum: 5,
				first:   1,
				second:  2,
			},
			want: 4,
		},
		{
			name: "Positive #3",
			args: args{
				copyNum: 1,
				first:   1,
				second:  2,
			},
			want: 1,
		},
		{
			name: "Negative #1",
			args: args{
				copyNum: 5,
				first:   -1,
				second:  2,
			},
			want: -1,
		},
		{
			name: "Negative #2",
			args: args{
				copyNum: 0,
				first:   1,
				second:  2,
			},
			want: -1,
		},
		{
			name: "Negative #3",
			args: args{
				copyNum: 5,
				first:   0,
				second:  -2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := SimpleTask(tt.args.copyNum, tt.args.first, tt.args.second)

			require.Equal(t, tt.want, got)
		})
	}
}
