# gin-cors
CORS Basic Middleware for [Gin Gonic]

## Installation

``` bash
$ go get github.com/afboteros/gin-cors
```

## Usage

``` go
import (
    "github.com/gin-gonic/gin"
    "github.com/afboteros/gin-cors"
)

func main() {
	g := gin.New()
	g.Use(cors.Middleware(cors.Options{}))
}
```

## AngularJS OPTIONS Management
When working with AngularJS ngResource, Options method will return error with some libraries, due to status return on this method
``` go
if c.Request.Method == Options {
    c.AbortWithStatus(http.StatusNoContent)
}
```

[Gin Gonic]: http://gin-gonic.github.io/gin/