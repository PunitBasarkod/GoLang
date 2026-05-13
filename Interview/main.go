package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/myproject/myname", getName)

    router.Run(":8080")
}

func getName(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello World Punit Kumar",
    })
}
