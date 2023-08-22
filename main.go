// main.go
package main

import (
	"https://github.com/Pepinonte/go-crud/blob/master/models"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.ConnectDatabase()  // new!

    // ...

    router.Run("localhost:8080")
}