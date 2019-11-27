package mock

import "github.com/leobrines/pfinance-api/domain"

type MethodDB struct {
	GetByIDFn func(id int) (*domain.Method, error)
	CreateFn  func(m *domain.Method) (*domain.Method, error)
}

func (db *MethodDB) GetByID(id int) (*domain.Method, error) {
	return db.GetByIDFn(id)
}

func (db *MethodDB) Create(m *domain.Method) (*domain.Method, error) {
	return db.CreateFn(m)
}
