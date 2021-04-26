package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mserebryaakov/course-golang/BinanceRequest/postgresql"
)

type Data struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

//Get запрос для https://www.binance.com/api/v3/ticker/price
func runGet() *[]Data {
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

	return &data
}

func main() {
	st := postgresql.New()
	st.Open()
	defer st.Close()
	data := runGet()
	for _, value := range *data {
		st.InsertBinanceData(value.Symbol, value.Price)
	}
}
