package sem3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteDuplicates(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive #1",
			args: args{
				s: "cdffd",
			},
			want: "c",
		},
		{
			name: "Positive #2",
			args: args{
				s: "uioouiouuo",
			},
			want: "uiui",
		},
		{
			name: "Positive #3",
			args: args{
				s: "xyyx",
			},
			want: "",
		},
		{
			name: "Positive #4",
			args: args{
				s: "fqffqzzsd",
			},
			want: "fsd",
		},
		{
			name: "Positive #5",
			args: args{
				s: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteDuplicates(tt.args.s)
			require.Equal(t, tt.want, got)
		})
	}
}
