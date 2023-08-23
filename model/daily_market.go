package model

type DailyMarket struct {
	Id   int64   `json:"id"`
	Date string  `json:"datum"`
	Sum  float64 `json:"sum"`
}
