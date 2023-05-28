package model

import "time"

type HotToken struct {
	BurnAmount      string `json:"burn_amount"`
	Chain           string `json:"chain"`
	CurrentPriceUsd string `json:"current_price_usd"`
	Holders         string `json:"holders"`
	HotRank         string `json:"hot_rank"`
	LockAmount      string `json:"lock_amount"`
	LogoUrl         string `json:"logo_url"`
	OpenPrice       string `json:"open_price"`
	OtherAmount     string `json:"other_amount"`
	PriceChange     string `json:"price_change"`
	Rate            string `json:"rate"`
	RiskScore       string `json:"risk_score"`
	Symbol          string `json:"symbol"`
	Token           string `json:"token"`
	Total           string `json:"total"`
	TxAmount24H     string `json:"tx_amount_24h"`
	TxCount24H      string `json:"tx_count_24h"`
	TxVolumeU24H    string `json:"tx_volume_u_24h"`
	CreatedAt       time.Time
}

func (h *HotToken) TableName() string {
	return "hot_token"
}
