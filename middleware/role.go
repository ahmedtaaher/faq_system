package middleware

import (
	"faq_sys_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exists := c.Get("userType")
		if !exists || userType != "admin" {
			utils.ErrorResponse(c, http.StatusForbidden, "Admin access required")
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireMerchant() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exists := c.Get("userType")
		if !exists || userType != "merchant" {
			utils.ErrorResponse(c, http.StatusForbidden, "Merchant access required")
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireAdminOrMerchant() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exists := c.Get("userType")
		if !exists || (userType != "admin" && userType != "merchant") {
			utils.ErrorResponse(c, http.StatusForbidden, "Admin or Merchant access required")
			c.Abort()
			return
		}
		c.Next()
	}
}