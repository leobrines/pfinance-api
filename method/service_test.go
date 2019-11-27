package method_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/leobrines/pfinance-api/domain"
	"github.com/leobrines/pfinance-api/method"
	"github.com/leobrines/pfinance-api/mock"
)

var mockdb *mock.MethodDB
var service *method.Service

func TestGetExistingMethodByID_Service(t *testing.T) {
	givenMockMethodDB()

	mockdb.GetByIDFn = func(id int) (*domain.Method, error) {
		return &domain.Method{Id: 1}, nil
	}

	givenServiceWithMockDB()

	user, err := service.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)
}

func givenMockMethodDB() {
	mockdb = &mock.MethodDB{}
}

func givenServiceWithMockDB() {
	service = &method.Service{
		DB: mockdb,
	}
}
