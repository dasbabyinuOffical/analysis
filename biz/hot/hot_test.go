package hot

import (
	"analysis/config"
	"analysis/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func setup() {
	config.Init()
}

func TestAddHotPairs(t *testing.T) {
	setup()
	token0 := &model.Token{
		Id:     "token0",
		Symbol: "token0",
	}
	token1 := &model.Token{
		Id:     "token1",
		Symbol: "token1",
	}
	pairs := &model.Pair{
		Id:         uuid.New().String(),
		Token0Id:   token0.Id,
		Token1Id:   token1.Id,
		ReserveUSD: 111.1111,
		VolumeUSD:  2222.2222,
		DateTime:   time.Now(),
	}
	err := AddHotPairs([]*model.Pair{pairs})
	assert.Nil(t, err)
}
