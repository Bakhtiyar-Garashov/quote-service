package middlewares

import (
	"bytes"
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
	lm := limiter.NewRateLimiter(time.Minute, 10, func(ctx *gin.Context) (string, error) {
		b, err := ioutil.ReadAll(ctx.Request.Body)

		if err != nil {
			fmt.Println("Couldn't parse request body:", err)
		}

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

		quoteRequestBody := new(dto.QuoteRequest)

		json.Unmarshal(b, quoteRequestBody)

		key := strconv.FormatUint(uint64(quoteRequestBody.UserID), 10)

		if key != "" {
			return key, nil
		}

		return "", errors.New("User id is missing")
	})
	return lm.Middleware()

}
