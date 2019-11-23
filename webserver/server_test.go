package webserver_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/leobrines/pfinance-api/webserver"
)

func TestMain(m *testing.M) {
	os.Setenv("PORT", "8080")
	quit := m.Run()
	os.Exit(quit)
}

func TestNewWebserverWithInvalidPort(t *testing.T) {
	os.Setenv("PORT", "test")

	assert.Panics(t, func() {
		webserver.New()
	})

	os.Setenv("PORT", "8080")
}

func TestStart(t *testing.T) {
	server := webserver.New()
	var err error

	go func() {
		err = server.Start()
	}()

	time.Sleep(time.Second * 2)
	assert.Nil(t, err)
}
