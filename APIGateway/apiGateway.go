package main

import (
	"net/http"
	"net/url"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
)

func proxyRequest(target string) gin.HandlerFunc{
	return func(c *gin.Context) {
		url, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing URL"})
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(url) 

		c.Request.URL.Scheme = url.Scheme
		c.Request.URL.Host = url.Host
		c.Request.Host = url.Host

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()

	// Route to task API
	r.Any("/task/*any",proxyRequest("http://localhost:8081"))
	r.Any("/task",proxyRequest("http://localhost:8081"))

	// Route to user API
	r.Any("/user/*any",proxyRequest("http://localhost:8082"))
	r.Any("/user",proxyRequest("http://localhost:8082"))

	r.Run(":8080")
}

