package webserver

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Port   int
}

func New() *Server {
	strport := os.Getenv("PORT")
	port, err := strconv.Atoi(strport)

	if err != nil {
		panic("PORT \"" + strport + "\"invalid.")
	}

	return &Server{
		Router: gin.Default(),
		Port:   port,
	}
}

func (s *Server) Start() error {
	return s.Router.Run(":" + strconv.Itoa(s.Port))
}
