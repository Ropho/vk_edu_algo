package sem_2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteNode(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
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
			want: ConstructList([]Note{
				{
					Val:       11,
					NextIndex: 1,
				},
				{
					Val:       3,
					NextIndex: 2,
				},
				{
					Val:       9,
					NextIndex: 3,
				},
				{
					Val:       6,
					NextIndex: 4,
				},
				{
					Val:       11,
					NextIndex: 5,
				},
				{
					Val:       16,
					NextIndex: 6,
				},
				{
					Val:       17,
					NextIndex: -1,
				},
			},
			),
		},
		{
			name: "Positive #2",
			args: args{
				head: ConstructList([]Note{
					{
						Val:       8,
						NextIndex: 1,
					},
					{
						Val:       9,
						NextIndex: 2,
					},
					{
						Val:       6,
						NextIndex: 3,
					},
					{
						Val:       11,
						NextIndex: 4,
					},
					{
						Val:       16,
						NextIndex: 5,
					},
					{
						Val:       17,
						NextIndex: -1,
					},
				},
				),
			},
			want: ConstructList([]Note{
				{
					Val:       9,
					NextIndex: 1,
				},
				{
					Val:       6,
					NextIndex: 2,
				},
				{
					Val:       11,
					NextIndex: 3,
				},
				{
					Val:       16,
					NextIndex: 4,
				},
				{
					Val:       17,
					NextIndex: -1,
				},
			},
			),
		}, {
			name: "Positive #3: No 8 value",
			args: args{
				head: ConstructList([]Note{
					{
						Val:       11,
						NextIndex: -1,
					},
				},
				),
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
			name: "Positive #4",
			args: args{
				head: ConstructList([]Note{}),
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			head := DeleteNode(tt.args.head, 8)

			require.True(t, IsEqual(head, tt.want))
		})

	}
}
