package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"strconv"
)

func RateLimiter(c *gin.Context) {
	limit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_REQUEST"))
	burst, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_BURST"))
	limiter := rate.NewLimiter(rate.Limit(limit), burst)
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
		c.Abort()
		return
	}
	c.Next()
}
