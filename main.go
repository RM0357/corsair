package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://api-pub.bitfinex.com/v2/candles/trade%3A1m%3AtBTCUSD/hist"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
