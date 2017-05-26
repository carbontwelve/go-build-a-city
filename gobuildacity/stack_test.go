package gobuildacity

import "testing"

// Based upon example here http://en.cppreference.com/w/cpp/container/stack/top
func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(2)
	s.Push(6)
	s.Push(51)

	if s.Len() != 3 {
		t.Error("Expected Len to return 3, got ", s.Len())
	}

	if (s.Peek() != 51) {
		t.Error("Expected Peek to return 51, got ", s.Peek())
	}

	if s.Len() != 3 {
		t.Error("Expected Len to return 3, got ", s.Len())
	}

	p, _ := s.Pop()

	if p != 51 {
		t.Error("Expected Pop to return 51, got ", p)
	}

	if s.Len() != 2 {
		t.Error("Expected Len to return 2, got ", s.Len())
	}

	if (s.Peek() != 6) {
		t.Error("Expected Peek to return 6, got ", s.Peek())
	}
}