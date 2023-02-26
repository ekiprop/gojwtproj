package handlers

import (
	"net/http"

	"github.com/ekiprop/gojwtproj/models"
	"github.com/ekiprop/gojwtproj/utils"
	"github.com/gin-gonic/gin"
)

type NewGrocery struct {
	Name     string `json: "name" binding: "required"`
	Quantity int    `json: "quantity" binding: "required"`
}

func (s *Server) GetGroceries(c *gin.Context) {

	var groceries []models.Grocery

	if err := s.db.Find(&groceries).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, groceries)

}

func (s *Server) PostGrocery(c *gin.Context) {

	var grocery NewGrocery

	if err := c.ShouldBindJSON(&grocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGrocery := models.Grocery{Name: grocery.Name, Quantity: grocery.Quantity}

	if err := s.db.Create(&newGrocery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newGrocery)
}

func (s *Server) GetGroceries(c *gin.Context) {

	user, err := utils.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user.Groceries})

}

func (s *Server) PostGrocery(c *gin.Context) {

	var grocery models.Grocery

	if err := c.ShouldBindJSON(&grocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grocery.UserId = user.ID

	if err := s.db.Create(&grocery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grocery)
}
