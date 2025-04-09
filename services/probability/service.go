package probability

import "errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// P(A) = m / n
func(s *Service) Calculate(m, n int) (float64, error) {
    if n == 0 {
        return 0, errors.New("общее число исходов не может быть равно нулю")
    }
    if m > n {
        return 0, errors.New("число благоприятных исходов не может превышать общее число исходов")
    }
    if n < 0 || m < 0 {
        return 0, errors.New("число исходов не может быть отрицательным")
    }
    return float64(m) / float64(n), nil
}