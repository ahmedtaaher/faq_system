package main

import (
	"faq_sys_go/config"
	"faq_sys_go/db"
	"faq_sys_go/routes"
	"faq_sys_go/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	utils.InitJWT(cfg.JWTSecret)
	db.InitDB(cfg)
	router := gin.Default()
	routes.SetupRoutes(router)
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}