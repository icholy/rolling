package rolling

import (
	"strconv"
	"testing"

	"gotest.tools/assert"
)

func TestStateLen(t *testing.T) {
	tests := []struct {
		size   int
		adds   int
		length int
	}{
		{size: 1, adds: 0, length: 0},
		{size: 1, adds: 1, length: 1},
		{size: 1, adds: 10, length: 1},
		{size: 2, adds: 1, length: 1},
		{size: 2, adds: 10, length: 2},
		{size: 42, adds: 12, length: 12},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := state{size: tt.size}
			for i := 0; i < tt.adds; i++ {
				s.add()
			}
			assert.Equal(t, s.length, tt.length)
		})
	}
}

func TestStateAt(t *testing.T) {
	tests := []struct {
		size  int
		adds  int
		index int
		want  int
	}{
		{size: 1, adds: 1, index: 0, want: 0},
		{size: 1, adds: 2, index: 0, want: 0},
		{size: 2, adds: 1, index: 0, want: 0},
		{size: 2, adds: 2, index: 1, want: 1},
		{size: 2, adds: 3, index: 0, want: 1},
		{size: 5, adds: 2, index: 2, want: 2},
		{size: 5, adds: 5, index: 5, want: 0},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := state{size: tt.size}
			for i := 0; i < tt.adds; i++ {
				s.add()
			}
			assert.Equal(t, s.at(tt.index), tt.want)
		})
	}
}

func TestStateAdd(t *testing.T) {
	tests := []struct {
		size int
		adds []int
	}{
		{size: 1, adds: []int{0, 0, 0}},
		{size: 2, adds: []int{0, 1, 0, 1}},
		{size: 3, adds: []int{0, 1, 2, 0, 1, 2}},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := state{size: tt.size}
			adds := make([]int, len(tt.adds))
			for i := range tt.adds {
				adds[i] = s.add()
			}
			assert.DeepEqual(t, adds, tt.adds)
		})
	}
}
