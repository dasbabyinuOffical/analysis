package model

import (
	"time"
)

type Pair struct {
	Id         string    `gorm:"column:id" json:"id"`
	Token0Id   string    `gorm:"column:token0_id" json:"token0Id"`
	Token1Id   string    `gorm:"column:token1_id" json:"token1Id"`
	ReserveUSD float64   `json:"reserveUSD"`
	VolumeUSD  float64   `json:"volumeUSD"`
	DateTime   time.Time `json:"dateTime"`
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
