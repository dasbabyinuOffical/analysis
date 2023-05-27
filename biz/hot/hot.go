package hot

import (
	"analysis/config"
	"analysis/model"
)

func AddHotPairs(pairs []*model.Pair) (err error) {
	db := config.ClickDB()
	err = db.CreateInBatches(pairs, len(pairs)).Error
	return
}
