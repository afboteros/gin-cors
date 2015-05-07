package cors

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"strings"
)

// HTTP Methods
const (
	Get     = "GET"
	Post    = "POST"
	Put     = "PUT"
	Delete  = "DELETE"
	Options = "OPTIONS"
	Patch   = "PATCH"
	Head    = "HEAD"
)

// HTTP Headers
const (
	ContentType    = "Content-Type"
	ContentLength  = "Content-Length"
	AcceptEncoding = "Accept-Encoding"
	XCSRFToken     = "X-CSRF-Token"
	Authorization  = "Authorization"
	Accept         = "Accept"
	Origin         = "Origin"
	CacheControl   = "Cache-Control"
	XRequestedWith = "X-Requested-With"
)

// Default values for Options
var (
	defaultAllowOrigins     = []string{"*"}
	defaultAllowHeaders     = []string{ContentType, ContentLength, AcceptEncoding, XCSRFToken, Authorization, Accept, Origin, CacheControl, XRequestedWith}
	defaultAllowMethods     = []string{Get, Post, Put, Delete, Patch, Head} // Not managing OPTIONS as default method in order to manage it individually
	defaultAllowCredentials = true
)

// Options for Handler
type Options struct {
	AllowOrigins     []string
	AllowHeaders     []string
	AllowMethods     []string
	AllowCredentials bool
}

// Middleware for setting headers on every managed request
func Middleware(options Options) gin.HandlerFunc {
	// Setting the default origins in case not specified
	if options.AllowOrigins == nil {
		options.AllowOrigins = defaultAllowOrigins
	}
	// Setting the default headers in case not specified
	if options.AllowHeaders == nil {
		options.AllowHeaders = defaultAllowHeaders
	}
	// Setting the default methods in case not specified
	if options.AllowMethods == nil {
		options.AllowMethods = defaultAllowMethods
	}
	// Setting the default credentials in case not specified
	if options.AllowCredentials == nil {
		options.AllowCredentials = defaultAllowCredentials
	}
	// Request managing func
	return func(c *gin.Context) {
		if len(options.AllowOrigins > 0) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", strings.Join(options.AllowOrigins, " "))
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", options.AllowOrigins)
		}

		if len(options.AllowHeaders > 0) {
			c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(options.AllowHeaders, ","))
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Headers", options.AllowHeaders)
		}

		if len(options.AllowMethods > 0) {
			c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(options.AllowMethods, ","))
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Methods", options.AllowMethods)
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(options.AllowCredentials))

		/** OPTIONS Method returns no content status, this is important for example
		when requesting server when AngularJS Resources in order to avoid OPTIONS Request error
		*/
		if c.Request.Method == Options {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

