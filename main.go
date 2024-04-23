package main

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    fmt.Println("REQUEST HEADERS ==> ")
    fmt.Println(c.Request.Header)
    fmt.Println("Client IP =>", c.ClientIP())
	jsonData := []byte(`{"author":"rudimk", "message":"in remembrance of earth's past."}`)

	c.Data(http.StatusOK, "application/json", jsonData)
  })

  return r
}

func main() {
  r := setupRouter()
  // Listen and Server in 0.0.0.0:8080
  r.Run(":3000")
}