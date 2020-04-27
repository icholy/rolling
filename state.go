package rolling

// state is a reusable implementation of the
// circular buffer logic
type state struct {
	size   int
	start  int
	length int
}

// add returns the index where the next value
// should be inserted
func (s *state) add() int {
	end := (s.start + s.length) % s.size
	if s.length < s.size {
		s.length++
	} else {
		s.start = (s.start + 1) % s.size
	}
	return end
}

// index returns the underlying index is for the
// value at index i
func (s state) at(i int) int {
	return (s.start + i) % s.size
}
