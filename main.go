package main

import (
	"io"
	"net/http"

	// "./handlers/handlers"

	"github.com/gin-gonic/gin"
)

func getPrice(tf string) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := "https://api-pub.bitfinex.com/v2/candles/trade:" + tf + ":tBTCUSD/hist" // 5m,30m,1h,1d 90req/min
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

func main() {
	r := gin.Default()
	r.Handle("GET", "/btc5", getPrice("5m"))
	r.Handle("GET", "/btc30", getPrice("30m"))
	r.Handle("GET", "/btc1h", getPrice("1h"))
	r.Handle("GET", "/btc1d", getPrice("1D"))

	r.Run(":7777")

}
