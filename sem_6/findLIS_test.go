package sem_6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive #1",
			args: args{
				nums: []int{3, 2, 8, 9, 5, 10},
			},
			want: 3,
		},
		{
			name: "Positive #2",
			args: args{
				nums: []int{1, 2, 7, 9, 0, 10},
			},
			want: 4,
		},
		{
			name: "Positive #3",
			args: args{
				nums: []int{8, 8, 8, 8},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLIS(tt.args.nums)

			require.Equal(t, tt.want, got)
		})
	}
}
