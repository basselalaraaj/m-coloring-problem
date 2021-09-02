package main

import (
	"reflect"
	"testing"
)

func TestGraphColoringBFS(t *testing.T) {
	testCases := []struct {
		name  string
		graph [][]int
		n, m  int
		want  []int
	}{
		{
			name: "should return the colors to all vertices",
			graph: [][]int{
				{0, 1, 1, 1},
				{1, 0, 1, 0},
				{1, 1, 0, 1},
				{1, 0, 1, 0},
			},
			n:    4,
			m:    3,
			want: []int{1, 2, 3, 2},
		},
		{
			name: "should return empty slice",
			graph: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			n:    4,
			m:    3,
			want: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := graphColoringBFS(tc.graph, tc.n, tc.m)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
