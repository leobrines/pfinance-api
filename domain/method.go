package domain

type Method struct {
	Id         int
	UserId     int
	CurrencyId int
	Balance    int
}

type MethodRepo interface {
	GetByID(id int) (*Method, error)
	Create(m *Method) (*Method, error)
}
