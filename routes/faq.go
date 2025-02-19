package routes

import (
	"faqs/controllers"

	"github.com/gin-gonic/gin"
)

func FAQRoutes(router *gin.Engine) {
	faqRoutes := router.Group("/faqs")
	{
		faqRoutes.GET("/", controllers.GetFAQs)
		faqRoutes.GET("/:id", controllers.GetFAQByID)
		faqRoutes.GET("/category/:category", controllers.GetFAQByCategory)
		faqRoutes.POST("/", controllers.CreateFAQ)
		faqRoutes.PUT("/:id", controllers.UpdateFAQ)
		faqRoutes.DELETE("/:id", controllers.DeleteFAQ)
	}
}
