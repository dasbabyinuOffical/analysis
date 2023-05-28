package ave

import (
	"analysis/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
)

func GetHotTokens() (tokens []*model.HotToken, err error) {
	chain := BSCChain
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
	return
}
