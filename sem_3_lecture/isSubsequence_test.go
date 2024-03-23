package sem3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSubsequenceQueue(t *testing.T) {
	type args struct {
		seq    string
		subseq string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive #1",
			args: args{
				seq:    "quefio",
				subseq: "uio",
			},
			want: true,
		},
		{
			name: "Positive #2",
			args: args{
				seq:    "🤐😀😚🤣🤑😇🥰",
				subseq: "😀🤣😇🥰",
			},
			want: true,
		},
		{
			name: "Positive #3",
			args: args{
				seq:    "",
				subseq: "",
			},
			want: true,
		},
		{
			name: "Positive #4",
			args: args{
				seq:    "qwe",
				subseq: "",
			},
			want: true,
		},
		{
			name: "Negative #1",
			args: args{
				seq:    "quefio",
				subseq: "uinpaoisdpoijasdpoijcnnqwopejo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := isSubsequenceQueue(tt.args.seq, tt.args.subseq)

			require.Equal(t, tt.want, got)
		})
	}
}

func Test_isSubsequencePointers(t *testing.T) {
	type args struct {
		seq    string
		subseq string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive #1",
			args: args{
				seq:    "quefio",
				subseq: "uio",
			},
			want: true,
		},
		{
			name: "Positive #2",
			args: args{
				seq:    "🤐😀😚🤣🤑😇🥰",
				subseq: "😀🤣😇🥰",
			},
			want: true,
		},
		{
			name: "Positive #3",
			args: args{
				seq:    "",
				subseq: "",
			},
			want: true,
		},
		{
			name: "Positive #4",
			args: args{
				seq:    "qwe",
				subseq: "",
			},
			want: true,
		},
		{
			name: "Negative #1",
			args: args{
				seq:    "quefio",
				subseq: "uinpaoisdpoijasdpoijcnnqwopejo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := isSubsequencePointers(tt.args.seq, tt.args.subseq)

			require.Equal(t, tt.want, got)
		})
	}
}
