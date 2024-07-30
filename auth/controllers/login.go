package login

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "github.com/rs/zerolog/log"
)

var jwtKey = []byte("my_secret_key")


func Login(c *gin.Context) {
    log.Info().Msg("Hello from Zerolog logger in login function")
    var loginData struct {
        Email string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if loginData.Email != "myuser" || loginData.Password != "mypassword" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
        return
    }

    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &jwt.StandardClaims{
        ExpiresAt: expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }

        token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }

            return jwtKey, nil
        })

        if err != nil {
            if err == jwt.ErrSignatureInvalid {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
                return
            }

            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        if !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Next()
    }
}
