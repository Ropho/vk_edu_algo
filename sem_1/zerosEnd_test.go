package sem_1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZerosEnd(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive 1",
			args: args{
				arr: []int{0, 0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0, 0},
		},
		{
			name: "Positive 2",
			args: args{
				arr: []int{0, 33, 57, 88, 60, 0, 0, 80, 99},
			},
			want: []int{33, 57, 88, 60, 80, 99, 0, 0, 0},
		},
		{
			name: "Positive 3",
			args: args{
				arr: []int{0, 0, 0, 18, 16, 0, 0, 77, 99},
			},
			want: []int{18, 16, 77, 99, 0, 0, 0, 0, 0},
		},
		{
			name: "Positive 4",
			args: args{
				arr: nil,
			},
			want: nil,
		},
		{
			name: "Positive 5",
			args: args{
				arr: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ZerosEnd(tt.args.arr)

			require.Equal(t, tt.want, got)
		})
	}
}
