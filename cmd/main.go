package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"summershare/config"
	"summershare/internal/routes"
	"summershare/pkg/database"
	"summershare/pkg/middleware"
	"summershare/pkg/utils"
)

var (
	DevMode = flag.Bool("dev", false, "enable dev mode")
)

func init() {
	flag.Parse()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if *DevMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	config.InitConfig(*DevMode)
	database.InitDatabase()
	database.AutoMigrate(database.DB)
}

func main() {
	app := fiber.New()

	middleware.RegisterMiddleware(app)

	routes.RegisterAuthRoutes(app, database.DB)
	routes.RegisterPostRoute(app, database.DB)

	if *DevMode {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
