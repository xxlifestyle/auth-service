package main

import (
	"fmt"

	"github.com/xxlifestyle/notes_auth-service/internal/config"
	"github.com/xxlifestyle/notes_auth-service/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("logger was successfully initialized")
	logger.Info("config initialization start")

	cfg := config.GetConfig()
	fmt.Println(cfg.Listen.BindIP)
	router := httprouter.New()

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {

}
