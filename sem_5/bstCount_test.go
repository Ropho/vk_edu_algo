package sem_5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBSTCount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive #1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "Positive #2",
			args: args{
				n: 0,
			},
			want: 1,
		},
		{
			name: "Positive #3",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "Positive #4",
			args: args{
				n: 10,
			},
			want: 16796,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := BSTCount(tt.args.n)

			require.Equal(t, got, tt.want)
		})
	}
}
