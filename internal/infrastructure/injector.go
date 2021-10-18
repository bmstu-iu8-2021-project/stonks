package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"stonks/internal/config"

	"stonks/internal/controllers/market/details"
	nc "stonks/internal/controllers/news"
	"stonks/internal/repos/details"
	nr "stonks/internal/repos/news"
	"stonks/internal/services/details"
	ns "stonks/internal/services/news"
)

type IInjector interface {
	InjectNewsController() nc.NewsControllers
	InjectDetailsController() details.CompanyDetailsControllers
}

var env *environment

type environment struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func (e *environment) InjectNewsController() nc.NewsControllers {
	return nc.NewsControllers{
		Log: e.logger,
		NewsService: &ns.NewsService{
			NewsRepo: &nr.NewsRepo{},
			Config:   e.cfg,
		},
		Validator: validator.New(),
	}
}

func (e *environment) InjectDetailsController() details.CompanyDetailsControllers {
	return details.CompanyDetailsControllers{
		Log: e.logger,
		CompanyDetailsService: &details_service.CompanyDetailsService{
			DetailsRepo: &details_repo.CompanyDetailsRepo{},
			Config:      e.cfg,
		},
		Validator: validator.New(),
	}
}

func Injector(logger *zap.SugaredLogger, config *config.Config) (IInjector, error) {
	env = &environment{
		logger: logger,
		cfg:    config,
	}
	return env, nil
}
