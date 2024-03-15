package sem_1

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReversePart(t *testing.T) {
	type args struct {
		arr        []int
		startIndex int
	}

	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{
			name: "Positive 1",
			args: args{
				arr:        []int{1, 2, 3, 4, 5, 6, 7},
				startIndex: 1,
			},
			want:    []int{7, 1, 2, 3, 4, 5, 6},
			wantErr: nil,
		},
		{
			name: "Positive 2",
			args: args{
				arr:        []int{1, 2, 3, 4, 5, 6, 7},
				startIndex: 3,
			},
			want:    []int{5, 6, 7, 1, 2, 3, 4},
			wantErr: nil,
		},
		{
			name: "Negative: empty slice",
			args: args{
				arr:        []int{},
				startIndex: 10,
			},
			want:    nil,
			wantErr: errors.New("validation failed for array"),
		},
		{
			name: "Negative: nil slice",
			args: args{
				arr:        nil,
				startIndex: 10,
			},
			want:    nil,
			wantErr: errors.New("validation failed for array"),
		},
		{
			name: "Negative start index",
			args: args{
				arr:        []int{1, 2, 3, 4, 5, 6, 7},
				startIndex: -1,
			},
			want:    nil,
			wantErr: errors.New("validation failed for endIndex"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got, err := ReversePart(tt.args.arr, tt.args.startIndex)

			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
