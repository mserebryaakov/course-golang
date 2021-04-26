package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

//Get запрос для https://www.binance.com/api/v3/ticker/price
func RunGet() (*[]Data, error) {
	URL := "https://www.binance.com/api/v3/ticker/price"
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	var data []Data
	if err = json.Unmarshal(respBody, &data); err != nil {
		log.Fatal(err)
	}

	return &data, nil
}
