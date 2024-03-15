package sem_2

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseLinkedList(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
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
						NextIndex: -1,
					},
				},
				),
			},
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
						NextIndex: -1,
					},
				},
				),
			},
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
		},
		{
			name: "Positive #4",
			args: args{
				head: ConstructList([]Note{}),
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			var notReversed []int
			cur := tt.args.head
			for cur != nil {
				notReversed = append(notReversed, cur.Val)
				cur = cur.Next
			}

			newHead := ReverseLinkedList(tt.args.head)

			var reversed []int
			cur = newHead
			for cur != nil {
				reversed = append(reversed, cur.Val)
				cur = cur.Next
			}

			sort.SliceStable(reversed, func(i, j int) bool {
				return true
			})

			require.Equal(t, notReversed, reversed)
		})

	}
}
