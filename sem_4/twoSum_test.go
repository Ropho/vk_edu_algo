package sem_4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		data   []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want1 int
		want2 int
	}{
		{
			name: "Positive #1",
			args: args{
				data:   []int{10, 1, 2, 3, 4},
				target: 4,
			},
			want1: 1,
			want2: 3,
		},
		{
			name: "Negative #1",
			args: args{
				data:   []int{10, 1, 2, 3, 4},
				target: 2,
			},
			want1: 0,
			want2: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := twoSum(tt.args.data, tt.args.target)

			require.Equal(t, got1, tt.want1)
			require.Equal(t, got2, tt.want2)
		})
	}
}
