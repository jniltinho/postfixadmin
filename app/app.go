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
	Echo   *echo.Echo
	Secret string
	SSL    bool
	Host   string
}

func NewApp(cfg *viper.Viper) *AppConfig {

	return &AppConfig{
		Config: cfg,
		Echo:   echo.New(),
		Secret: cfg.GetString("security.jwt.key"),
		SSL:    cfg.GetBool("http.ssl"),
		Host:   cfg.GetString("http.host"),
	}
}

func (r *AppConfig) AppRun() {
	// Initialize database connection
	config.InitDBConnection(r.Config)
	// Echo instance
	//r.Echo.Logger.SetLevel(log.INFO)

	//Routes
	r.runRoutes()

	// Start server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		r.runServer()
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Echo.Shutdown(ctx); err != nil {
		r.Echo.Logger.Fatal(err)
	}

	// Close database connections when the program terminates.
	defer config.CloseDBConnection()
}

func (r *AppConfig) runServer() {
	cfg := r.Config

	switch r.SSL {
	case true:
		log.INFO("Starting server on https://%s", r.Host)
		if err := r.Echo.StartTLS(r.Host, cfg.GetString("http.cert"), cfg.GetString("http.key")); err != nil && err != http.ErrServerClosed {
			r.Echo.Logger.Fatal("Shutting down the server")
		}
	case false:
		log.INFO("Starting server on http://%s", r.Host)
		if err := r.Echo.Start(r.Host); err != nil && err != http.ErrServerClosed {
			r.Echo.Logger.Fatal("Shutting down the server")
		}
	}
}
