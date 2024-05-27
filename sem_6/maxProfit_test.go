package sem_6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive #1",
			args: args{
				prices: []int{8, 9, 3, 7, 4, 16, 12},
			},
			want: 13,
		},
		{
			name: "Positive #2",
			args: args{
				prices: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: 6,
		},
		{
			name: "Positive #3",
			args: args{
				prices: []int{8, 7, 6, 5, 4, 3, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.args.prices)

			require.Equal(t, tt.want, got)
		})
	}
}
