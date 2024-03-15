package sem_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasCycle(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive #1",
			args: args{
				head: ConstructList([]Note{
					{
						Val:       11,
						NextIndex: 1,
					},
					{
						Val:       3,
						NextIndex: 2,
					},
					{
						Val:       8,
						NextIndex: 3,
					},
					{
						Val:       9,
						NextIndex: 4,
					},
					{
						Val:       6,
						NextIndex: 5,
					},
					{
						Val:       11,
						NextIndex: 6,
					},
					{
						Val:       16,
						NextIndex: 7,
					},
					{
						Val:       17,
						NextIndex: 2,
					},
				},
				),
			},
			want: true,
		},
		{
			name: "Positive #2",
			args: args{
				head: ConstructList([]Note{
					{
						Val:       11,
						NextIndex: 1,
					},
					{
						Val:       3,
						NextIndex: 2,
					},
					{
						Val:       8,
						NextIndex: 3,
					},
					{
						Val:       9,
						NextIndex: 4,
					},
					{
						Val:       6,
						NextIndex: 5,
					},
					{
						Val:       11,
						NextIndex: 6,
					},
					{
						Val:       16,
						NextIndex: 7,
					},
					{
						Val:       17,
						NextIndex: -1,
					},
				},
				),
			},
			want: false,
		},
		{
			name: "Positive #3",
			args: args{
				head: ConstructList([]Note{
					{
						Val:       11,
						NextIndex: -1,
					},
				},
				),
			},
			want: false,
		},
		{
			name: "Positive #4",
			args: args{
				head: ConstructList([]Note{}),
			},
			want: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			hasCycle := HasCycle(tt.args.head)

			require.Equal(t, hasCycle, tt.want)
		})

	}
}
