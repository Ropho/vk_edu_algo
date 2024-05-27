package sem_6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCoins(t *testing.T) {
	type args struct {
		coins  []int
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
				coins:  []int{1, 2, 5},
				target: 11,
			},
			want: 3,
		},
		{
			name: "Positive #2",
			args: args{
				coins:  []int{1, 2, 4},
				target: 13,
			},
			want: 4,
		},
		{
			name: "Negative #1",
			args: args{
				coins:  []int{2},
				target: 3,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Coins(tt.args.coins, tt.args.target)

			require.Equal(t, tt.want, got)
		})
	}
}
