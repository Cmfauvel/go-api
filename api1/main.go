package main

import (
    "github.com/gin-gonic/gin"
    "github.com/cmfauvel/go-api/auth/controllers"
    "github.com/rs/zerolog/log"
)

func main() {
    router := gin.Default()
    log.Info().Msg("Hello from Zerolog logger in login function")

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, world!",
        })
    })

    router.POST("/login",login.Login)

    auth := router.Group("/auth")
    auth.Use(login.AuthMiddleware())

    router.Run(":8080")
}