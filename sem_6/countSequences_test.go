package sem_6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountSequences(t *testing.T) {
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
				n: 4,
			},
			want: 13,
		},
		{
			name: "Positive #2",
			args: args{
				n: 3,
			},
			want: 7,
		},
		{
			name: "Positive #3",
			args: args{
				n: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountSequences(tt.args.n)

			require.Equal(t, tt.want, got)
		})
	}
}
