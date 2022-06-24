package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"summershare/config"
)

func fiberConnURL() string {
	PORT := config.Config.App.PORT
	if PORT == "" {
		PORT = "3000"
	}

	return fmt.Sprintf("0.0.0.0:%s", PORT)
}

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Error().Err(err).Msg("Oops... Server is not shutting down! Reason")
		}

		close(idleConnsClosed)
	}()

	fiberConnURL := fiberConnURL()

	if err := a.Listen(fiberConnURL); err != nil {
		log.Error().Err(err).Msg("Oops... Server is not running! Reason")
	}

	<-idleConnsClosed
}

func StartServer(a *fiber.App) {
	fiberConnURL := fiberConnURL()

	if err := a.Listen(fiberConnURL); err != nil {
		log.Error().Err(err).Msg("Oops... Server is not running! Reason")
	}
}
