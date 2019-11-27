package currency

import (
	"github.com/go-redis/redis"
	"github.com/leobrines/pfinance-api/domain"
)

const currencyKey = "currency"

type DB struct {
	Redis *redis.Client
}

func (db *DB) GetByID(id int) (*domain.Currency, error) {
	name, err := db.Redis.LIndex(currencyKey, int64(id)-1).Result()

	if err != nil {
		return nil, err
	}

	return &domain.Currency{
		Id:   id,
		Name: name,
	}, nil
}
