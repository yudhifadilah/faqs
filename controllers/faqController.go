package controllers

import (
	"net/http"

	"faqs/config"
	"faqs/models"

	"github.com/gin-gonic/gin"
)

// Get All FAQs
func GetFAQs(c *gin.Context) {
	var faqs []models.FAQ
	config.DB.Find(&faqs)
	c.JSON(http.StatusOK, faqs)
}

// Get FAQ by ID
func GetFAQByID(c *gin.Context) {
	var faq models.FAQ
	id := c.Param("id")

	if err := config.DB.First(&faq, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}

	c.JSON(http.StatusOK, faq)
}

// Get FAQ by Category
func GetFAQByCategory(c *gin.Context) {
	var faqs []models.FAQ
	category := c.Param("category")

	if err := config.DB.Where("category = ?", category).Find(&faqs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch FAQs"})
		return
	}

	if len(faqs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No FAQs found for this category"})
		return
	}

	c.JSON(http.StatusOK, faqs)
}

// Create New FAQ
func CreateFAQ(c *gin.Context) {
	var faq models.FAQ

	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&faq)
	c.JSON(http.StatusCreated, faq)
}

// Update FAQ
func UpdateFAQ(c *gin.Context) {
	var faq models.FAQ
	id := c.Param("id")

	if err := config.DB.First(&faq, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}

	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&faq)
	c.JSON(http.StatusOK, faq)
}

// Delete FAQ
func DeleteFAQ(c *gin.Context) {
	var faq models.FAQ
	id := c.Param("id")

	if err := config.DB.First(&faq, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}

	config.DB.Delete(&faq)
	c.JSON(http.StatusOK, gin.H{"message": "FAQ deleted successfully"})
}
