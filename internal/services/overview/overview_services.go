package overview_service

import (
	"fmt"
	"net/http"
	"net/url"
	"stonks/internal/config"
	"stonks/internal/interfaces/overview_interfaces"
	om "stonks/internal/models/overview"

	"stonks/internal/constants/market"
)

type OverviewService struct {
	OverviewRepo overview_interfaces.IOverviewRepo
	Config       *config.Config
}

func (s *OverviewService) GetOverview(queryParams url.Values) (om.Overview, error) {
	queryParams.Set("function", market_constants.Overview)
	queryParams.Set("apikey", s.Config.MarketKey)
	req, _ := http.NewRequest(http.MethodGet, market_constants.URL, nil)
	req.URL.Path = market_constants.Path
	req.URL.RawQuery = queryParams.Encode()
	sss := req.URL.String()
	fmt.Print(sss)
	resp, err := s.OverviewRepo.GetOverview(req)
	if err != nil {
		return om.Overview{}, err
	}

	return resp, nil
}
