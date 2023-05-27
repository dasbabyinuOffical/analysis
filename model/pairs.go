package model

import (
	"time"
)

type Pair struct {
	Id         string    `gorm:"column:id" json:"id"`
	Token0Id   string    `gorm:"column:token0_id" json:"token0Id"`
	Token1Id   string    `gorm:"column:token1_id" json:"token1Id"`
	ReserveUSD float64   `gorm:"column:reserve_usd" json:"reserveUSD"`
	VolumeUSD  float64   `gorm:"column:volume_usd" json:"volumeUSD"`
	DateTime   time.Time `gorm:"column:date_time" json:"dateTime"`
}

func (p *Pair) TableName() string {
	return "pair"
}

type Token struct {
	Id     string `gorm:"column:id" json:"id"`
	Symbol string `gorm:"column:symbol" json:"symbol"`
}

func (t *Token) TableName() string {
	return "token"
}
