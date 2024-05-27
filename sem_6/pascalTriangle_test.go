package sem_6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPascalTriangle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive #1",
			args: args{
				n: 1,
			},
			want: []int{1},
		},
		{
			name: "Positive #2",
			args: args{
				n: 2,
			},
			want: []int{1, 1},
		},
		{
			name: "Positive #3",
			args: args{
				n: 3,
			},
			want: []int{1, 2, 1},
		},
		{
			name: "Positive #4",
			args: args{
				n: 4,
			},
			want: []int{1, 3, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PascalTriangle(tt.args.n)

			require.Equal(t, tt.want, got)
		})
	}
}
