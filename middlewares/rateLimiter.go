package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/Bakhtiyar-Garashov/quote-service/dto"
	"github.com/gin-gonic/gin"
	limiter "github.com/julianshen/gin-limiter"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	lm := limiter.NewRateLimiter(time.Minute, 1, func(ctx *gin.Context) (string, error) {
		b, err := ioutil.ReadAll(ctx.Request.Body)

		ctx.Request.Close = true

		quoteRequestBody := new(dto.QuoteRequest)

		if err != nil {
			fmt.Println("Couldn't parse request body:", err)
		}

		json.Unmarshal(b, quoteRequestBody)

		key := strconv.FormatUint(uint64(quoteRequestBody.UserID), 10)

		if key != "" {
			return key, nil
		}

		return "", errors.New("User id is missing")
	})
	return lm.Middleware()

}
