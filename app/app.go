package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"postfixadmin/config"
	"postfixadmin/log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Config *viper.Viper
	App    *echo.Echo
	Secret string
	SSL    bool
}

func NewAppConfig(config *viper.Viper) *AppConfig {
	return &AppConfig{Config: config}
}

func (a *AppConfig) AppRun() {
	conf := a.Config
	a.App = echo.New()
	a.Secret = conf.GetString("security.jwt.key")
	a.SSL = conf.GetBool("http.ssl")

	config.InitDBConnection(conf)
	// Echo instance
	app := a.App
	//app.Logger.SetLevel(log.INFO)

	//Routes
	a.Routes()

	// Start server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		a.StartServer()
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}

	// Close database connections when the program terminates.
	defer config.CloseDBConnection()
}

func (a *AppConfig) StartServer() {
	conf := a.Config
	app := a.App
	host := conf.GetString("http.host")

	switch a.SSL {
	case true:
		log.INFO("Starting server on https://%s", host)
		if err := app.StartTLS(host, conf.GetString("http.cert"), conf.GetString("http.key")); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("Shutting down the server")
		}
	case false:
		log.INFO("Starting server on http://%s", host)
		if err := app.Start(host); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("Shutting down the server")
		}
	}
}
