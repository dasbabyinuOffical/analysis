package hot

import (
	"analysis/config"
	"analysis/gather/subgraph"
	"analysis/model"
)

func AddHotPairs(pairs []*model.Pair) (err error) {
	db := config.ClickDB()
	err = db.CreateInBatches(pairs, len(pairs)).Error
	return
}

func AddTokens(tokenMap map[string]*model.Token) (err error) {
	return
}

func SyncHotPairs() (err error) {
	pairs, tokenMap, err := subgraph.GetHotPairs()
	if err != nil {
		return
	}
	err = AddHotPairs(pairs)
	if err != nil {
		return
	}

	err = AddTokens(tokenMap)
	if err != nil {
		return
	}
	return
}
