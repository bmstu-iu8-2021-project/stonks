package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"stonks/internal/config"
	"stonks/internal/infrastructure"
)

var (
	cfg *config.Config
	ctx context.Context
	log *zap.SugaredLogger
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("error loading logger: %s", err)
		os.Exit(1)
		return
	}

	log = logger.Sugar()

	cfg = config.New()
	if err != nil {
		log.Fatalf("config init error: %s", err)
	}
	log.Infof("Config loaded:\n%+v", cfg)
}

func main() {
	injector, err := infrastructure.Injector(log, ctx, cfg)
	if err != nil {
		log.Fatal("main :: inject failing")
	}
	newsController := injector.InjectNewsController()
	overviewController := injector.InjectDetailsController()
	quotesController := injector.InjectQuotesController()
	growthController := injector.InjectGrowthController()
	chooseController := injector.InjectChooseController()

	router := gin.Default()

	v1 := router.Group("/stonks/v1")
	{
		v1.GET("/news", newsController.GetNews)
		v1.GET("/details", overviewController.GetCompanyDetails)
		v1.GET("/quotes", quotesController.GetQuotes)
		v1.GET("/growth", growthController.GetGrowth)
		v1.GET("/choose", chooseController.GetChoose)
	}

	err = router.Run()
	if err != nil {
		log.Fatal("main :: router start error")
	}
}
