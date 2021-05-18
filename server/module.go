package server

import (
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/server/middlewares"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"os"
	"regexp"
)

var Module = fx.Options(
	factories,
	fx.Invoke(MapRouter),
)

var factories = fx.Provide(
	ConfigureServer,
)

var rxURL = regexp.MustCompile(`^/regexp\d*`)

func ConfigureServer(conf *config.Config) *Server {
	r := gin.New()
	configureLog(r)
	server := &Server{conf, r}
	return server
}

func configureLog(r *gin.Engine) {
	r.Use(logger.SetLogger())
	r.Use(middlewares.ErrorHandler)

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
