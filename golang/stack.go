package golang

type IntStack struct {
	data []int
}

func NewIntStack() *IntStack {
	return &IntStack{}
}

func (s *IntStack) Push(d int) {
	s.data = append(s.data, d)
}

func (s *IntStack) Pop() int {
	t := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return t
}

func (s *IntStack) Top() int {
	if !s.IsEmpty() {
		return s.data[len(s.data)-1]
	}
	return 0
}

func (s *IntStack) IsEmpty() bool {
	return len(s.data) == 0
}
