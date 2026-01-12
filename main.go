package main

import (
	"faq_sys_go/config"
	"faq_sys_go/db"

	"github.com/gin-gonic/gin"
)

func main() {
  cfg := config.LoadConfig()
  db.InitDB(cfg)
  server := gin.Default()
  server.Run(":8080")
}