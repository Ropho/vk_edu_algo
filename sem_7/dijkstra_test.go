package sem7

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDijkstra(t *testing.T) {
	type args struct {
		graph map[string]map[string]int
		start string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Positive #1",
			args: args{
				graph: map[string]map[string]int{
					"A": {
						"B": 1, "C": 5,
					},
					"B": {
						"A": 1, "C": 2, "D": 3,
					},
					"C": {
						"A": 5, "B": 2, "D": 1,
					},
					"D": {
						"B": 3, "C": 1,
					},
				},
				start: "A",
			},
			want: map[string]int{
				"A": 0,
				"B": 1,
				"C": 3,
				"D": 4,
			},
		},
		{
			name: "Positive #2",
			args: args{
				graph: map[string]map[string]int{
					"0": {
						"1": 4, "2": 1,
					},
					"1": {
						"3": 1,
					},
					"2": {
						"1": 2, "3": 5,
					},
					"3": {
						"4": 3,
					},
					"4": {},
				},
				start: "0",
			},
			want: map[string]int{
				"0": 0,
				"1": 3,
				"2": 1,
				"3": 4,
				"4": 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dijkstra(tt.args.graph, tt.args.start)

			require.Equal(t, tt.want, got)
		})
	}
}

func Test_heap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    heap
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("heap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    heap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("heap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    heap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_heap_Push(t *testing.T) {
	type args struct {
		x any
	}
	tests := []struct {
		name string
		h    *heap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
		})
	}
}

func Test_heap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *heap
		want any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
