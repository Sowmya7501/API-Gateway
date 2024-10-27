package main

import (
	"net/http"
	"net/url"
	"net/http/httputil"
	"time" // Import time package for measuring duration
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define metrics with additional path label
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "queried_path"}, // Added queried_path label
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of HTTP request durations in seconds",
			Buckets: prometheus.DefBuckets, // Default buckets
		},
		[]string{"method", "path", "queried_path"}, // Added queried_path label
	)
)

func init() {
	// Register metrics with Prometheus
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func proxyRequest(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timing the request

		url, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing URL"})
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(url)

		c.Request.URL.Scheme = url.Scheme
		c.Request.URL.Host = url.Host
		c.Request.Host = url.Host

		proxy.ServeHTTP(c.Writer, c.Request) // Serve the proxied request

		duration := time.Since(start).Seconds() // Calculate duration in seconds

		// Increment the counter and record the duration with the queried path
		httpRequestsTotal.WithLabelValues(c.Request.Method, c.Request.URL.Path, c.FullPath()).Inc()
		httpRequestDuration.WithLabelValues(c.Request.Method, c.Request.URL.Path, c.FullPath()).Observe(duration)
	}
}

func main() {
	r := gin.Default()

	// Route to task API
	r.Any("/task/*any", proxyRequest("http://task-manager:8082"))
	r.Any("/task", proxyRequest("http://task-manager:8082"))

	// Route to user API
	r.Any("/user/*any", proxyRequest("http://user-manager:8083"))
	r.Any("/user", proxyRequest("http://user-manager:8083"))

	// Start the server on port 8081 for the REST API
	go func() {
		if err := r.Run(":8081"); err != nil {
			panic(err)
		}
	}()

	// Start a separate server for metrics on port 8080
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}