package forth

type stack struct {
	values []interface{}
}

func (s *stack) push(v interface{}) error {
	s.values = append(s.values, v)
	return nil
}

func (s *stack) popNumber() float64 {
	return s.pop().(float64)
}

func (s *stack) pop() interface{} {
	lg := len(s.values)
	if lg == 0 {
		panic("unable to pop")
	}

	v := s.values[lg-1]
	s.values = s.values[:lg-1]

	return v
}
