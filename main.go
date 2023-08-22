// main.go
package main

import (
		"./models"
		"./bdd.go"
)

func main() {
    router := gin.Default()

    models.ConnectDatabase()  // new!

    // ...

    router.Run("localhost:8080")
}