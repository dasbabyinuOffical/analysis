package hot

import (
	"analysis/config"
	"analysis/gather/ave"
	"analysis/model"
)

func AddHotTokens(tokens []*model.HotToken) (err error) {
	db := config.ClickDB()
	err = db.CreateInBatches(tokens, len(tokens)).Error
	return
}

func SyncHotPairs() (err error) {
	tokens, err := ave.GetHotTokens()
	if err != nil {
		return
	}
	err = AddHotTokens(tokens)
	if err != nil {
		return
	}
	return
}
