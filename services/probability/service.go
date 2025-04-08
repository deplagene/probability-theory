package probability

import "errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Calculate(n, m int) (int, error) {
    if n < 0 || m < 0 {
        return 0, errors.New("cannot be zero")
    }
    return n * m, nil
}