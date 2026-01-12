package routes

import (
	"faq_sys_go/db"
	"faq_sys_go/handler"
	"faq_sys_go/middleware"
	"faq_sys_go/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	db := db.GetDB()

	userRepo := repository.NewUserRepository(db)
	storeRepo := repository.NewStoreRepository(db)
	categoryRepo := repository.NewFAQCategoryRepository(db)
	faqRepo := repository.NewFAQRepository(db)
	translationRepo := repository.NewTranslationRepository(db)

	authHandler := handler.NewAuthHandler(userRepo, storeRepo)
	categoryHandler := handler.NewFAQCategoryHandler(categoryRepo)
	faqHandler := handler.NewFAQHandler(faqRepo)
	translationHandler := handler.NewTranslationHandler(translationRepo, faqRepo)

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authHandler.SignUp)
			auth.POST("/login", authHandler.Login)
		}

		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", authHandler.GetProfile)

			categories := protected.Group("/categories")
			{
				categories.GET("", categoryHandler.GetAllCategories)
				categories.GET("/:id", categoryHandler.GetCategoryByID)

				adminCategories := categories.Group("/")
				adminCategories.Use(middleware.RequireAdmin())
				{
					adminCategories.POST("", categoryHandler.CreateCategory)
					adminCategories.PUT("/:id", categoryHandler.UpdateCategory)
					adminCategories.DELETE("/:id", categoryHandler.DeleteCategory)
				}
			}

			faqs := protected.Group("/faqs")
			{
				faqs.GET("", faqHandler.GetAllFAQs)
				faqs.GET("/:id", faqHandler.GetFAQByID)

				adminMerchantFaqs := faqs.Group("/")
				adminMerchantFaqs.Use(middleware.RequireAdminOrMerchant())
				{
					adminMerchantFaqs.POST("", faqHandler.CreateFAQ)
					adminMerchantFaqs.PUT("/:id", faqHandler.UpdateFAQ)
					adminMerchantFaqs.DELETE("/:id", faqHandler.DeleteFAQ)
				}
			}

			translations := protected.Group("/translations")
			{
				translations.GET("/faq/:faq_id", translationHandler.GetTranslationsByFAQID)
				adminMerchantTranslations := translations.Group("/")
				adminMerchantTranslations.Use(middleware.RequireAdminOrMerchant())
				{
					adminMerchantTranslations.POST("", translationHandler.CreateTranslation)
					adminMerchantTranslations.PUT("/:id", translationHandler.UpdateTranslation)
					adminMerchantTranslations.DELETE("/:id", translationHandler.DeleteTranslation)
				}
			}
		}
	}
}