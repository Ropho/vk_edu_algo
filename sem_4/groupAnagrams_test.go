package sem_4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupAnagrams(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Positive #1",
			args: args{
				words: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "tea", "eat"}},
		},
		{
			name: "Positive #2",
			args: args{
				words: []string{"won", "now", "aaa", "ooo", "ooo"},
			},
			want: [][]string{{"aaa"}, {"ooo", "ooo"}, {"won", "now"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.args.words)

			sumGot := 0
			for _, slice := range got {
				for _, word := range slice {
					sumGot += anagramHash(word)
				}
			}

			sumWant := 0
			for _, slice := range tt.want {
				for _, word := range slice {
					sumWant += anagramHash(word)
				}
			}

			require.Equal(t, sumWant, sumGot)
		})
	}
}

func Test_anagramHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagramHash(tt.args.s); got != tt.want {
				t.Errorf("anagramHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
