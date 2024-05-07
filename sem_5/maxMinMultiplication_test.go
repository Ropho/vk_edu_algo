package sem_5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxMinMultiplication(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive #1",
			args: args{
				data: []int{16, 12, 18, 11, 15, 17, 21, 0, 0, 0, 0, 0, 0, 19, 24},
			},
			want: 264,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := MaxMinMultiplication(tt.args.data)

			require.Equal(t, tt.want, got)
		})
	}
}
