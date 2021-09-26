# Gin-gonic framework installation

> $ go get -u github.com/gin-gonic/gin

# Making a router

```
// dependency
import "github.com/gin-gonic/gin
```

```go
router:= gin.Default()
```

# Define route handler

```go
// to use http.StatusOK const value, net/http has to be imported
import "net/http"

router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK,
    gin.H{ // gin.H is map[string] interface{}
    //gin.H is used to make JSON output in gin
    //nesting is possible
        "What is": "gin.H?",
    })
})
```

# Starting App

```
router.Run()
```

# In case of app size becomes bigger

1. Make `routes.go` and write `initializeRoutes()`
2. Call `initializeRoutes()` from `main.go`
