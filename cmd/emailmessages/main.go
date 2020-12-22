package main

import (
	"context"
	"email-messages/app"
	"email-messages/config"
	"email-messages/delivery/http"
	"email-messages/repository"
	"email-messages/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var postgresDB *gorm.DB

func main() {

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logrus.Infof("Received %s signal", <-c)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	er := repository.NewEntryRepository(postgresDB)

	es := service.NewMessagesService(er)
	a := app.NewApp(es)

	http.NewHandler(e, a)

	log.Fatal(e.Start(":" + config.Cfg.Port))
}

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v\n", err)
	}
}
