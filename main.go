package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPrice(c *gin.Context) {
	url := "https://api-pub.bitfinex.com/v2/candles/trade:1m:tBTCUSD/hist"
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

func main() {
	r := gin.Default()

	r.GET("/bitcoin", getPrice)

	r.Run(":7500")
	if err != nil {
		fmt.Printf("Error %v at opening app at port 7500 ", err)
	}
}
