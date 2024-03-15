package sem_1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvenFirst(t *testing.T) {
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
				arr: []int{7, 3, 2, 4, 1, 11, 8, 9},
			},
			want: []int{2, 4, 8, 3, 1, 11, 7, 9},
		},
		{
			name: "Positive 2",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{0, 2, 4, 6, 8, 10, 3, 7, 1, 9, 5},
		},
		{
			name: "Positive 3",
			args: args{
				arr: []int{},
			},
			want: []int{},
		},
		{
			name: "Positive 4",
			args: args{
				arr: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := EvenFirst(tt.args.arr)

			require.Equal(t, tt.want, got)
		})
	}
}
