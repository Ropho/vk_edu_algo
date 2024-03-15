package sem_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeSortedLists(t *testing.T) {
	type args struct {
		head1 *Node
		head2 *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Positive #1",
			args: args{
				head1: ConstructList([]Note{
					{
						Val:       3,
						NextIndex: 1,
					},
					{
						Val:       6,
						NextIndex: 2,
					},
					{
						Val:       8,
						NextIndex: -1,
					},
				},
				),
				head2: ConstructList([]Note{
					{
						Val:       4,
						NextIndex: 1,
					},
					{
						Val:       7,
						NextIndex: 2,
					},
					{
						Val:       9,
						NextIndex: 3,
					},
					{
						Val:       11,
						NextIndex: -1,
					},
				},
				),
			},
			want: ConstructList([]Note{
				{
					Val:       3,
					NextIndex: 1,
				},
				{
					Val:       4,
					NextIndex: 2,
				},
				{
					Val:       6,
					NextIndex: 3,
				},
				{
					Val:       7,
					NextIndex: 4,
				},
				{
					Val:       8,
					NextIndex: 5,
				},
				{
					Val:       9,
					NextIndex: 6,
				},
				{
					Val:       11,
					NextIndex: -1,
				},
			},
			),
		},
		{
			name: "Positive #2",
			args: args{
				head1: ConstructList([]Note{
					{
						Val:       11,
						NextIndex: -1,
					},
				},
				),
				head2: nil,
			},
			want: ConstructList([]Note{
				{
					Val:       11,
					NextIndex: -1,
				},
			},
			),
		},
		{
			name: "Positive #3",
			args: args{
				head1: ConstructList([]Note{
					{
						Val:       3,
						NextIndex: 1,
					},
					{
						Val:       3,
						NextIndex: 2,
					},
					{
						Val:       3,
						NextIndex: -1,
					},
				},
				),
				head2: ConstructList([]Note{
					{
						Val:       4,
						NextIndex: 1,
					},
					{
						Val:       7,
						NextIndex: 2,
					},
					{
						Val:       9,
						NextIndex: 3,
					},
					{
						Val:       11,
						NextIndex: -1,
					},
				},
				),
			},
			want: ConstructList([]Note{
				{
					Val:       3,
					NextIndex: 1,
				},
				{
					Val:       3,
					NextIndex: 2,
				},
				{
					Val:       3,
					NextIndex: 3,
				},
				{
					Val:       4,
					NextIndex: 4,
				},
				{
					Val:       7,
					NextIndex: 5,
				},
				{
					Val:       9,
					NextIndex: 6,
				},
				{
					Val:       11,
					NextIndex: -1,
				},
			},
			),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			head := MergeSortedLists(tt.args.head1, tt.args.head2)

			require.True(t, IsEqual(head, tt.want))
		})

	}
}
