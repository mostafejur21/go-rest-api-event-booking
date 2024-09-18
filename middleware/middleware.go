package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/mostafejur21/go_rest_api/utils"
)

func Authenticate (context *gin.Context){

    // extracting token
    // looking for the token for the Authorization header.
    token := context.Request.Header.Get("Authorization")
    if token == "" {
        // AbortWithStatusJSON will abort current request and send below request
        context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})

        return
    }

    userId, err := utils.VerifyJwtToken(token)
    if err != nil {

        context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
        return
    }
    context.Set("userId" , userId)
    context.Next()
}
