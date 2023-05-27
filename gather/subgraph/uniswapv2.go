package subgraph

import (
	"analysis/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Ret struct {
	Data struct {
		Pairs []struct {
			Id         string       `json:"id"`
			Token0     *model.Token `json:"token0"`
			Token1     *model.Token `json:"token1"`
			ReserveUSD string       `json:"reserveUSD"`
			VolumeUSD  string       `json:"volumeUSD"`
		} `json:"pairs,omitempty"`
	} `json:"data,omitempty"`
}

const (
	URL = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2"
)

func GetHotPairs() (pairs []*model.Pair, tokenMap map[string]*model.Token, err error) {
	query := `
{
  pairs(
    first: 1000
    where: {reserveUSD_gt: "1000000", volumeUSD_gt: "50000"}
    orderBy: reserveUSD
    orderDirection: desc
  ) {
    id
    token0 {
      id
      symbol
    }
    token1 {
      id
      symbol
    }
    reserveUSD
    volumeUSD
  }
}
`
	request := map[string]string{
		"query": query,
	}
	req, _ := json.Marshal(request)
	resp, err := http.Post(URL, "application/json", bytes.NewReader(req))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var ret Ret
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}

	now := time.Now()
	tokenMap = make(map[string]*model.Token)
	for _, p := range ret.Data.Pairs {
		reserveUSD, _ := strconv.ParseFloat(p.ReserveUSD, 64)
		volumeUSD, _ := strconv.ParseFloat(p.VolumeUSD, 64)
		pairs = append(pairs, &model.Pair{
			Id:         p.Id,
			Token0Id:   p.Token0.Id,
			Token1Id:   p.Token1.Id,
			ReserveUSD: reserveUSD,
			VolumeUSD:  volumeUSD,
			DateTime:   now,
		})
		tokenMap[p.Token0.Id] = p.Token0
		tokenMap[p.Token1.Id] = p.Token1
	}
	return
}
