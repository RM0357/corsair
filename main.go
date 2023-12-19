package main

import (
	"io"
	"log"
	"net/http"

	// "./handlers/handlers"

	"github.com/gin-gonic/gin"
)

func getPrice(tf string, coin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := "https://api-pub.bitfinex.com/v2/candles/trade:" + tf + ":" + coin + "/hist" // 5m,30m,1h,1d 90req/min tBTCUSD
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		req.Header.Add("accept", "application/json")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request"})
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}

		response := string(body)

		c.JSON(http.StatusOK, map[string]interface{}{
			"response": response,
		})
	}
}

func getDataFromApi(tf string, coin string) string {
	url := "https://api-pub.bitfinex.com/v2/candles/trade:" + tf + ":" + coin + "/hist" // 5m,30m,1h,1d 90req/min tBTCUSD
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("ERROR")
	}

	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("ERROR")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("ERROR")
	}
	return string(body)
}

func getOHLCVData(tf []string, coin string) gin.HandlerFunc { // Proper go routines. wait group. // FOR each coin in the list, do for each time frame and save as own result
	return func(c *gin.Context) {

		var m5, m30, h1, D1 string

		m5 = getDataFromApi(tf[0], coin)
		m30 = getDataFromApi(tf[1], coin)
		h1 = getDataFromApi(tf[2], coin)
		D1 = getDataFromApi(tf[3], coin)

		c.JSON(http.StatusOK, map[string]interface{}{
			"m5":  m5,
			"m30": m30,
			"1h":  h1,
			"1D":  D1,
		})
	}
}

func main() { // TODO:  add go routines and auto restart in timer in next patch. Also start saving into the database

	r := gin.Default()
	tfList := []string{"5m", "30m", "1h", "1D"}

	r.Handle("GET", "/bitcoin", getOHLCVData(tfList, "tBTCUSD"))
	r.Handle("GET", "/ethereum", getOHLCVData(tfList, "tETHUSD"))
	r.Handle("GET", "/chainlink", getOHLCVData(tfList, "tLINK:USD"))
	r.Handle("GET", "/cardano", getOHLCVData(tfList, "tADAUSD"))
	r.Handle("GET", "/polkadot", getOHLCVData(tfList, "tDOTUSD"))
	r.Run(":7777")

}
