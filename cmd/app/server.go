package app

import (
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/server"
	"github.com/crisaltmann/fundament-stock-api/server/middlewares"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"os"
	"regexp"
)

var Server = fx.Options(
	serverfactories,
	fx.Invoke(server.MapRouter),
)

var serverfactories = fx.Provide(
	configureServer,
)

var rxURL = regexp.MustCompile(`^/regexp\d*`)

func configureServer(conf *config.Config) *server.Server {
	r := gin.New()
	configureMiddlewares(r)
	server := &server.Server{
		Config: conf,
		Server: r,
	}
	return server
}

func configureMiddlewares(r *gin.Engine) {
	r.Use(logger.SetLogger())
	r.Use(middlewares.ErrorHandler)
	r.Use(middlewares.CORSMiddleware())

	// Custom logger
	subLog := zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Logger()

	r.Use(logger.SetLogger(logger.Config{
		Logger:         &subLog,
		UTC:            true,
		SkipPath:       []string{"/skip"},
		SkipPathRegexp: rxURL,
	}))
}
