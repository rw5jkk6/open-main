package lib

type S struct{
	X, Y int
}

func (s *S)Swap(){
	tmp := s.X
	s.X = s.Y
	s.Y = tmp
}