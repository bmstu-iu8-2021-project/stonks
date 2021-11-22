package api_interfaces

import (
	"net/http"
	qmodels "stonks/internal/models/quotes"
)

type IQuotesApiRepo interface {
	GetIntraday1Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday5Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday15Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday30Quotes(*http.Request) (qmodels.TSMongo, error)
	GetIntraday60Quotes(*http.Request) (qmodels.TSMongo, error)
	GetDailyQuotes(*http.Request) (qmodels.TSMongo, error)
	GetWeeklyQuotes(*http.Request) (qmodels.TSMongo, error)
	GetMonthlyQuotes(*http.Request) (qmodels.TSMongo, error)
}
