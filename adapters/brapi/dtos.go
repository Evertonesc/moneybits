package brapi

type TickerQuoteResponse struct {
	Results []TickerResult `json:"results"`
}

type TickerResult struct {
	Currency                   string         `json:"currency"`
	MarketCap                  int64          `json:"marketCap"`
	ShortName                  string         `json:"shortName"`
	LongName                   string         `json:"longName"`
	RegularMarketChange        float64        `json:"regularMarketChange"`
	RegularMarketChangePercent float64        `json:"regularMarketChangePercent"`
	RegularMarketTime          string         `json:"regularMarketTime"`
	RegularMarketPrice         float64        `json:"regularMarketPrice"`
	RegularMarketDayHigh       float64        `json:"regularMarketDayHigh"`
	RegularMarketDayRange      string         `json:"regularMarketDayRange"`
	RegularMarketDayLow        float64        `json:"regularMarketDayLow"`
	RegularMarketVolume        int64          `json:"regularMarketVolume"`
	RegularMarketPreviousClose float64        `json:"regularMarketPreviousClose"`
	RegularMarketOpen          float64        `json:"regularMarketOpen"`
	FiftyTwoWeekRange          string         `json:"fiftyTwoWeekRange"`
	FiftyTwoWeekLow            float64        `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh           float64        `json:"fiftyTwoWeekHigh"`
	Symbol                     string         `json:"symbol"`
	SummaryProfile             SummaryProfile `json:"summaryProfile"`
	PriceEarnings              float64        `json:"priceEarnings"`
	EarningsPerShare           float64        `json:"earningsPerShare"`
	LogoURL                    string         `json:"logourl"`
}

type SummaryProfile struct {
	Symbol              string `json:"symbol"`
	Address1            string `json:"address1"`
	City                string `json:"city"`
	State               string `json:"state"`
	Country             string `json:"country"`
	Phone               string `json:"phone"`
	Website             string `json:"website"`
	Industry            string `json:"industry"`
	IndustryKey         string `json:"industryKey"`
	IndustryDisp        string `json:"industryDisp"`
	Sector              string `json:"sector"`
	SectorKey           string `json:"sectorKey"`
	SectorDisp          string `json:"sectorDisp"`
	LongBusinessSummary string `json:"longBusinessSummary"`
	FullTimeEmployees   int    `json:"fullTimeEmployees"`
	UpdatedAt           string `json:"updatedAt"`
}
