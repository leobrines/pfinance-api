package method

import "github.com/leobrines/pfinance-api/domain"

type Service struct {
	DB domain.MethodDB
}

func (s *Service) GetByID(id int) (*domain.Method, error) {
	return s.DB.GetByID(id)
}
