package ave

import (
	"analysis/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Ret struct {
	Status     int    `json:"status"`
	Msg        string `json:"msg"`
	DataType   int    `json:"data_type"`
	EncodeData string `json:"encode_data"`
}

type Chain string

const (
	URL            = "https://api.hserpcvice.com/v1api/v2/discover/token_list?chain=%s&category=hot&pageSize=1000"
	BSCChain Chain = "bsc"
	ARBChain Chain = "arbitrum"
	EthChain Chain = "eth"
)

var (
	chains = []Chain{BSCChain, ARBChain, EthChain}
)

func GetHotTokens() (hotTokens []*model.HotToken, err error) {
	for _, chain := range chains {
		tokens, err := getHotTokens(chain)
		if err != nil {
			return nil, err
		}
		hotTokens = append(hotTokens, tokens...)
	}
	return
}

func getHotTokens(chain Chain) (tokens []*model.HotToken, err error) {
	host := fmt.Sprintf(URL, chain)
	resp, err := http.Get(host)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var (
		ret Ret
	)
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return
	}

	if ret.Status != 1 || ret.Msg != "SUCCESS" {
		return
	}

	data, err := base64.StdEncoding.DecodeString(ret.EncodeData)
	if err != nil {
		return
	}

	escapeData, err := url.QueryUnescape(string(data))
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(escapeData), &tokens)
	if err != nil {
		return
	}

	now := time.Now()
	for _, token := range tokens {
		token.CreatedAt = now
	}

	return
}
