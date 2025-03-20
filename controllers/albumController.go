package controllers

import (
	"api/models"
	"api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) { //Using gin api context
	c.IndentedJSON(http.StatusOK, repositories.GetAbums()) //Get from the repositories the albums and parsed to JSON
}

func AddAlbum(c *gin.Context) { //Gin context = request context
	var newAlbum models.Album                     //Creates a new "Object not object" album
	if err := c.BindJSON(&newAlbum); err != nil { //Gets the body request from incoming json, and converts the json to the album struct

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request my friend"})
		return
	}
	repositories.AddAlbum(newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}
