package main

import (
	"faq_sys_go/db"

	"github.com/gin-gonic/gin"
)

func main() {
  db.InitDB()
	server := gin.Default()
  server.Run(":8080")
}