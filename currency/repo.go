package currency

import (
	"github.com/go-redis/redis"
	. "github.com/leobrines/pfinance-api/domain"
)

const currencyKey = "currency"

type Repo struct {
	Redis *redis.Client
}

func (r *Repo) GetByID(id int) (*Currency, error) {
	name, err := r.Redis.LIndex(currencyKey, int64(id)-1).Result()

	if err != nil {
		return nil, err
	}

	return &Currency{
		Id:   id,
		Name: name,
	}, nil
}
