package sem_4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShellSortInt(t *testing.T) {
	type args struct {
		sl []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive #1",
			args: args{
				sl: []int{-1, 2, -2, -1, 0, 1, 0},
			},
			want: []int{-2, -1, -1, 0, 0, 1, 2},
		},
		{
			name: "Positive #2",
			args: args{
				sl: []int{12, 11, 9, 8, 16, 24, 19, 3, 7, 88, 72, 63},
			},
			want: []int{3, 7, 8, 9, 11, 12, 16, 19, 24, 63, 72, 88},
		},
		{
			name: "Positive #3",
			args: args{
				sl: []int{},
			},
			want: []int{},
		},
		{
			name: "Positive #4",
			args: args{
				sl: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := ShellSort[int](tt.args.sl)

			require.Equal(t, got, tt.want)
		})
	}
}

func TestShellSortString(t *testing.T) {
	type args struct {
		sl []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Positive #1",
			args: args{
				sl: []string{"b", "a", "qqq", "aa"},
			},
			want: []string{"a", "aa", "b", "qqq"},
		},
		{
			name: "Positive #3",
			args: args{
				sl: []string{},
			},
			want: []string{},
		},
		{
			name: "Positive #4",
			args: args{
				sl: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := ShellSort[string](tt.args.sl)

			require.Equal(t, got, tt.want)
		})
	}
}
