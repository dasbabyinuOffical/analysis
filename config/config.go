package config

import (
	"analysis/model"
	ck "github.com/ClickHouse/clickhouse-go/v2"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"time"
)

var (
	clickDB *gorm.DB
)

func initClickhouse() {
	sqlDB := ck.OpenDB(&ck.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: ck.Auth{
			Database: "default",
			Username: "default",
			Password: "123456",
		},
		Settings: ck.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &ck.Compression{
			Method: ck.CompressionLZ4,
		},
		Debug: true,
	})
	db, err := gorm.Open(clickhouse.New(clickhouse.Config{Conn: sqlDB}))
	if err != nil {
		panic(err)
	}
	clickDB = db
}

func ClickDB() *gorm.DB {
	return clickDB
}

func registerModel() {
	clickDB.AutoMigrate(
		&model.Token{},
		&model.Pair{},
	)
}

func Init() {
	initClickhouse()
	registerModel()
}
