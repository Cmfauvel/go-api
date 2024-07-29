package main

import (
    "github.com/gin-gonic/gin"
    "github.com/cmfauvel/go-api/auth/controllers"
)
//import "github.com/cmfauvel/go-api/auth/controllers"

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, world!",
        })
    })

    router.POST("/login",login.Login)

    auth := router.Group("/auth")
    auth.Use(login.AuthMiddleware())
    /* {
        auth.GET("/hello", hello)
    } */

    router.Run(":8080")
}