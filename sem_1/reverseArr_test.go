package sem_1

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseArr(t *testing.T) {
	type args struct {
		arr   []int
		start int
		end   int
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
				arr:   []int{1, 2, 3, 4, 5, 6, 7},
				start: 0,
				end:   6,
			},
			want:    []int{7, 6, 5, 4, 3, 2, 1},
			wantErr: nil,
		},
		{
			name: "Positive 2",
			args: args{
				arr:   []int{3, 8, 6, 9, 9, 8, 6},
				start: 0,
				end:   6,
			},
			want:    []int{6, 8, 9, 9, 6, 8, 3},
			wantErr: nil,
		},
		{
			name: "Negative: empty slice",
			args: args{
				arr:   []int{},
				start: 0,
				end:   0,
			},
			want:    nil,
			wantErr: errors.New("validation failed for array"),
		},
		{
			name: "Negative: nil slice",
			args: args{
				arr:   []int{},
				start: 0,
				end:   0,
			},
			want:    nil,
			wantErr: errors.New("validation failed for array"),
		},
		{
			name: "Negative start index",
			args: args{
				arr:   []int{1, 2, 3, 4, 5, 6, 7},
				start: -1,
				end:   6,
			},
			want:    nil,
			wantErr: errors.New("validation failed for start index"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got, err := ReverseArr(tt.args.arr, tt.args.start, tt.args.end)

			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
