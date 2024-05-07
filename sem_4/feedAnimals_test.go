package sem_4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFeedAnimals(t *testing.T) {
	type args struct {
		animals []int
		food    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive #1",
			args: args{
				animals: []int{3, 4, 7},
				food:    []int{8, 1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FeedAnimals(tt.args.animals, tt.args.food)

			require.Equal(t, tt.want, got)
		})
	}
}
