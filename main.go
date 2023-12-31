// main.go
package main

import (
	"/Pepinonte/go-crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.ConnectDatabase() 

    router.Run("localhost:8080")
}