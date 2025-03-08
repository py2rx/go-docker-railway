package main

import (
        "net/http"
        "os"

        "github.com/gin-gonic/gin"
)

type PingResponse struct {
        Status  int    `json:"status"`
        Message string `json:"message"`
        Result  string `json:"result"`
}

func main() {
        r := gin.Default()

        r.GET("/ping", func(c *gin.Context) {
                response := PingResponse{
                        Status:  http.StatusOK,
                        Message: "pong",
                        Result:  "success",
                }
                c.JSON(http.StatusOK, response)
        })

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080" // Default port
        }

        r.Run(":" + port)
}