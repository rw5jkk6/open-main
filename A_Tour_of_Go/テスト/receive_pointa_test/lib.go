package main

type SwapStr struct{
	a, b int
}

func (s *SwapStr)Swap(){
	tmp := s.a
	s.a = s.b
	s.b = tmp
}