package details_models

// Overview function
type Overview struct {
	Symbol               string `json:"Symbol"`
	AssetType            string `json:"AssetType"`
	Name                 string `json:"Name"`
	Description          string `json:"Description"`
	Exchange             string `json:"Exchange"`
	Currency             string `json:"Currency"`
	Country              string `json:"Country"`
	Sector               string `json:"Sector"`
	Industry             string `json:"Industry"`
	FiscalYearEnd        string `json:"FiscalYearEnd"`
	DividendPerShare     string `json:"DividendPerShare"`
	ProfitMargin         string `json:"ProfitMargin"`
	MarketCapitalization string `json:"MarketCapitalization"`
	BookValue            string `json:"BookValue"`
	WeekHigh52           string `json:"52WeekHigh"`
	WeekLow52            string `json:"52WeeLow"`
	DayMovingAverage50   string `json:"50DayMovingAverage"`
	DayMovingAverage200  string `json:"200DayMovingAverage"`
	PayoutRatio          string `json:"PayoutRatio"`
	DividendDate         string `json:"DividendDate"`
}

// IncomeStatement function

type IncomeStatementReports struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       string `json:"grossProfit"`
	TotalRevenue                      string `json:"totalRevenue"`
	CostOfRevenue                     string `json:"costOfRevenue"`
	CostofGoodsAndServicesSold        string `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   string `json:"operatingIncome"`
	SellingGeneralAndAdministrative   string `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            string `json:"researchAndDevelopment"`
	OperatingExpenses                 string `json:"operatingExpenses"`
	InvestmentIncomeNet               string `json:"investmentIncomeNet"`
	NetInterestIncome                 string `json:"netInterestIncome"`
	InterestIncome                    string `json:"interestIncome"`
	InterestExpense                   string `json:"interestExpense"`
	NonInterestIncome                 string `json:"nonInterestIncome"`
	OtherNonOperatingIncome           string `json:"otherNonOperatingIncome"`
	DepreciationAndAmortization       string `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense"`
	InterestAndDebtExpense            string `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       string `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              string `json:"ebit"`
	Ebitda                            string `json:"ebitda"`
	NetIncome                         string `json:"net_income"`
}

type IncomeStatement struct {
	Symbol           string                   `json:"symbol"`
	AnnualReports    []IncomeStatementReports `json:"annualReports"`
	QuarterlyReports []IncomeStatementReports `json:"quarterlyReports"`
}

// Earnings

type AnnualEarnings struct {
	FiscalDateEnding string `json:"fiscalDateEnding"`
	ReportedEPS      string `json:"reportedEPS"`
}

type QuarterlyEarnings struct {
	FiscalDateEnding   string `json:"fiscalDateEnding"`
	ReportedDate       string `json:"reportedDate"`
	ReportedEPS        string `json:"reportedEPS"`
	EstimatedEPS       string `json:"estimatedEPS"`
	Surprise           string `json:"surprise"`
	SurprisePercentage string `json:"surprisePercentage"`
}

type Earnings struct {
	Symbol            string              `json:"symbol"`
	AnnualEarnings    []AnnualEarnings    `json:"annualEarnings"`
	QuarterlyEarnings []QuarterlyEarnings `json:"quarterlyEarnings"`
}

// Balance sheet

type BalanceSheetReports struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            string `json:"totalAssets"`
	TotalCurrentAssets                     string `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  string `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            string `json:"cashAndShortTermInvestments"`
	Inventory                              string `json:"inventory"`
	CurrentNetReceivables                  string `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  string `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 string `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE string `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       string `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      string `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               string `json:"goodwill"`
	Investments                            string `json:"investments"`
	LongTermInvestments                    string `json:"longTermInvestments"`
	ShortTermInvestments                   string `json:"shortTermInvestments"`
	OtherCurrentAssets                     string `json:"otherCurrentAssets"`
	OtherNonCurrrentAssets                 string `json:"otherNonCurrrentAssets"`
	TotalLiabilities                       string `json:"totalLiabilities"`
	TotalCurrentLiabilities                string `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 string `json:"currentAccountsPayable"`
	DeferredRevenue                        string `json:"deferredRevenue"`
	CurrentDebt                            string `json:"currentDebt"`
	ShortTermDebt                          string `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             string `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                string `json:"capitalLeaseObligations"`
	LongTermDebt                           string `json:"longTermDebt"`
	CurrentLongTermDebt                    string `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 string `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 string `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                string `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             string `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 string `json:"totalShareholderEquity"`
	TreasuryStock                          string `json:"treasuryStock"`
	RetainedEarnings                       string `json:"retainedEarnings"`
	CommonStock                            string `json:"commonStock"`
	CommonStockSharesOutstanding           string `json:"commonStockSharesOutstanding"`
}

type BalanceSheet struct {
	Symbol           string                `json:"symbol"`
	AnnualReports    []BalanceSheetReports `json:"annualReports"`
	QuarterlyReports []BalanceSheetReports `json:"quarterlyReports"`
}

// Cash Flow

type CashFlow struct {
	Symbol           string                `json:"symbol"`
	AnnualReports    []BalanceSheetReports `json:"annualReports"`
	QuarterlyReports []BalanceSheetReports `json:"quarterlyReports"`
}

// DetailsRequest

type DetailsRequest struct {
	Symbol   string `json:"symbol" validate:"required"`
	Function string `json:"function" validate:"required,oneof=OVERVIEW EARNINGS INCOME_STATEMENT BALANCE_SHEET CASH_FLOW EARNINGS_CALENDAR"`
}
