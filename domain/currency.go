package domain

type Currency struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CurrencyRepo interface {
	GetByID(id int) (*Currency, error)
}
